package master

import (
  "bufio"
  "encoding/binary"
  "io"
  "reflect"
  "time"
  "unsafe"

  "vertesan/hailstorm/rich"
  "vertesan/hailstorm/utils"
)

func Parse[T any](src io.Reader, label string, instance *T) []T {
  bufr := bufio.NewReader(src)

  buf := make([]byte, 2)
  // magic number 0xDA00
  if _, err := io.ReadFull(bufr, buf); err != nil {
    panic(err)
  }
  if buf[0] != 0xDA || buf[1] != 0x00 {
    rich.Panic("Magic number mismatched, except: 0xDA00, given: %X.", buf)
  }

  // idk what is this
  if _, err := io.ReadFull(bufr, buf); err != nil {
    panic(err)
  }

  // vlq: numRows
  rowNum, err := binary.ReadUvarint(bufr)
  if err != nil {
    panic(err)
  }
  if rowNum < 1 {
    rich.Info("Database file %q has 0 rows.", label)
    return nil
  }

  // vlq: numFields
  fieldNum, err := binary.ReadUvarint(bufr)
  if err != nil {
    panic(err)
  }
  if fieldNum < 1 {
    rich.Warning("Table %q has no fields.", label)
    return nil
  }
  stNum := uint64(reflect.TypeOf(*instance).NumField())
  if stNum < fieldNum {
    rich.Warning("Incoming table %q has %d fields, but struct has %d fields. Perhaps the DB structure was changed.", label, fieldNum, stNum)
    rich.Warning("struct fields num %d will be used to avoid 'IndexOutOfBound' panic.", stNum)
    fieldNum = stNum
  }

  var fieldNames []uint32
  var fieldTypes []uint32

  for i := 0; i < int(fieldNum); i++ {
    buf4b := make([]byte, 4)

    // uint32: fieldNames
    if _, err = io.ReadFull(bufr, buf4b); err != nil {
      panic(err)
    }
    fn := binary.BigEndian.Uint32(buf4b)
    fieldNames = append(fieldNames, fn)

    // uint32: fieldTyps
    if _, err = io.ReadFull(bufr, buf4b); err != nil {
      panic(err)
    }
    ty := binary.BigEndian.Uint32(buf4b)
    fieldTypes = append(fieldTypes, ty)
  }

  // Since T is interface{} here, we cannot simple make([]T, rowNum).
  // Instead, assign a shallow copy of the instance to every items
  // inside the slice to make sure those items are not assigned to nil.
  results := make([]T, rowNum)
  for i := range results {
    results[i] = *instance
  }

  for fieldIdx := 0; fieldIdx < int(fieldNum); fieldIdx++ {
    // fieldCrc := fieldNames[i]
    givenType := fieldTypes[fieldIdx]
    // field := GetField(label, fieldIdx)

    // fieldName := reflect.ValueOf(*instance).Field(fieldIdx).Type().Name()
    fieldName := reflect.TypeOf(*instance).Field(fieldIdx).Name
    field := Field{
      Name: fieldName,
      Type: givenType,
    }

    if !ensureTypeOrSkipRows(field.Type, givenType) {
      rich.Warning("Missmatched types, skip reading field %q.", field.Name)
      rich.Warning("Expect: %X, given: %X.", field.Type, givenType)
      continue
    }

    // I don't think it can be larger than MaxInt32 by any chance...
    for rowIdx := 0; rowIdx < int(rowNum); rowIdx++ {
      read1Cell(&results[rowIdx], bufr, field.Name, givenType)
    }
  }
  if _, err = bufr.ReadByte(); err != io.EOF {
    rich.Warning("Except io.EOF but redundant bytes are detected during parsing tsv: %q.", label)
    // rich.PanicError("Except io.EOF but redundant bytes are detected during parsing tsv: %q.", err, label)
  }
  rich.Info("Database file %q was successfully parsed.", label)
  return results
}

func read1Cell[T any](instance *T, r *bufio.Reader, field string, givenType uint32) {
  // get Type info
  typeField, found := reflect.TypeOf(*instance).FieldByName(field)
  if !found {
    rich.Panic("Field name %q not found.", field)
  }

  switch givenType {
  case 0x10: // string, DateTime
    buf := utils.ReadUntilNext(r, 0x00)
    switch typeField.Type.Name() {
    case "string":
      setValueToInterface(instance, field, reflect.ValueOf(string(buf)))
    case "Time":
      t, err := time.Parse(time.DateTime, string(buf))
      if err != nil {
        panic(err)
      }
      setValueToInterface(instance, field, reflect.ValueOf(t))
    default:
      rich.Panic("Unimplemented type: %q of %q found.", typeField.Type.Name(), field)
    }
  case 0x20: // int
    uNum, err := binary.ReadUvarint(r)
    if err != nil {
      panic(err)
    }
    excepted := typeField.Type.Name()
    if excepted == "int" {
      signedInt32 := *(*int32)(unsafe.Pointer(&uNum))
      setValueToInterface(instance, field, reflect.ValueOf(int(signedInt32)))
    } else if excepted == "int64" { // not sure if this is needed or not
      setValueToInterface(instance, field, reflect.ValueOf(int64(uNum)))
    } else {
      rich.Panic("Unimplemented excepted type %q.", excepted)
    }
  case 0x33: // long
    numBuf := make([]byte, 8)
    if _, err := io.ReadFull(r, numBuf); err != nil {
      panic(err)
    }
    uNum := binary.BigEndian.Uint64(numBuf)
    setValueToInterface(instance, field, reflect.ValueOf(int64(uNum)))
  default:
    rich.Panic("Unkonwn data type [%X] is detected in %q.", givenType, reflect.TypeOf(*instance).Name())
  }
}

func ensureTypeOrSkipRows(except uint32, given uint32) bool {
  if except == given || except == 0x33 && given == 0x20 {
    return true
  }
  return false
}

// https://stackoverflow.com/questions/63421976/panic-reflect-call-of-reflect-value-fieldbyname-on-interface-value#63422049
func setValueToInterface[T any](instance *T, fieldName string, value reflect.Value) {
  // v is the interface{}
  v := reflect.ValueOf(instance).Elem()
  // Allocate a temporary variable with type of the struct.
  // v.Elem() is the true value contained in the interface.
  tmp := reflect.New(v.Elem().Type()).Elem()
  // Copy the struct value contained in interface to
  // the temporary variable.
  tmp.Set(v.Elem())
  // Now we can set the field.
  tmp.FieldByName(fieldName).Set(value)
  // Set the interface back to the modified struct value.
  v.Set(tmp)
}
