package main

import (
	"sync"
	"time"
)

// 计数器 固定时间窗口
type limitRate struct {
	threshold int           // 阈值
	begin     time.Time     // 计数开始时间
	cycle     time.Duration // 计数周期
	count     int           // 收到的请求数
	lock      sync.Mutex    // 锁
}

func (limit *limitRate) Allow() bool {
	limit.lock.Lock()
	defer limit.lock.Unlock()

	// 判断请求数是否达到阈值
	if limit.count == limit.threshold-1 {
		now := time.Now()
		// 达到阈值后，判断是否是请求周期内
		if now.Sub(limit.begin) >= limit.cycle {
			limit.Reset(now)
			return true
		}
		return false

	} else {
		limit.count++
		return true
	}
}

func (limit *limitRate) Set(rate int, cycle time.Duration) {
	limit.threshold = rate
	limit.begin = time.Now()
	limit.cycle = cycle
	limit.count = 0
}

func (limit *limitRate) Reset(begin time.Time) {
	limit.begin = begin
	limit.count = 0
}
