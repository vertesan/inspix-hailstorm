package manifest

import (
	"bufio"
	"bytes"
	"crypto/md5"
	"encoding/binary"
	"io"
	"log"
	"reflect"
	"slices"

	"vertesan/hailstorm/crypto"
	"vertesan/hailstorm/rich"
	"vertesan/hailstorm/utils"
)

type Entry struct {
  Priority           uint32
  ResourceType       uint32
  NumDeps            uint32
  NumContents        uint32
  NumCategories      uint32
  Size               uint64
  ContentTypeCrcs    []byte
  TypeCrc            uint32
  CategoryCrcs       []byte
  LabelCrc           uint64
  ContentNameCrcs    []byte
  DepCrcs            []byte
  RecDepCrcs         []byte
  NumRecDepCrcs      uint32
  Checksum           uint64
  Seed               uint64
  StrTypeCrc         string
  StrContentTypeCrcs []string
  StrCategoryCrcs    []string
  StrLabelCrc        string
  StrContentNameCrcs []string
  StrDepCrcs         []string
  RealName           string
}

type Catalog struct {
  Entries []Entry
}

func (c *Catalog) Init(mani *Manifest, src io.Reader) {
  buf := new(bytes.Buffer)
  revMap := make(map[uint64]int)
  c.Parse(mani, src, buf)
  c.ParseTransposedArray(buf, revMap)
  c.ResolveAllDeps(revMap)
  c.ResolveAllRealNames()
}

func (c *Catalog) Parse(mani *Manifest, src io.Reader, dst io.Writer) {
  DecodeAsset(&mani.Asset, dst, src)
}

func (c *Catalog) ParseTransposedArray(bBuf *bytes.Buffer, revMap map[uint64]int) {
  bBuff := bufio.NewReader(bBuf)

  var magic uint16
  binary.Read(bBuff, binary.BigEndian, &magic)
  if magic != 0xCA01 {
    rich.Panic("First 2 bytes are [%X], expect [CA01]", magic)
  }

  var nothing uint16
  binary.Read(bBuff, binary.BigEndian, &nothing)

  var _size uint64
  _size, err := binary.ReadUvarint(bBuff)
  if err != nil {
    panic(err)
  }
  size := int(_size)

  c.Entries = make([]Entry, size)

  // iterate all
  itrAll := func(fn func(int)) {
    for i := 0; i < size; i++ {
      fn(i)
    }
  }

  // iterate reading vlq
  itrReadUvarint := func(field string) {
    itrAll(func(i int) {
      v, err := binary.ReadUvarint(bBuff)
      if err != nil {
        panic(err)
      }
      reflect.ValueOf(&c.Entries[i]).Elem().FieldByName(field).SetUint(v)
    })
  }

  itrReadUvarint("Priority")
  itrReadUvarint("ResourceType")
  itrReadUvarint("NumDeps")
  itrReadUvarint("NumContents")
  itrReadUvarint("NumCategories")
  itrReadUvarint("Size")

  // TypeCrc
  itrAll(func(i int) {
    rawBytes := utils.ReadUntilNext(bBuff, 0x00)
    c.Entries[i].TypeCrc = crypto.UpdateCrc32(0, rawBytes, len(rawBytes))
    c.Entries[i].StrTypeCrc = string(rawBytes)
  })

  // ContentTypeCrcs
  itrAll(func(i int) {
    entry := &c.Entries[i]
    n := entry.NumContents
    for n >= 1 {
      rawBytes := utils.ReadUntilNext(bBuff, 0x00)
      crc32 := crypto.UpdateCrc32(0, rawBytes, len(rawBytes))
      bs := make([]byte, 4)
      binary.BigEndian.PutUint32(bs, crc32)
      entry.ContentTypeCrcs = append(entry.ContentTypeCrcs, bs...)
      entry.StrContentTypeCrcs = append(entry.StrContentTypeCrcs, string(rawBytes))
      n--
    }
  })

  // CategoryCrcs
  itrAll(func(i int) {
    entry := &c.Entries[i]
    n := entry.NumCategories
    for n >= 1 {
      rawBytes := utils.ReadUntilNext(bBuff, 0x00)
      crc32 := crypto.UpdateCrc32(0, rawBytes, len(rawBytes))
      bs := make([]byte, 4)
      binary.BigEndian.PutUint32(bs, crc32)
      entry.CategoryCrcs = append(entry.CategoryCrcs, bs...)
      entry.StrCategoryCrcs = append(entry.StrCategoryCrcs, string(rawBytes))
      n--
    }
  })

  // LabelCrc & revMap
  itrAll(func(i int) {
    rawBytes := utils.ReadUntilNext(bBuff, 0x00)
    crc64 := crypto.UpdateCrc64(0, rawBytes, len(rawBytes), nil)
    c.Entries[i].LabelCrc = crc64
    c.Entries[i].StrLabelCrc = string(rawBytes)
    revMap[crc64] = i
  })

  // ContentNameCrcs
  itrAll(func(i int) {
    entry := &c.Entries[i]
    n := entry.NumContents
    for n >= 1 {
      rawBytes := utils.ReadUntilNext(bBuff, 0x00)
      crc64 := crypto.UpdateCrc64(0, rawBytes, len(rawBytes), nil)
      bs := make([]byte, 8)
      binary.BigEndian.PutUint64(bs, crc64)
      entry.ContentNameCrcs = append(entry.ContentNameCrcs, bs...)
      entry.StrContentNameCrcs = append(entry.StrContentNameCrcs, string(rawBytes))
      n--
    }
  })

  // DepCrcs
  itrAll(func(i int) {
    entry := &c.Entries[i]
    n := entry.NumDeps
    for n >= 1 {
      rawBytes := utils.ReadUntilNext(bBuff, 0x00)
      crc64 := crypto.UpdateCrc64(0, rawBytes, len(rawBytes), nil)
      bs := make([]byte, 8)
      binary.BigEndian.PutUint64(bs, crc64)
      entry.DepCrcs = append(entry.DepCrcs, bs...)
      entry.StrDepCrcs = append(entry.StrDepCrcs, string(rawBytes))
      n--
    }
  })

  // CheckSum
  itrAll(func(i int) {
    var checkSum uint64
    binary.Read(bBuff, binary.BigEndian, &checkSum)
    c.Entries[i].Checksum = checkSum
  })

  // Seed
  itrAll(func(i int) {
    peek, err := bBuff.Peek(1)
    if err != nil {
      panic(err)
      // if err == io.EOF && i >= len(c.Entries)-4 {
      //   peek = []byte{0x00}
      // } else {
      //   panic(err)
      // }
    }

    var seed uint64
    if uint32(peek[0]) != 0 {
      binary.Read(bBuff, binary.BigEndian, &seed)
    } else {
      _seed, err := bBuff.ReadByte()
      if err != nil {
        panic(err)
      }
      seed = uint64(_seed)
    }
    c.Entries[i].Seed = seed
  })

  _, err = bBuff.ReadByte()
  if err != io.EOF {
    log.Panic("Remaining bytes exist after ParseTransposedArray, perhaps the file is broken.")
  }
}

func (c *Catalog) ResolveAllDeps(revMap map[uint64]int) {
  maxLen := len(c.Entries)

  var resolveRecursiveDeps func(label uint64, ordDict *[]uint64)
  resolveRecursiveDeps = func(label uint64, ordDict *[]uint64) {
    if slices.Contains(*ordDict, label) {
      return
    }
    *ordDict = append(*ordDict, label)
    revIdx, ok := revMap[label]
    if !ok {
      log.Panicf("Label: %q not inside revMap", label)
    }
    revIdxEntry := &c.Entries[revIdx]
    numDeps := revIdxEntry.NumDeps
    for i := 0; i < int(numDeps); i++ {
      depCrcs := binary.BigEndian.Uint64(revIdxEntry.DepCrcs[i*8:])
      resolveRecursiveDeps(depCrcs, ordDict)
    }
  }

  for i := 0; i < maxLen; i++ {
    entry := &c.Entries[i]
    buf := new(bytes.Buffer)
    var ordDict []uint64
    resolveRecursiveDeps(entry.LabelCrc, &ordDict)
    binary.Write(buf, binary.BigEndian, ordDict)
    entry.RecDepCrcs = buf.Bytes()
    if len(c.Entries) > 16396 {
      log.Default()
    }
  }
}

func (c *Catalog) ResolveAllRealNames() {
  for idx := range c.Entries {
    entry := &c.Entries[idx]
    entry.RealName = GetRealName(entry.Checksum, entry.LabelCrc, entry.Size)
  }
}

func GetRealName(checksum uint64, labelCrc uint64, size uint64) string {
  buf := new(bytes.Buffer)
  binary.Write(buf, binary.BigEndian, checksum)
  binary.Write(buf, binary.BigEndian, labelCrc)

  varuintSlice := make([]byte, binary.MaxVarintLen64)
  n := binary.PutUvarint(varuintSlice, size)
  buf.Write(varuintSlice[:n])
  md5 := md5.Sum(buf.Bytes())

  return crypto.Base32Encoder.EncodeToString(md5[:])
}
