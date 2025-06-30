package main

import (
  "flag"
  "fmt"
  "os"
  "reflect"
  "regexp"

  "vertesan/hailstorm/analyser"
  "vertesan/hailstorm/manifest"
  "vertesan/hailstorm/master"
  "vertesan/hailstorm/network"
  "vertesan/hailstorm/rich"
  "vertesan/hailstorm/utils"
)

var (
  manifestSaveDir        = "cache"
  assetsSaveDir          = "cache/assets"
  decrpytedAssetsSaveDir = "cache/plain"
  dbSaveDir              = "masterdata"

  catalogVersionFile  = "cache/currentVersion.txt"
  catalogJsonFile     = "cache/catalog.json"
  catalogJsonFilePrev = "cache/catalog_prev.json"
  catalogJsonDiffFile = "cache/catalog_diff.json"
  updatedFlagFile     = "cache/updated"
)

func doAnalyze() {
  rich.Info("Start analyzing code...")
  analyser.Analyze()
  rich.Info("Analysis completed.")
}

// Removes already existed or unchanged entries.
func diff(catalog *manifest.Catalog, outDatedCatalog *manifest.Catalog) {
  rich.Info("Start doing diff.")
  oldMap := make(map[uint64]manifest.Entry)
  for _, entry := range outDatedCatalog.Entries {
    oldMap[entry.LabelCrc] = entry
  }
  entries := []manifest.Entry{}

  for _, entry := range catalog.Entries {
    if oldEntry, ok := oldMap[entry.LabelCrc]; ok {
      if entry.Checksum == oldEntry.Checksum {
        continue
      }
    }
    rich.Info("Found a new or updated entry [%s].", entry.StrLabelCrc)
    entries = append(entries, entry)
  }
  utils.WriteToJsonFile(entries, catalogJsonDiffFile)
  catalog.Entries = entries
}

// Filter out non-DB typed assets.
func filterDb(catalog *manifest.Catalog) {
  s := []manifest.Entry{}
  for _, entry := range catalog.Entries {
    if entry.StrTypeCrc == "tsv" {
      s = append(s, entry)
    }
  }
  catalog.Entries = s
}

// Filter out assets that don't match regex patterns.
func filterByRegex(catalog *manifest.Catalog, pattern string) {
  re, err := regexp.Compile(pattern)
  if err != nil {
    rich.Error("Invalid regex pattern: %v", err)
    return
  }
  
  s := []manifest.Entry{}
  for _, entry := range catalog.Entries {
    if re.MatchString(entry.StrLabelCrc) {
      s = append(s, entry)
    }
  }
  catalog.Entries = s
}

func main() {
  // parse arguments
  fAnalyze := flag.Bool("analyze", false, "Do code analysis and exit.")
  fDbOnly := flag.Bool("dbonly", false, "Only download and decrypt DB files, put assets aside.")
  fForce := flag.Bool("force", false, "Ignore current cached version and update caches.")
  fKeepRaw := flag.Bool("keepraw", false, "Do not delete encrypted raw asset files after decrypting.")
  fConvert := flag.Bool("convert", false, "Only generate cache/plain from existing cache/assets without downloading.")
  fKeepPath := flag.Bool("keep-path", false, "Imitate url download path on file system for assets.")
  fClientVersion := flag.String("client-version", "", "Specify client version manually.")
  fResInfo := flag.String("res-info", "", "Specify resource info manually.")
  fFilterRegex := flag.String("filter-regex", "", "Only download assets that match the regex pattern. eg. --filter-regex=\"bgm_.*\"")
  flag.Parse()

  if *fAnalyze {
    doAnalyze()
    return
  }

  if *fConvert {
    rich.Info("Convert mode: generating cache/plain from existing cache/assets...")
    
    // Read existing catalog if it exists
    if _, err := os.Stat(catalogJsonFile); os.IsNotExist(err) {
      rich.Panic("No existing catalog found. Run without -convert first to download assets.")
    }
    
    entries := []manifest.Entry{}
    if err := utils.ReadFromJsonFile(catalogJsonFile, &entries); err != nil {
      panic(err)
    }
    
    catalog := &manifest.Catalog{
      Entries: entries,
    }
    
    // Only decrypt existing assets
    manifest.DecryptAllAssets(catalog, decrpytedAssetsSaveDir, assetsSaveDir)
    
    rich.Info("Conversion completed.")
    return
  }

  if err := os.Remove(updatedFlagFile); err != nil {
    if !os.IsNotExist(err) {
      panic(err)
    }
  }
  clientVersion := *fClientVersion
  resInfo := *fResInfo
  if clientVersion == "" {
    var err error
    if clientVersion, err = network.GetPlayVersion(); err != nil {
      panic(err)
    }
  }
  if resInfo == "" {
  	resInfo = network.Login(clientVersion)
  }

  currentVer, err := os.ReadFile(catalogVersionFile)
  if err != nil {
    if !os.IsNotExist(err) {
      panic(err)
    }
  }

  if !*fForce && resInfo == string(currentVer) {
    rich.Info("Nothing updated, will be stopping process.")
    return
  }

  rich.Info("New resource version: %q.", resInfo)

  // init manifest
  mani := new(manifest.Manifest)
  mani.Init(resInfo, clientVersion)

  // download catalog
  network.DownloadManifestSync(mani.RealName, manifestSaveDir)
  catalogFile, err := os.Open(fmt.Sprintf("%v/%v", manifestSaveDir, mani.RealName))
  if err != nil {
    panic(err)
  }

  // init catalog from downloaded file
  catalog := new(manifest.Catalog)
  catalog.Init(mani, catalogFile)

  if err = catalogFile.Close(); err != nil {
    panic(err)
  }
  if err = os.Remove(fmt.Sprintf("%v/%v", manifestSaveDir, mani.RealName)); err != nil {
    panic(err)
  }

  // rename outdated catalog file
  if err := os.Rename(catalogJsonFile, catalogJsonFilePrev); err != nil {
    if !os.IsNotExist(err) {
      panic(err)
    }
  }
  rich.Info("Outdated catalog was renamed to '%s'.", catalogJsonFilePrev)

  // marshal catalog to a json file
  utils.WriteToJsonFile(catalog.Entries, catalogJsonFile)

  oldEntries := []manifest.Entry{}
  if err := utils.ReadFromJsonFile(catalogJsonFilePrev, &oldEntries); err != nil {
    if !os.IsNotExist(err) {
      panic(err)
    }
  }

  oldCatalog := &manifest.Catalog{
    Entries: oldEntries,
  }

  if !*fForce {
    diff(catalog, oldCatalog)
  }

  if *fDbOnly {
    filterDb(catalog)
  }

  if *fFilterRegex != "" {
    filterByRegex(catalog, *fFilterRegex)
  }

  if len(catalog.Entries) == 0 {
    rich.Info("Nothing is updated, will be stopping process.")
    return
  }

  // download all assets
  network.DownloadAssetsAsync(catalog, assetsSaveDir, fKeepPath)

  // decrypt all assets
  manifest.DecryptAllAssets(catalog, decrpytedAssetsSaveDir, assetsSaveDir)

  // convert db files to yaml
  if err := os.MkdirAll(dbSaveDir, 0755); err != nil {
    panic(err)
  }
  errCount := 0
  for _, entry := range catalog.Entries {
    if entry.StrTypeCrc != "tsv" {
      continue
    }
    dbFile, err := os.Open(decrpytedAssetsSaveDir + "/" + entry.StrLabelCrc)
    if err != nil {
      panic(err)
    }
    // hand a instance to master.Parse so the compiler can do the inference
    ins, ok := master.MasterMap[entry.StrLabelCrc]
    if !ok {
      rich.Error("Database %q does not exist. Perhaps `master.MasterMap` needs update.", entry.StrLabelCrc)
      errCount++
      continue
    }
    rows, err := master.Parse(dbFile, entry.StrLabelCrc, &ins)
    if err != nil {
      rich.Error("An error occurred when parsing database.")
      rich.Error(err.Error())
      continue
    }
    // marshal catalog to a yaml file
    utils.WriteToYamlFile(rows, dbSaveDir+"/"+reflect.TypeOf(ins).Name()+".yaml")
  }
  cvf, err := os.Create(catalogVersionFile)
  if err != nil {
    panic(err)
  }
  if _, err := cvf.WriteString(resInfo); err != nil {
    panic(err)
  }
  if errCount > 0 {
    rich.Error("%d Error(s) occurred during parsing, please check the log.", errCount)
  }
  rich.Info("All databases parsed.")

  if _, err = os.Create(updatedFlagFile); err != nil {
    panic(err)
  }

  if !*fKeepRaw {
    if err := os.RemoveAll(assetsSaveDir); err != nil {
      panic(err)
    }
  }
}
