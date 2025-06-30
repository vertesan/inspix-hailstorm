package network

import (
  "bufio"
  "context"
  "fmt"
  "net/http"
  "os"
  "path"
  "sync"
  "time"

  "vertesan/hailstorm/manifest"
  "vertesan/hailstorm/rich"

  "golang.org/x/sync/semaphore"
)

const MAX_CONCURRENCY = 10
const MAX_RETRIES = 3
const ORIGIN = "https://assets.link-like-lovelive.app"

var (
  ctx       = context.Background()
  transport = &http.Transport{
    DisableKeepAlives: true,
  }
  client = &http.Client{
    // Timeout seconds for downloading a single file.
    // As of 2024.2, the largest file size is 363MB (one of the feslive videos),
    // make sure this value is large enough for downloading large files.
    Timeout:   1200 * time.Second,
    Transport: transport,
  }
)

type SafeCounter struct {
  mutex sync.Mutex
  num   int
}

func (c *SafeCounter) Increase() {
  c.mutex.Lock()
  c.num++
  c.mutex.Unlock()
}

func (c *SafeCounter) Value() int {
  c.mutex.Lock()
  defer c.mutex.Unlock()
  return c.num
}

func (c *SafeCounter) Clear() {
  c.mutex.Lock()
  c.num = 0
  c.mutex.Unlock()
}

func DownloadManifestSync(realName string, saveDir string) {
  counter := &SafeCounter{}
  entry := &manifest.Entry{
    RealName:     realName,
    ResourceType: 999,
    StrLabelCrc:  "Manifest",
  }
  if err := os.MkdirAll(saveDir, 0755); err != nil {
    panic(err)
  }
  downloadOne(entry, saveDir, assetHeader, nil, counter, 1)
  rich.Info("Manifest is successfully downloaded.")
}

func DownloadAssetsAsync(catalog *manifest.Catalog, downloadDir string, keepPath *bool) {
  sem := semaphore.NewWeighted(MAX_CONCURRENCY)
  dlAmount := len(catalog.Entries)
  counter := &SafeCounter{}
  if err := os.MkdirAll(downloadDir, 0755); err != nil {
    panic(err)
  }

  for _, entry := range catalog.Entries {
    // if already exists, skip downloading
    // if _, err := os.Stat(downloadDir + "/" + entry.RealName); err == nil {
    //   counter.Increase()
    //   rich.Info("(%d/%d) File %q already exists, skip downloading.", counter.Value(), dlAmount, entry.RealName)
    //   continue
    // }
    saveToDir := downloadDir
    if *keepPath {
      var resType string
      if entry.ResourceType <= 1 {
        resType = "android"
      } else {
        resType = "raw"
      }
      saveToDir = path.Join(downloadDir, resType, entry.RealName[:2])
      if err := os.MkdirAll(saveToDir, 0755); err != nil {
        panic(err)
      }
    }
    if err := sem.Acquire(ctx, 1); err != nil {
      panic(err)
    }
    go downloadOne(&entry, saveToDir, assetHeader, sem, counter, dlAmount)
  }

  // wait all concurrencies completed
  if err := sem.Acquire(ctx, MAX_CONCURRENCY); err != nil {
    panic(err)
  }
  rich.Info("Successfully downloaded all assets.")
}

func downloadOne(
  entry *manifest.Entry,
  saveDir string,
  header http.Header,
  sem *semaphore.Weighted,
  counter *SafeCounter,
  amount int,
) {
  if sem != nil {
    defer sem.Release(1)
  }
  request := prepareRequest(entry, header)
  for i := range MAX_RETRIES {
    res, err := client.Do(request)
    if err != nil {
      rich.Error("%v", err)
      rich.Warning("An internal error was occurred when downloading %v, retrying...(%d/%d)", request.URL, i+1, MAX_RETRIES)
      continue
    }
    if res.StatusCode != 200 {
      rich.Error("Status code: %d, message: %v.", res.StatusCode, res.Status)
      rich.Warning("A HTTP error was occurred when downloading %v, retrying...(%d/%d)", request.URL, i+1, MAX_RETRIES)
      if err := res.Body.Close(); err != nil {
        panic(err)
      }
      continue
    }
    // save response stream to a file
    fs, err := os.Create(fmt.Sprintf("%v/%v", saveDir, entry.RealName))
    if err != nil {
      panic(err)
    }
    defer fs.Close()
    bufw := bufio.NewWriter(fs)
    if _, err := bufw.ReadFrom(res.Body); err != nil {
      rich.Error("Error reading response body: %v", err)
      fs.Close()
      res.Body.Close()
      if err := os.Remove(fmt.Sprintf("%v/%v", saveDir, entry.RealName)); err != nil {
        rich.Warning("Failed to remove incomplete file: %v", err)
      }
      rich.Warning("Failed to read response body for %v, retrying...(%d/%d)", request.URL, i+1, MAX_RETRIES)
      continue
    }
    if err := bufw.Flush(); err != nil {
      rich.Error("Error flushing buffer: %v", err)
      fs.Close()
      res.Body.Close()
      if err := os.Remove(fmt.Sprintf("%v/%v", saveDir, entry.RealName)); err != nil {
        rich.Warning("Failed to remove incomplete file: %v", err)
      }
      rich.Warning("Failed to flush buffer for %v, retrying...(%d/%d)", request.URL, i+1, MAX_RETRIES)
      continue
    }

    counter.Increase()
    rich.Info("(%d/%d) Download completed: %q(%v).", counter.Value(), amount, entry.StrLabelCrc, entry.RealName)
    res.Body.Close()
    return
  }
  // max retry exhausted
  rich.Panic("Max retries exhausted when downloading %v. Will be stopping process.", request.URL)
}

func prepareRequest(entry *manifest.Entry, header http.Header) *http.Request {
  var resType string
  if entry.ResourceType <= 1 {
    resType = "android"
  } else {
    resType = "raw"
  }
  url := fmt.Sprintf("%v/%v/%v/%v", ORIGIN, resType, entry.RealName[:2], entry.RealName)
  request, err := http.NewRequest("GET", url, nil)
  if err != nil {
    panic(err)
  }
  request.Header = header
  return request
}
