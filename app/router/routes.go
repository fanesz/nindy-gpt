package router

import (
	"fmt"
	NindyGPTController "nindy-gpt/app/chore/controller"
)

func InitializeRoutes() {
	fmt.Println("===== Initialize Routes =====")
	router := GetRouterInstance()

	NindyGPTController.NewNindyGPTController().Register(router)

	fmt.Println("✓ Initialize", len(router.Routes()), "routes")
}
