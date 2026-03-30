package manifest

import (
  "encoding/json"
  "os"
  "sync"
  "vertesan/hailstorm/rich"
)

type DownloadStatus string

const (
  StatusPending DownloadStatus = "pending"
  StatusFailed DownloadStatus = "failed"
  StatusDownloaded DownloadStatus = "downloaded"
  StatusDecrypted DownloadStatus = "decrypted"
)

type ProgressEntry struct {
  RealName string `json:"real_name"`
  Label    string `json:"label"`
  Status   DownloadStatus `json:"status"`
}

type Progress struct {
  ResInfo string `json:"res_info"`
  Entries map[string]ProgressEntry `json:"entries"`
  mu sync.Mutex
  dirty int
  FlushInterval int `json:"flush_interval"`
}

func LoadProgress(file string) *Progress {
  if _, err := os.Stat(file); os.IsNotExist(err) {
    return nil
  }
  data, err := os.ReadFile(file)
  if err != nil {
    rich.Panic("failed to read progress file: %v", err)
  }
  var progress Progress
  err = json.Unmarshal(data, &progress)
  if err != nil {
    rich.Panic("failed to unmarshal progress file: %v", err)
  }
  return &progress
}

func NewProgress(resInfo string, catalog *Catalog, flushInterval int) *Progress {
  entries := make(map[string]ProgressEntry, len(catalog.Entries))
  for _, e := range catalog.Entries {
    entries[e.RealName] = ProgressEntry{
      RealName: e.RealName,
      Label: e.StrLabelCrc,
      Status: StatusPending,
    }
  }
  if flushInterval <= 0 {
    flushInterval = 50
  }
  return &Progress{
    ResInfo: resInfo,
    Entries: entries,
    FlushInterval: flushInterval,
  }
}

func (p *Progress) saveLocked(path string) {
  data, err := json.MarshalIndent(p, "", "  ")
  if err != nil {
    rich.Panic("failed to marshal progress file: %v", err)
  }
  if err := os.WriteFile(path, data, 0644); err != nil {
    rich.Panic("failed to write progress file: %v", err)
  }
}

func (p *Progress) Save(path string) {
  p.mu.Lock()
  defer p.mu.Unlock()
  p.saveLocked(path)
}

func (p *Progress) GetStatus(realName string) DownloadStatus {
  p.mu.Lock()
  defer p.mu.Unlock()
  if entry, ok := p.Entries[realName]; ok {
    return entry.Status
  }
  return StatusPending
}

func (p *Progress) UpdateStatus(realName string, status DownloadStatus, path string) {
  p.mu.Lock()
  defer p.mu.Unlock()
  entry := p.Entries[realName]
  entry.Status = status
  p.Entries[realName] = entry
  p.dirty++
  if p.dirty >= p.FlushInterval {
    p.dirty = 0
    p.saveLocked(path)
  }
}

func (p *Progress) Cleanup(path string) bool {
  p.mu.Lock()
  defer p.mu.Unlock()
  for _, entry := range p.Entries {
    if entry.Status == StatusFailed {
      p.saveLocked(path)
      return false
    }
  }
  os.Remove(path)
  return true
}