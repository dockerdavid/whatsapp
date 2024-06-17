package routes

import (
	"compi-whatsapp/pkg/services"

	"github.com/gin-gonic/gin"
)

func InitRoutes(e *gin.Engine) {
	api := e.Group("/api")

	api.POST("/send-message", services.HandleSendMessage)
}
