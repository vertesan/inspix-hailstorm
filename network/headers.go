package network

import "net/http"

var (
  assetHeader = http.Header{
    "User-Agent":      {"UnityPlayer/2021.3.16f1 (UnityWebRequest/1.0, libcurl/7.84.0-DEV)"},
    "Accept":          {"*/*"},
    "Accept-Encoding": {"deflate, gzip"},
    "X-Unity-Version": {"2021.3.16f1"},
  }
  loginHeader = &http.Header{
    "X-Idempotency-Key":       {"eb6afd7c69cd9a87ca1fb167b21ae95c"},
    "X-Client-Version":        {"1.10.40"},
    "User-Agent":              {"inspix-android/1.10.40"},
    "x-res-version":           {"R2402010"},
    "x-device-type":           {"android"},
    "inspix-user-api-version": {"1.0.0"},
    "Accept":                  {"application/json"},
    "x-api-key":               {"4e769efa67d8f54be0b67e8f70ccb23d513a3c841191b6b2ba45ffc6fb498068"},
    "Content-Type":            {"application/json"},
    "Accept-Encoding":         {"gzip, deflate"},
  }
)
