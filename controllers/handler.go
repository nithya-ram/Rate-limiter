package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Request struct {
	UserID  string `json:"user_id"`
	Payload string `json:"payload"`
}

var limiter = NewRateLimiter()

func HandleRequest(c *gin.Context) {
	var req Request

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	if !limiter.Allow(req.UserID) {
		c.JSON(http.StatusTooManyRequests, gin.H{
			"error": "rate limit exceeded (max 5 per minute)",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "request accepted",
	})
}
