package main

import (
	"sync"
	"time"
)

// 滑动窗口

type timeSlot struct {
	timestamp time.Time
	count     int
}

// 统计整个时间窗口中已经发生的请求次数
func countReq(win []*timeSlot) int {
	var count int
	for _, ts := range win {
		count += ts.count
	}
	return count
}

type SlidingWindowLimiter struct {
	mu           sync.Mutex
	SlotDuration time.Duration
	WinDuration  time.Duration
	numSlots     int
	windows      []*timeSlot
	maxReq       int
}

func NewSliding(slotDuration time.Duration, winDuration time.Duration, maxReq int) *SlidingWindowLimiter {
	return &SlidingWindowLimiter{
		SlotDuration: slotDuration,
		WinDuration:  winDuration,
		numSlots:     int(winDuration / slotDuration),
		maxReq:       maxReq,
	}
}

func (l *SlidingWindowLimiter) Allow() bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	now := time.Now()
	timeoutOffset := -1
	for i, ts := range l.windows {
		if ts.timestamp.Add(l.WinDuration).After(now) {
			break
		}
		timeoutOffset = i
	}

	if timeoutOffset > -1 {
		l.windows = l.windows[timeoutOffset+1:]
	}

	var result bool
	if countReq(l.windows) < l.maxReq {
		result = true
	}

	var lastSlot *timeSlot
	if len(l.windows) > 0 {
		lastSlot = l.windows[len(l.windows)-1]
		if lastSlot.timestamp.Add(l.SlotDuration).Before(now) {
			lastSlot = &timeSlot{timestamp: now, count: 1}
			l.windows = append(l.windows, lastSlot)

		} else {
			lastSlot.count++
		}

	} else {
		lastSlot = &timeSlot{timestamp: now, count: 1}
		l.windows = append(l.windows, lastSlot)
	}

	return result
}
