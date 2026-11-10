package main

import (
  "fmt"
  "net"
)

func main() {
  for {
    l, err := net.Listen("tcp", "localhost:8080")
    if (err != nil) {
      fmt.Println(1)
    }
    c, err := l.Accept()
    if (err != nil) {
      fmt.Println(2)
    }
    _, err = c.Write([]byte("hello"))
    if (err != nil) {
      fmt.Println(3)
    }
  }
}
