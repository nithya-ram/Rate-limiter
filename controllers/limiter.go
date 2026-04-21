package controllers

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type UserData struct {
	Count     int
	Timestamp time.Time
	LastReset time.Time
}

type RateLimiter struct {
	mu    sync.Mutex
	store map[string]*UserData
}

func NewRateLimiter() *RateLimiter {
	return &RateLimiter{
		store: make(map[string]*UserData),
	}
}
func (r *RateLimiter) Allow(userID string) bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	now := time.Now()
	data, ok := r.store[userID]
	if !ok {
		r.store[userID] = &UserData{Count: 1, LastReset: now}
		return true
	}

	// MAGIC LINE - Last 60 sec only count
	if now.Sub(data.LastReset) > time.Minute {
		data.Count = 0
		data.LastReset = now
	}

	if data.Count >= 5 {
		return false
	}
	data.Count++
	return true
}
func GetStatus(c *gin.Context) {
	limiter.mu.Lock() 
	defer limiter.mu.Unlock()

	stats := make(map[string]int)
	for user, data := range limiter.store {
		stats[user] = data.Count
	}
	c.JSON(http.StatusOK, stats)
}
