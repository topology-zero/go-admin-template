package redis

import (
	"context"
	"time"
)

const (
	lockKey    = "LOCK:"
	defaultTTL = 10 * time.Second
)

// GetLock 获取锁
func GetLock(key string, t ...time.Duration) (bool, error) {
	ttl := defaultTTL
	if len(t) != 0 {
		ttl = t[0]
	}
	return Client.SetNX(context.Background(), lockKey+key, 1, ttl).Result()
}

// ReleaseLock 释放锁
func ReleaseLock(key string) {
	Client.Del(context.Background(), lockKey+key)
}
