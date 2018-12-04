package main

import (
	"bufio"
  "fmt"
  "log"
	"os"
  "strconv"
)

func main() {
  cwd, _ := os.Getwd()
  file, err := os.Open(cwd + "\\data\\1")
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  i := 0
  for scanner.Scan() {
    t := scanner.Text()
    if s, err := strconv.Atoi(t); err == nil {
      i += s
    }
  }

  fmt.Println(i)

  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }
}
