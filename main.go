package main

import (
//  "fmt"
  "net/http"
)

var Usrs []http.ResponseWriter

func h(w http.ResponseWriter, r *http.Request) {
  for i := 0; i < len(Usrs); i++ {
    Usrs[i].Write([]byte("Hello world"))
  }
  Usrs = append(Usrs, w)
}

func main() {
  http.HandleFunc("/", h)
  _ = http.ListenAndServe(":8080", nil)
}
