package handler

import (
	"github.com/KiseSaiyajin/testTask/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{services: service}
}
func (handler *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	auth := router.Group("/auth")
	{
		auth.POST("/sing-up", handler.signUp)
		auth.POST("/sing-in", handler.signIn)

	}
	return router
}
