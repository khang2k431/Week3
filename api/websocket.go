package api

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"realtime-chat/service"
	"realtime-chat/util"
	"sync"
)

// Nâng cấp HTTP -> WebSocket
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

var (
	clients   = make(map[*websocket.Conn]bool) // danh sách client kết nối
	broadcast = make(chan string)              // channel để phát tin nhắn
	lock      sync.Mutex
)

func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "WebSocket upgrade failed", http.StatusInternalServerError)
		return
	}

	lock.Lock()
	clients[conn] = true
	lock.Unlock()

	// Đánh dấu user online trong Redis
	util.SetUserOnline(conn.RemoteAddr().String())

	go readMessages(conn)
	go writeMessages()
}

func readMessages(conn *websocket.Conn) {
	defer func() {
		lock.Lock()
		delete(clients, conn)
		lock.Unlock()
		conn.Close()
		util.SetUserOffline(conn.RemoteAddr().String())
	}()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}

		// Check rate limit
		if !util.AllowRequest(conn.RemoteAddr().String()) {
			conn.WriteMessage(websocket.TextMessage, []byte("Too many messages, slow down!"))
			continue
		}

		// Lưu vào Redis (message history)
		service.SaveMessage(string(msg))

		// Gửi qua channel
		broadcast <- fmt.Sprintf("%s: %s", conn.RemoteAddr().String(), msg)
	}
}

func writeMessages() {
	for {
		msg := <-broadcast
		lock.Lock()
		for client := range clients {
			client.WriteMessage(websocket.TextMessage, []byte(msg))
		}
		lock.Unlock()
	}
}
