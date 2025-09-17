package pubsub

import (
	"context"
	"fmt"
	"log"
	"week3-chat-app/internal/storage"

	"github.com/redis/go-redis/v9"
)

const ChannelName = "chat-messages"

func PublishMessage(sender, msg string) {
	err := storage.Rdb.Publish(context.Background(), ChannelName, fmt.Sprintf("%s: %s", sender, msg)).Err()
	if err != nil {
		log.Println("Error publishing:", err)
	}
}

func SubscribeMessages(onMessage func(msg string)) {
	subscriber := storage.Rdb.Subscribe(context.Background(), ChannelName)
	ch := subscriber.Channel()

	go func() {
		for msg := range ch {
			onMessage(msg.Payload)
		}
	}()
}
