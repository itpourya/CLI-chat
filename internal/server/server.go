package main

import (
	"log"
	"net"
	"os"
	"sync"
	"time"

	"github.com/joho/godotenv"
)

type Client struct {
	Username   string
	Connection net.Conn
	Message    chan string
	Join       time.Time
}

var (
	clients  = make(map[net.Conn]Client)
	mu       = sync.Mutex{}
	_        = godotenv.Load("../../.env")
	hostType = os.Getenv("SERVER_TYPE")
	addr     = os.Getenv("SERVER_HOST")
	port     = os.Getenv("SERVER_PORT")
)

func main() {
	listener, err := net.Listen(hostType, addr+":"+port)
	if err != nil {
		log.Fatal("Failed to listen on ", addr+":"+port)
	}
	defer listener.Close()

	// Server start
	log.Println("Server Started. Listening on: " + addr + ":" + port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Failed accepting connections")
		}

		go handleConnection(conn)
	}
}
