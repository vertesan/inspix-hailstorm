package network

import (
  "errors"
  "fmt"
  "net/http"
  "regexp"
  "strings"
  "vertesan/hailstorm/rich"

  "github.com/PuerkitoBio/goquery"
)

const GAME_ID = "com.oddno.lovelive"

func GetPlayVersion() (string, error) {
  url := fmt.Sprintf("https://play.google.com/store/apps/details?id=%s", GAME_ID)
  req, err := http.NewRequest("GET", url, nil)
  if err != nil {
    panic(err)
  }
  header := http.Header{
    "User-Agent": {"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36 Edg/119.0.0.0"},
  }
  req.Header = header
  res, err := http.DefaultClient.Do(req)
  if err != nil {
    panic(err)
  }
  defer res.Body.Close()
  if res.StatusCode != 200 {
    rich.Panic("Abnormal HTTP status code: %d. Message: %s", res.StatusCode, res.Status)
  }
  doc, err := goquery.NewDocumentFromReader(res.Body)
  if err != nil {
    panic(err)
  }
  var version string
  doc.Find("script").EachWithBreak(func(i int, s *goquery.Selection) bool {
    scriptContent := strings.TrimSpace(s.Text())
    if strings.Contains(scriptContent, "key: 'ds:5'") {
      reg := regexp.MustCompile(`\[\[\["([\d\.]+)"\]\],\[\[\[\d+\]\]\,\[\[\[\d+,"`)
      match := reg.FindStringSubmatch(scriptContent)
      version = match[1]
      return false
    }
    return true
  })
  if version == "" {
    return "", errors.New("Cannot find ds:5 in <script>, perhaps GooglePlay has changed its structure of webpage.")
  }
  return version, nil
}
