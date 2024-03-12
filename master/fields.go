package master

type Field struct {
  Name string
  Type uint32
}

// func GetField(label string, idx int) *Field {
//   return &fieldNamesMapping[label][idx]
// }

// var fieldNamesMapping = map[string][]Field{
//   "cardrarities.tsv": {
//     {"Id", 0x20},
//     {"RarityName", 0x10},
//   },
// }
