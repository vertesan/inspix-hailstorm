package manifest

import (
  "bufio"
  "bytes"
  "crypto/sha256"
  "encoding/binary"
  "encoding/hex"
  "io"
  "log"
  "os"

  "vertesan/hailstorm/crypto"
  "vertesan/hailstorm/rich"

  "github.com/pierrec/lz4/v4"
)

const (
  RAW AssetType = 1 + iota
  CATALOG
  RAW_STR     = "raw"
  CATALOG_STR = "android"
  PREFIX      = "c34ea77df4976cd8907096fa47bb97e61852305892494e3692ba0c7eb434f022c549c96cf7ca0ee1b6ba7f203b6c76e8679699ce9c44af7b1cb000173a515938"
)

type AssetType int

type Asset struct {
  Seed          uint64
  Size          uint64
  RealName      string
  CalcCrc64Name string
  Type          AssetType
  // Md5           [16]byte
}

func DecryptAllAssets(catalog *Catalog, dstDir string, srcDir string) {
  counter := 0
  amount := len(catalog.Entries)

  // decrypt (or just copy) asset files
  if err := os.MkdirAll(dstDir, 0755); err != nil {
    panic(err)
  }
  for _, entry := range catalog.Entries {
    counter++
    var suffix = ""
    if entry.ResourceType == 1 {
      suffix = ".assetbundle"
    }
    // check existing
    // if _, err := os.Stat(dstDir + "/" + entry.StrLabelCrc + suffix); err == nil {
    //   rich.Info("(%d/%d) Asset file %q(%v) already exists, skip decrypting.", counter, amount, entry.StrLabelCrc, entry.RealName)
    //   continue
    // }
    
    rawFile, err := os.Open(srcDir + "/" + entry.RealName)
    if err != nil {
      panic(err)
    }
    rawBuf := bufio.NewReader(rawFile)
    plainFile, err := os.Create(dstDir + "/" + entry.StrLabelCrc + suffix)
    if err != nil {
      panic(err)
    }
    plainBuf := bufio.NewWriter(plainFile)

    switch entry.ResourceType {
    case 0, 128: // do nothing
      if _, err := io.Copy(plainBuf, rawBuf); err != nil {
        panic(err)
      }
    case 1: // skip first 2 bytes
      offset, err := io.ReadFull(rawBuf, make([]byte, 2))
      if err != nil {
        panic(err)
      }
      if offset != 2 {
        rich.Panic("Failed to seek asset file %q.", rawFile.Name())
      }
      signature, err := rawBuf.Peek(7)
      if err != nil {
        panic(err)
      }
      if string(signature) != "UnityFS" {
        rich.Panic("AssetBundle signature mismatches for %q.", entry.StrLabelCrc)
      }
      if _, err := io.Copy(plainBuf, rawBuf); err != nil {
        panic(err)
      }
    case 192: // need decrypting
      asset := &Asset{
        Seed:          entry.Seed,
        Size:          entry.Size,
        RealName:      entry.RealName,
        CalcCrc64Name: entry.StrLabelCrc,
        Type:          RAW,
      }
      DecodeAsset(asset, plainBuf, rawBuf)
    }
    plainBuf.Flush()
    rich.Info("(%d/%d) Asset file %q(%v) was successfully processed.", counter, amount, entry.StrLabelCrc, entry.RealName)
    rawFile.Close()
    plainFile.Close()
  }
  rich.Info("All asset files processed.")
}

func DecodeAsset(asset *Asset, dst io.Writer, src io.Reader) {
  hexPrefix, _ := hex.DecodeString(PREFIX)

  buf := new(bytes.Buffer)

  // Step 1. Write a bunch of constant arbitrary hex data
  buf.Write(hexPrefix)

  // Step 2. Write Manifest.Seed
  var seed uint64
  switch asset.Type {
  case RAW:
    seed = asset.Seed & 0x7FFFFFFFFFFFFFFF
  case CATALOG:
    seed = asset.Seed
  }
  binary.Write(buf, binary.BigEndian, seed)

  // Step 3. Write CRC64 of crc64str
  crc64str := asset.CalcCrc64Name
  crc64 := crypto.UpdateCrc64(0, []byte(crc64str), len(crc64str), nil)
  binary.Write(buf, binary.BigEndian, crc64)

  // Step 4. Write CRC32
  var crc32Str string
  switch asset.Type {
  case RAW:
    crc32Str = RAW_STR
  case CATALOG:
    crc32Str = CATALOG_STR
  default:
    log.Panicf("Unkown asset type: %q", asset.Type)
  }
  crc32 := crypto.UpdateCrc32(0, []byte(crc32Str), len(crc32Str))
  binary.Write(buf, binary.BigEndian, crc32)

  // Step 5. Write VLQ of Manifest.Size
  vlqBuf := make([]byte, binary.MaxVarintLen64)
  n := binary.PutUvarint(vlqBuf, asset.Size)
  buf.Write(vlqBuf[:n])

  // Compute sha256 of all bytes
  keyiv := sha256.Sum256(buf.Bytes())

  key := keyiv[:16]
  iv := keyiv[16:]

  internalBuf := new(bytes.Buffer)
  // middleRwter := bufio.NewReadWriter(
  //   bufio.NewReader(internalBuf),
  //   bufio.NewWriter(internalBuf),
  // )

  // Decrypt
  crypto.Decrypt(key, iv, src, internalBuf)

  // Decompress
  lz4Reader := lz4.NewReader(internalBuf)
  _, err := io.Copy(dst, lz4Reader)
  if err != nil {
    panic(err)
  }
}
