package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "0.0.0.0:8081")
	if err != nil {
		log.Println(err)
	}

	message := make([]byte, 1024)
	n, err := conn.Read(message)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(string(message[:n]))
	defer conn.Close()
}
