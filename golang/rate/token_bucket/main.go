package main

import (
	"sync"
	"time"
)

// 令牌桶算法
type tokenBucket struct {
	rate         int64
	capacity     int64
	tokens       int64
	lastTokenSec int64
	lock         sync.Mutex
}

func (bucket *tokenBucket) Allow() bool {
	bucket.lock.Lock()
	defer bucket.lock.Unlock()

	now := time.Now().Unix()
	bucket.tokens = bucket.tokens + (now-bucket.lastTokenSec)*bucket.rate
	if bucket.tokens > bucket.capacity {
		bucket.tokens = bucket.capacity
	}

	bucket.lastTokenSec = now
	if bucket.tokens > 0 {
		bucket.tokens--
		return true
	} else {
		return false
	}
}

func (bucket *tokenBucket) Init(rate, cap int64) {
	bucket.rate = rate
	bucket.capacity = cap
	bucket.tokens = 0
	bucket.lastTokenSec = time.Now().Unix()
}
