package service

import (
	"context"
	"realtime-chat/util"
	"time"
)

var ctx = context.Background()

// Lưu tin nhắn vào Redis list
func SaveMessage(msg string) {
	util.Rdb.LPush(ctx, "chat:history", msg)
	util.Rdb.LTrim(ctx, "chat:history", 0, 49) // chỉ giữ 50 tin mới nhất
}

// Lấy lịch sử chat
func GetHistory() ([]string, error) {
	return util.Rdb.LRange(ctx, "chat:history", 0, 49).Result()
}

// Lưu trạng thái online
func SetUserPresence(user string, online bool) {
	if online {
		util.Rdb.Set(ctx, "user:"+user+":online", "1", time.Minute*5)
	} else {
		util.Rdb.Del(ctx, "user:"+user+":online")
	}
}
