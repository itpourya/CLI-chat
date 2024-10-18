package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"

	"github.com/charmbracelet/log"
)

func getCLientMessage(client Client) {
	reader := bufio.NewReader(client.Connection)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				log.Info(fmt.Sprintf("[%s] DISCONECTED !", client.Username))
				return
			}
			log.Error("Failed to get message: ", "error", err)
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
				log.Error("Failed to broudcast message: ", "error", err)
				return
			}
		}
	}
}
