package main

import (
	"log"
	"net"
)

func main() {
	// Создаём слушатель для порта
	ln, err := net.Listen("tcp", "0.0.0.0:8081")
	if err != nil {
		log.Println(err)
	}
	defer ln.Close()

	// Принимаем входящие подключения
	conn, err := ln.Accept()
	if err != nil {
		log.Println(err)
	}
	_, err = conn.Write([]byte("message"))
	if err != nil {
		log.Println(err)
	}
}
