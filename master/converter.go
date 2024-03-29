package master

import (
  "os"
  "reflect"

  "vertesan/hailstorm/manifest"
  "vertesan/hailstorm/utils"
)

func Convert2Json(catalog *manifest.Catalog, dbSaveDir, srcDir string) {
  if err := os.MkdirAll(dbSaveDir, 0755); err != nil {
    panic(err)
  }
  for _, entry := range catalog.Entries {
    if entry.StrTypeCrc != "tsv" {
      continue
    }
    dbFile, err := os.Open(srcDir + "/" + entry.StrLabelCrc)
    if err != nil {
      panic(err)
    }
    // hand a instance to master.Parse so the compiler can do the inference
    ins := MasterMap[entry.StrLabelCrc]
    rows, err := Parse(dbFile, entry.StrLabelCrc, &ins)
    if err != nil {
      panic(err)
    }
    // marshal catalog to a json file
    utils.WriteToJsonFile(rows, dbSaveDir+"/"+reflect.TypeOf(ins).Name()+".json")
  }
}
