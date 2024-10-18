package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
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
	client.Join = time.Now()
	mu.Lock()
	clients[conn] = client
	mu.Unlock()

	broudcatMessage(fmt.Sprintf("User %s joined the chat", client.Username), nil)

	go getCLientMessage(client)

	for {
		select {
		case message := <-client.Message:
			broudcatMessage(fmt.Sprintf("[%s]: %s", client.Username, message), conn)
		}
	}
}
