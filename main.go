package main

import (
  "fmt"
  "net/http"
)

func h(w http.ResponseWriter, r *http.Request) {
  fmt.Println("hello world")
}

/*var Usrs []http.ResponseWriter

func h(w http.ResponseWriter, r *http.Request) {
  for i := 0; i < len(Usrs); i++ {
    Usrs[i].Write([]byte("Hello world"))
  }
  w.Write([]byte("Hello world"))
  Usrs = append(Usrs, w)
  fmt.Println("Hello world")
}*/

func main() {
  http.HandleFunc("/", h)
  err := http.ListenAndServe(":8080", nil)
  if (err != nil) {
    fmt.Println(err)
  }
}
