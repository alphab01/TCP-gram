package main

import (
	"fmt"
	"net"
	"strconv"
	"sync"
)

type protectmap struct {
  ma map[net.Conn]string
  mu sync.Mutex
}

var lc protectmap

var k int

func obrabotka(c net.Conn) {
	defer c.Close()
	_, err := c.Write([]byte("hello"))
	if err != nil {
    lc[c] = "err"
		return
	}
  for {
    d := make([]byte, 10000)
    s, err := c.Read(d)
    if (err != nil) {
      lc[c] = "err"
			return
    }
		if (len(string(d[:s])) > 0) {
			fmt.Print(string(d[:s]))
			fmt.Println(c.RemoteAddr())
			for i, m := range lc {
        if (m != "err") {
  	    	i.Write([]byte(lc[c] + ":\n" + string(d[:s])))
        }
			}
		}
  }
}

func main() {
	lc = protectmap{ma: make(map[net.Conn]string)}
  k = 0
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
      lc.ma[c] = strconv.Itoa(k)
      k++
			go obrabotka(c)
		}
	}
}
