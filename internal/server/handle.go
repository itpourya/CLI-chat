package main

import (
	"bufio"
	"log"
	"net"
	"strings"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	log.Println("Client connected from: ", conn.RemoteAddr())

	client := Client{
		Connection: conn,
		Message:    make(chan string),
	}

	_, _ = conn.Write([]byte("Enter your Username: "))
	username, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Println("Failed to read name: ", err)
		return
	}

	client.Username = strings.TrimSpace(username)
	mu.Lock()
	clients[conn] = client
	mu.Unlock()

	go getCLientMessage(client)
}
