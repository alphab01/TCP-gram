package main

import (
  "fmt"
  "net"
)

func main() {
  l, err := net.Listen("tcp", "localhost:8080")
  if (err != nil) {
    fmt.Println(1)
  }
  for {
    c, err := l.Accept()
    if (err != nil) {
      fmt.Println(2)
      continue
    }
    _, err = c.Write([]byte("hello"))
    if (err != nil) {
      fmt.Println(3)
      continue
    }
    c.Close()
  }
}
