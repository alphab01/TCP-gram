package main

import (
//  "fmt"
  "net/http"
)

func h(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Hello world"))
}

func main() {
  http.HandleFunc("/", h)
  _ = http.ListenAndServe(":8080", nil)
}
