package main

import (
	"fmt"
	"log"
	"net/http"
	"realtime-chat/api"
	"realtime-chat/db"
	"realtime-chat/util"
)

func main() {
	// Kết nối database (giả sử PostgreSQL/MySQL)
	db.ConnectDB()

	// Kết nối Redis
	util.InitRedis()

	// Tạo router
	r := api.SetupRouter()

	fmt.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

