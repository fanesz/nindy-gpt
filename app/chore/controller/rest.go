package controller

import (
	"context"
	"net/http"
	"nindy-gpt/app/chore/entity"
	"nindy-gpt/app/chore/interfaces"
	"nindy-gpt/app/chore/service"
	"nindy-gpt/app/config"

	"github.com/fanesz/bindator"
	"github.com/gin-gonic/gin"
)

type nindyGPTController struct {
	service interfaces.NindyGPTService
}

func NewNindyGPTController() *nindyGPTController {
	client := config.GetClient()
	ctx := context.Background()

	return &nindyGPTController{
		service: service.NewNindyGPTService(client, ctx),
	}
}

func (c *nindyGPTController) Register(router *gin.Engine) {
	v1 := router.Group("/v1")

	v1.POST("/chat", func(ctx *gin.Context) {
		var req entity.ChatRequest

		val := bindator.BindBody(ctx, &req)
		if !val.Ok {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error":  val.Message,
				"fields": val.Errors,
			})
			return
		}

		resp, err := c.service.Chat(&req)
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(200, gin.H{"response": resp})
	})
}
