package util

import (
	"golang.org/x/time/rate"
	"sync"
)

var limiters = make(map[string]*rate.Limiter)
var mu sync.Mutex

// Tạo limiter cho từng user
func getLimiter(user string) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()

	limiter, exists := limiters[user]
	if !exists {
		// Cho phép 1 tin nhắn / giây, tối đa 5 burst
		limiter = rate.NewLimiter(1, 5)
		limiters[user] = limiter
	}
	return limiter
}

func AllowRequest(user string) bool {
	limiter := getLimiter(user)
	return limiter.Allow()
}
