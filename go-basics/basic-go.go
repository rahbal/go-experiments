package main

import (
  "fmt"
  "os"
)

func main() {
  fmt.Println("It's over", os.Args)
  if len(os.Args) != 2 {
    os.Exit(1)
  }
}
