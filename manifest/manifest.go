package manifest

import (
  "bytes"
  "encoding/base64"
  "encoding/binary"
  "strings"

  "vertesan/hailstorm/crypto"
  "vertesan/hailstorm/rich"
)

type Manifest struct {
  Asset
  SimpleResver  string
  Checksum      uint64
  LabelCrc      uint64
  ClientVersion string
}

// resHeader: something like `R2402010@B/FicABV0d3BUb8PQHvXSsDwHw==`
func (m *Manifest) Init(resHeader string, clientVer string) {
  m.Type = CATALOG
  m.ClientVersion = clientVer
  splitted := strings.Split(resHeader, "@")
  if len(splitted) != 2 {
    rich.Panic("resHeader %q cannot be splitted by `@`", resHeader)
  }
  m.SimpleResver = splitted[0]
  decoded, err := base64.StdEncoding.DecodeString(splitted[1])
  if err != nil {
    panic(err)
  }
  reader := bytes.NewReader(decoded)

  err = binary.Read(reader, binary.BigEndian, &m.Checksum)
  if err != nil {
    panic(err)
  }

  err = binary.Read(reader, binary.BigEndian, &m.Seed)
  if err != nil {
    panic(err)
  }

  m.Size, err = binary.ReadUvarint(reader)
  if err != nil {
    panic(err)
  }

  m.LabelCrc = crypto.UpdateCrc64(0, []byte(m.SimpleResver), len(m.SimpleResver), nil)

  m.RealName = GetRealName(m.Checksum, m.LabelCrc, m.Size)

  m.CalcCrc64Name = m.ClientVersion + ":" + m.SimpleResver

  rich.Info("Initialize manifest completed.")
  rich.Info("Manifest realname: %q.", m.RealName)
}
