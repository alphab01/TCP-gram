package main

import (
	"fmt"
	"net"
)

func obrabotka(l net.Listener) {
	c, err := l.Accept()
	if err != nil {
		fmt.Println(2)
		return
	}
	_, err = c.Write([]byte("hello"))
	if err != nil {
		fmt.Println(3)
		return
	}
	c.Close()
}

func main() {
	l, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println(1)
	}
	for {
		go obrabotka(l)
	}
}
