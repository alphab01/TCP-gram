package main

import (
  "net/http"
  "fmt"
  "io"
)

func main() {
  req, _ := http.Get("http://localhost:8080")
  defer req.Body.Close()
  dat, _ := io.ReadAll(req.Body)
  fmt.Println(string(dat))
}
