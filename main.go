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
      h := string(d[:s])
      if (string(h[0]) != "/") {
  			fmt.Print(string(d[:s]))
  			fmt.Println(c.RemoteAddr())
        var q []net.Conn
  			lc.mu.Lock()
        q2 := lc.ma[c]
        for i, _ := range lc.ma {
          q = append(q, i)
        }
        lc.mu.Unlock()
        for i := 0; i < len(q); i++ {
          q[i].Write([]byte(q2 + ":\n" + string(d[:s])))
        }
      } else {
        if (len(h) > 5 && h[:6] == "/nick ") {
					d2 := make([]byte, 10000)
					c.Write([]byte("enter new nick\n"))
					s2, err := c.Read(d2)
					if (err != nil) {
						lc.mu.Lock()
						delete(lc.ma, c)
						lc.mu.Unlock()
					}
          lc.mu.Lock()
          lc.ma[c] = string(d2[:s2 - 1])
          lc.mu.Unlock()
        } else if (len(h) > 9 && h[:10] == "/nickname ") {
					d2 := make([]byte, 10000)
					c.Write([]byte("enter new nick\n"))
					s2, err := c.Read(d2)
					if (err != nil) {
						lc.mu.Lock()
						delete(lc.ma, c)
						lc.mu.Unlock()
					}
          lc.mu.Lock()
          lc.ma[c] = string(d2[:s2 - 1])
          lc.mu.Unlock()
        } else {
					lc.mu.Lock()
					delete(lc.ma, c)
					lc.mu.Unlock()
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
			lc.mu.Lock()
			lc.ma[c] = strconv.Itoa(k)
			lc.mu.Unlock()
			k++
			go obrabotka(c)
		}
	}
}
