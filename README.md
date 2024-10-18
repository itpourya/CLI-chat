# TCP Chat Server

A simple TCP chat server written in Go that allows multiple clients to connect and chat over a TCP network connection. This project demonstrates how to set up a basic chat application using the TCP protocol for communication.

## Features
- **Multiple Clients**: Supports multiple clients connecting to the server.
- **Real-time Chat**: Clients can chat in real-time over the TCP layer.
- **Golang Implementation**: Built entirely using Go's standard `net` package.
- **Localhost Communication**: Easily test on a local server before deploying.

## Requirements

- Go 1.16+ installed on your machine.
- Basic understanding of TCP networking.

## Installation

To install and run the TCP chat server on your local machine, follow these steps:

1. Clone the repository:
   ```bash
   git clone https://github.com/itpourya/CLI-chat.git
   ```

2. Navigate to the project directory:
   ```bash
   cd CLI-chat
   ```

3. Run the server:
   ```bash
   go run internal/server/server.go
   ```

4. Run the client(s) (in separate terminals):
   ```bash
   go run internal/client/client.go
   ```

## Usage

### Running the Server

To start the server, simply execute:

```bash
go run server.go
```

The server will start and listen for incoming client connections on `localhost` and the specified port (default is `8080`).

### Connecting Clients

Clients can connect to the server by running:

```bash
go run client.go
```

Once connected, the clients will be able to send and receive messages in real-time. Any message sent by a client will be broadcast to all connected clients.

### Example Output

When the server is running, you should see messages like this:

```bash
Client connected from: 127.0.0.1:54321
Broadcasting message: Hello, everyone!
Client disconnected: 127.0.0.1:54321
```

And on the client side, you'll see:

```bash
Connected to the server at localhost:8080
Enter your message:
> Hello, everyone!
Message from server: Hello, everyone!
```

## Project Structure

```
.
├── client.go       # Client implementation
├── server.go       # Server implementation
├── README.md       # Project documentation
```

## How It Works

- **Server**: The server listens on a specified TCP port for incoming client connections. Each connected client is handled in a separate goroutine, allowing multiple clients to communicate simultaneously. The server broadcasts all messages received from clients to all other connected clients.

- **Client**: The client establishes a TCP connection to the server and continuously listens for incoming messages while also allowing the user to send messages to the server.

## Configuration

You can change the default server port and host by modifying the constants in the `.env` file:

```bash
    SERVER_HOST = "localhost"
    SERVER_PORT = "8080"
    SERVER_TYPE = "tcp"
```

## Contributing

Feel free to submit pull requests, issues, or suggestions. Contributions are welcome!

1. Fork the repository
2. Create a new branch (`git checkout -b feature-branch`)
3. Make your changes and commit (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin feature-branch`)
5. Create a new Pull Request

## Contact

For any inquiries or feedback, feel free to contact me via [itpourya@yahoo.com](mailto:itpourya@yahoo.com).

---

Feel free to adjust it according to your project specifics!
