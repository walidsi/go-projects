package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Chat represents the chat room with shared connection and mutex
type Chat struct {
	connMutex sync.Mutex
	conns     map[*websocket.Conn]struct{}
}

// NewChat creates a new chat instance
func NewChat() *Chat {
	return &Chat{
		conns: make(map[*websocket.Conn]struct{}),
	}
}

func (c *Chat) addConnection(conn *websocket.Conn) {
	c.connMutex.Lock()
	defer c.connMutex.Unlock()
	c.conns[conn] = struct{}{}
}

func (c *Chat) removeConnection(conn *websocket.Conn) {
	c.connMutex.Lock()
	defer c.connMutex.Unlock()
	delete(c.conns, conn)
}

func (c *Chat) broadcast(messageType int, p []byte) {
	c.connMutex.Lock()
	defer c.connMutex.Unlock()
	for conn := range c.conns {
		err := conn.WriteMessage(messageType, p)
		if err != nil {
			fmt.Println(err)
			c.removeConnection(conn)
		}
	}
}

// Handle WebSocket connections
func (chat *Chat) handleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	chat.addConnection(conn)
	fmt.Println("Client connected")

	// Handle messages from the client
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			chat.removeConnection(conn)
			return
		}
		fmt.Printf("Received message: %s\n", p)

		// Broadcast the message to all clients
		chat.broadcast(messageType, p)
	}
}

func main() {
	chat := NewChat()

	http.HandleFunc("/ws", chat.handleConnections)

	// Serve static files
	fs := http.FileServer(http.Dir("."))
	http.Handle("/", fs)

	// Start the server on localhost:8080
	fmt.Println("Server is running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
