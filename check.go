package main

import (
	"fmt"
	"net"
)

func obrabotka(c net.Conn) {
	_, err := c.Write([]byte("hello"))
	if err != nil {
		fmt.Println(3)
		return
	}
  for {
    var d []byte
    _, err := c.Read(d)
    if (err != nil) {
      return
    }
		if (len(string(d)) > 0) {
			fmt.Println(string(d))
		}
  }
}

func main() {
	l, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println(1)
		l.Close()
		return
	}
	defer l.Close()
	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(2)
		} else {
			go obrabotka(c)
		}
	}
}
