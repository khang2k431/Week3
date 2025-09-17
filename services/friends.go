package services

import "sync"

var (
	friendships = make(map[string][]string) // userID -> list friends
	fmu         sync.Mutex
)

func AddFriend(user, friend string)  {
	fmu.Lock()
	defer fmu.Unlock()
	friendships[user] = append(friendships[user], friend)

	// Notify friend request via WebSocket if online
	if conn, ok := connections[friend]; ok {
		conn.WriteMessage(1, []byte("New friend request from: "+user))
	}
}

func RemoveFriend(user, friend string) {
	fmu.Lock()
	defer fmu.Unlock()
	newList := []string{}
	for _, f := range friendships[user] {
		if f != friend {
			newList = append(newList, f)
		}
	}
	friendships[user] = newList
}

func ListFriends(user string) []string {
	fmu.Lock()
	defer fmu.Unlock()
	return friendships[user]
}

func OnlineFriends(user string) []string {
	fmu.Lock()
	defer fmu.Unlock()
	result := []string{}
	for _, f := range friendships[user] {
		if _, online := connections[f]; online {
			result = append(result, f)
		}
	}
	return result
}
