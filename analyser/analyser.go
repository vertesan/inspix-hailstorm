package analyser

import (
  "bufio"
  "fmt"
  "io"
  "os"
  "regexp"
  "strings"
)

var (
  tablePtn       = regexp.MustCompile(`// Namespace: Silverflame\.SFL\n\[Table\("(?<tableName>\w+)"\)\][\s\S]+?\n}\n\n`)
  columnPtn      = regexp.MustCompile(`\[Column\("(?<columnName>\w+)"\)\][\s\S]+?public (?<type>\w+) (?<fieldName>\w+)`)
  structTemplate = "  $columnName $type `yaml:\"$columnName\"`\n"
  mapTemplate    = "    \"$lower.tsv\": $tableName{},\n"

  dumpFilePath    = "cache/dump.cs"
  structFieldPath = "cache/structs.go"
  mapFilePath     = "cache/masterMap.go"

  typeMap = map[string]string{
    "int":      "int",
    "string":   "string",
    "long":     "int64",
    "DateTime": "time.Time",
  }
)

func Analyze() {
  f, err := os.Open(dumpFilePath)
  if err != nil {
    panic(err)
  }
  defer f.Close()
  bufr := bufio.NewReader(f)
  // read dump file to string builder
  sb := new(strings.Builder)
  _, err = io.Copy(sb, bufr)
  if err != nil {
    panic(err)
  }

  // create to be generated files
  mapFile, err := os.Create(mapFilePath)
  if err != nil {
    panic(err)
  }
  defer mapFile.Close()
  mapFileBuf := bufio.NewWriter(mapFile)

  structFile, err := os.Create(structFieldPath)
  if err != nil {
    panic(err)
  }
  defer structFile.Close()
  structFileBuf := bufio.NewWriter(structFile)

  // match all classes
  contents := tablePtn.FindAllStringSubmatch(sb.String(), -1)

  // write prefixes
  mapFileBuf.WriteString("// Generated code. DO NOT EDIT!\npackage master\n\nvar (\n  MasterMap = map[string]any{\n")
  structFileBuf.WriteString("// Generated code. DO NOT EDIT!\npackage master\n\nimport \"time\"\n\n")

  for _, oneClass := range contents {
    content := oneClass[0]
    tableName := oneClass[1]
    writeMap(mapFileBuf, tableName)
    writeStruct(structFileBuf, tableName, content)
  }

  // write suffix
  mapFileBuf.WriteString("  }\n)\n")

  // flush
  err = structFileBuf.Flush()
  if err != nil {
    panic(err)
  }
  err = mapFileBuf.Flush()
  if err != nil {
    panic(err)
  }
}

func writeStruct(w *bufio.Writer, tableName string, content string) {
  // write struct prefix
  w.WriteString(fmt.Sprintf("type %v struct {\n", tableName))

  result := []byte{}
  for _, submatches := range columnPtn.FindAllStringSubmatchIndex(content, -1) {
    columnName := content[submatches[2]:submatches[3]]
    csType := content[submatches[4]:submatches[5]]
    goType := typeMap[csType]
    // result = columnPtn.ExpandString(result, structTemplate, content, submatches)
    line := structTemplate
    line = strings.Replace(line, "$columnName", columnName, -1)
    line = strings.Replace(line, "$type", goType, 1)
    result = append(result, []byte(line)...)
  }

  w.Write(result)

  // write struct suffix
  w.WriteString("}\n\n")
}

func writeMap(w *bufio.Writer, tableName string) {
  line := mapTemplate
  line = strings.Replace(line, "$lower", strings.ToLower(tableName), 1)
  line = strings.Replace(line, "$tableName", tableName, 1)
  (*w).Write([]byte(line))
}
