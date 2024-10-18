package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/charmbracelet/log"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	log.Info("Client connected from:" + conn.RemoteAddr().String())

	client := Client{
		Connection: conn,
		Message:    make(chan string),
	}

	_, _ = conn.Write([]byte("Enter your Username: "))
	username, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Error("Failed to read name: ", "error", err)
		return
	}

	client.Username = strings.TrimSpace(username)
	client.Join = time.Now()
	mu.Lock()
	clients[conn] = client
	mu.Unlock()

	broudcatMessage(fmt.Sprintf("\n"+"User %s joined the chat", client.Username), nil)

	go getCLientMessage(client)

	for {
		select {
		case message := <-client.Message:
			broudcatMessage(fmt.Sprintf("[%s]: %s", client.Username, message), conn)
		}
	}
}
