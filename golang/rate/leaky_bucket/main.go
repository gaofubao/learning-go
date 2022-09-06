package main

import (
	"math"
	"sync"
	"time"
)

// 漏桶算法
type leakyBucket struct {
	rate       float64
	capacity   float64
	water      float64
	lastLeakMs int64
	lock       sync.Mutex
}

func (bucket *leakyBucket) Allow() bool {
	bucket.lock.Lock()
	defer bucket.lock.Unlock()

	now := time.Now().UnixNano() / 1e6
	leakyWater := bucket.water - (float64(now-bucket.lastLeakMs) * bucket.rate / 1000)
	bucket.water = math.Max(0, leakyWater)
	bucket.lastLeakMs = now
	if bucket.water+1 <= bucket.capacity {
		bucket.water++
		return true
	} else {
		return false
	}
}

func (bucket *leakyBucket) Init(rate, cap float64) {
	bucket.rate = rate
	bucket.capacity = cap
	bucket.water = 0
	bucket.lastLeakMs = time.Now().UnixNano() / 1e6
}
