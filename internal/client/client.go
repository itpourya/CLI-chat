package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"sync"

	"github.com/charmbracelet/log"

	"github.com/joho/godotenv"
)

var (
	mu       = sync.Mutex{}
	_        = godotenv.Load("../../.env")
	hostType = os.Getenv("SERVER_TYPE")
	addr     = os.Getenv("SERVER_HOST")
	port     = os.Getenv("SERVER_PORT")
)

func main() {
	conn, err := net.Dial(hostType, addr+":"+port)
	if err != nil {
		log.Fatal("Failed to connect on ", addr+":"+port)
	}
	defer conn.Close()

	log.Info("Connected to the server at ", addr+":"+port)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter your username: ")
	scanner.Scan()
	username := scanner.Text()

	_, err = conn.Write([]byte(username + "\n"))
	if err != nil {
		log.Error("Failed to send username ", "error", err)
		return
	}

	go reciverMessage(conn)

	for scanner.Scan() {
		message := scanner.Text()
		_, err := conn.Write([]byte(message + "\n"))
		if err != nil {
			log.Error("Failed to send message ", "error", err)
			break
		}
	}

}

func reciverMessage(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			log.Error("Failed to read message ", "error", err)
			return
		}

		fmt.Println(message)
	}
}
