package router

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var routerInstance *gin.Engine

func InitializeRouter() {
	fmt.Println("===== Initialize Router =====")
	router := gin.Default()
	router.Use(corsHeaderConfig())
	router.Use(corsConfig())
	router.Use(rateLimiterConfig())

	routerInstance = router

	fmt.Println("✓ Gin router initialized")
}

func GetRouterInstance() *gin.Engine {
	return routerInstance
}

func UnsyncRouter(server *http.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		fmt.Println("✗ Router closed with error:", err)
		return
	}

	fmt.Println("✓ Router closed")
}
