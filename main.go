package main

import (
	"nindy-gpt/app/config"
	"nindy-gpt/app/router"
	"nindy-gpt/pkg/env"
)

func main() {
	env.InitializeEnv()
	config.InitializeClient()
	config.InitializeThread()
	router.InitializeRouter()
	router.InitializeRoutes()

	routerInstance := router.GetRouterInstance()
	routerInstance.Run(":5000")
}
