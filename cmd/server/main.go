package main

import (
	"log"
	"net/http"
	"week3-chat-app/internal/handlers"
	"week3-chat-app/internal/pubsub"
	"week3-chat-app/internal/services"
	"week3-chat-app/internal/storage"
)

func main() {
	// Init Redis
	storage.InitRedis()

	// Subscribe Pub/Sub
	pubsub.SubscribeMessages(func(msg string) {
		services.BroadcastFromPubSub(msg)
	})

	// Routes
	http.HandleFunc("/ws", handlers.WebSocketHandler)
	http.HandleFunc("/friends/add", handlers.AddFriendHandler)
	http.HandleFunc("/friends/remove", handlers.UnfriendHandler)
	http.HandleFunc("/friends/list", handlers.ListFriendsHandler)
	http.HandleFunc("/friends/online", handlers.OnlineFriendsHandler)

	log.Println("Server running at :8080")
	http.ListenAndServe(":8080", nil)
}
