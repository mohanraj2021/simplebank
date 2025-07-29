package ratelimiter

import (
	"sync"
	"time"
)

type RateLimiter struct {
	Request      map[string]*Clientdata
	MaxRequest   int
	WindowSecond int
	Mu           sync.Mutex
}

type Clientdata struct {
	LastReset time.Time
	Count     int
}

func NewRateLimiter(maxrequest, windosecond int) *RateLimiter {
	return &RateLimiter{
		Request:      make(map[string]*Clientdata),
		MaxRequest:   maxrequest,
		WindowSecond: windosecond,
	}
}
