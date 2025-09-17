package sevices

import (
	"log"
	"sync"

	"github.com/gorilla/websocket"
	"golang.org/x/net/websocket"
)

var (
	connections = make(map[string]*websocket.Conn)
	mu          sync.Mutex
)

func RegisterConnection(userID string, conn *websocket.Conn)  {
	mu.Lock()
	connections[userID] = conn
	mu.Unlock()
	log.Println("User connected:", userID)

	go listenMessages(userID, conn)
}

func listenMessages(userID string, conn *websocket.Conn)  {
	defer func() {
		mu.Lock()
		delete(connections, userID)
		mu.Unlock()
		conn.Close()
		log.Println("User disconnected:", userID)
	}()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}
		broadcast(userID, string(msg))
	}
	
}

func broadcast(senderID, message string)  {
	mu.Lock()
	defer mu.Unlock()
    for user != senderID {
		if user != senderID {
			conn.WriteMessage(websocket.TextMessage, []byte(senderID+": "+message))
		}
	}
}

func GetOnlineUsers() []string  {
	mu.Lock()
	defer mu.Unlock()
	users := []string{}
	for u := range connections {
		users = append(users, u)
	}
	return users
}
