package router

import (
	"net/http"
	"time"

	"github.com/axiaoxin-com/ratelimiter"
	"github.com/gin-gonic/gin"
)

const REQUEST_LIMIT = 1
const LIMIT_INTERVAL = time.Second * 1

func rateLimiterConfig() gin.HandlerFunc {
	return ratelimiter.GinMemRatelimiter(ratelimiter.GinRatelimiterConfig{
		LimitKey: func(c *gin.Context) string {
			return c.ClientIP()
		},
		LimitedHandler: func(c *gin.Context) {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "Too many requests, please try again later",
			})
		},
		TokenBucketConfig: func(*gin.Context) (time.Duration, int) {
			return LIMIT_INTERVAL, REQUEST_LIMIT
		},
	})
}
