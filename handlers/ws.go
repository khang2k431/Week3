package handlers

import (
	"log"
	"net/http"
	"week3-chat-app/services"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true }, // allow any origin
}

func WebSocketHandler(w http.ResponseWriter, r *http.Request)  {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}
	//Get the assumed userID from the query
	userID := r.URL.Query().Get("user")
	if userID == "" {
		userID = "guest"
	}

	services.RegisterConnection(userID, conn)
}
