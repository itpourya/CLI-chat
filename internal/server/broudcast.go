package main

import (
	"bufio"
	"log"
	"net"
	"strings"
)

func getCLientMessage(client Client) {
	reader := bufio.NewReader(client.Connection)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			log.Println("Failed to get message: ", err)
			return
		}

		message = strings.TrimSpace(message)
		if message != "" && message != " " {
			client.Message <- message
		}
	}
}

func broudcatMessage(message string, sender net.Conn) {
	mu.Lock()
	defer mu.Unlock()

	for conn, _ := range clients {
		if conn != sender {
			_, err := conn.Write([]byte(message + "\n"))
			if err != nil {
				log.Println("Failed to broudcast message: ", err)
				return
			}
		}
	}
}
