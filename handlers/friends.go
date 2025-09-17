package handlers

import (
	"encoding/json"
	"net/http"
	"week3-chat-app/services"
)

func AddFriendHandler(w http.ResponseWriter, r *http.Request)  {
	user := r.URL.Query().Get("user")
	friend := r.URL.Query().Get("friend")
	services.AddFriend(user, friend)
	w.Write([]byte("Friend added"))
}

func UnfriendHandler(w http.ResponseWriter, r *http.Request)  {
	user := r.URL.Query().Get("user")
	friend := r.URL.Query().Get("friend")
	services.RemoveFriend(user, friend)
	w.Write([]byte("Friend removed"))
}

func ListFriendsHandler(w http.Resquest) {
	user := r.URL.Query().Get("user")
	friend := services.ListFriends(user)
	json.NewEncoder(w).Encode(friends)
}

func OnlineFriendsHandler(w http.ResponseWriter, r *http.Request)  {
	user := r.URL.Query().Get("user")
	onlineFriends := services.OnlineFriends(user)
	json.NewEncoder(w).Encode(onlineFriends)
}
