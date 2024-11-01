package main

import (
	"nindy-gpt/app/config"
	"nindy-gpt/app/database"
	"nindy-gpt/app/router"
	"nindy-gpt/pkg/env"
)

func main() {
	env.InitializeEnv()
	database.InitializeDatabase()
	config.InitializeClient()
	config.InitializeThread()
	router.InitializeRouter()
	router.InitializeRoutes()

	routerInstance := router.GetRouterInstance()
	routerInstance.Run(env.BEHost + ":" + env.BEPort)
}
