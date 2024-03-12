package rich

import (
  "fmt"

  "github.com/fatih/color"
)

func Info(text string, a ...any) {
  fmt.Println(color.BlueString(">>> [Info]"), fmt.Sprintf(text, a...))
}

func Error(text string, a ...any) {
  fmt.Println(color.RedString(">>> [Error]"), fmt.Sprintf(text, a...))
}

func Warning(text string, a ...any) {
  fmt.Println(color.YellowString(">>> [Warning]"), fmt.Sprintf(text, a...))
}

func Panic(text string, a ...any) {
  Error(text, a...)
  panic("Exiting program due to the aforementioned reasons.")
}

func PanicError(text string, err error, a ...any) {
  Error(text, a...)
  panic(err)
}
