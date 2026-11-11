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
		lc.mu.Lock()
		delete(lc.ma, c)
		lc.mu.Unlock()
		return
	}
	for {
		d := make([]byte, 10000)
		s, err := c.Read(d)
		if err != nil {
			lc.mu.Lock()
      delete(lc.ma, c)
			lc.mu.Unlock()
			return
		}
		if len(string(d[:s])) > 0 {
			fmt.Print(string(d[:s]))
			fmt.Println(c.RemoteAddr())
      var q []net.Conn
			lc.mu.Lock()
      for i, _ := range lc.ma {
        q = append(q, i)
      }
      lc.mu.Unlock()
      for i := 0; i < len(q); i++ {
        q[i].Write([]byte(lc.ma[c] + ":\n" + string(d[:s])))
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
			lc.mu.Lock()
			lc.ma[c] = strconv.Itoa(k)
			lc.mu.Unlock()
			k++
			go obrabotka(c)
		}
	}
}
