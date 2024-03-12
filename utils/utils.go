package utils

import (
  "bufio"
  "bytes"
  "os"
  "vertesan/hailstorm/rich"

  "github.com/goccy/go-json"
  "github.com/goccy/go-yaml"
)

func Must[T any](v T, err error) T {
  if err != nil {
    panic(err)
  }
  return v
}

func ReadUntilNext(buf *bufio.Reader, point byte) []byte {
  rawBytes, err := buf.ReadBytes(0x00)
  if err != nil {
    panic(err)
  }
  trimed, _ := bytes.CutSuffix(rawBytes, []byte{point})
  if len(trimed) == 0 {
    return nil
  }
  return trimed
}

func WriteToJsonFile(instance any, dst string) {
  jsonBytes, err := json.MarshalIndent(instance, "", "  ")
  if err != nil {
    panic(err)
  }
  jsonDbFile, err := os.Create(dst)
  if err != nil {
    panic(err)
  }
  jFileBuf := bufio.NewWriter(jsonDbFile)
  jFileBuf.Write(jsonBytes)
  jFileBuf.Flush()
  rich.Info("Writing json file '%s' done.", jsonDbFile.Name())
  jsonDbFile.Close()
}

func ReadFromJsonFile(src string, v any) error {
  jBytes, err := os.ReadFile(src)
  if err != nil {
    return err
  }
  if err := json.Unmarshal(jBytes, v); err != nil {
    return err
  }
  return nil
}

func WriteToYamlFile(instance any, dst string) {
  yamlBytes, err := yaml.Marshal(instance)
  if err != nil {
    panic(err)
  }
  yamlDbFile, err := os.Create(dst)
  if err != nil {
    panic(err)
  }
  yFileBuf := bufio.NewWriter(yamlDbFile)
  yFileBuf.Write(yamlBytes)
  yFileBuf.Flush()
  rich.Info("Writing yaml file '%s' done.", yamlDbFile.Name())
  yamlDbFile.Close()
}
