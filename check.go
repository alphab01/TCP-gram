package main

import (
	"fmt"
	"net"
	"strconv"
)

var lc []net.Conn

func obrabotka(c net.Conn) {
	_, err := c.Write([]byte("hello"))
	if err != nil {
		fmt.Println(3)
		return
	}
  for {
    d := make([]byte, 10000)
    _, err := c.Read(d)
    if (err != nil) {
      return
    }
		if (len(string(d)) > 0) {
			fmt.Print(string(d))
			fmt.Println(c.RemoteAddr())
			q := 0
			for i := 0; i < len(lc); i++ {
				if (c == lc[i]) {
					q = i
					break
				}
			}
			for i := 0; i < len(lc); i++ {
				lc[i].Write([]byte(strconv.Itoa(q) + ":\n" + string(d)))
			}
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
			lc = append(lc, c)
			go obrabotka(c)
		}
	}
}
