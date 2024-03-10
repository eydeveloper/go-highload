package handler

import (
	"github.com/eydeveloper/highload-social/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("api")
	{
		auth := api.Group("auth")
		{
			auth.POST("login", h.login)
			auth.POST("register", h.register)
		}

		user := api.Group("user")
		{
			user.GET(":id", h.getUserById)
		}
	}

	return router
}
