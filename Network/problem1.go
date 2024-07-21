package main

import (
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8081")
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()
	message := make([]byte, 1024)
	for i := 0; i < 3; i++ {
		n, err := conn.Read(message)
		if err != nil {
			log.Println(err)
		}
		fmt.Println(strings.ToUpper(string(message[:n])))
	}
}
