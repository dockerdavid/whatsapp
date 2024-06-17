package services

import (
	"compi-whatsapp/pkg/client"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mau.fi/whatsmeow/proto/waE2E"
	"go.mau.fi/whatsmeow/types"
)

type Message struct {
	To      string `json:"to" binding:"required"`
	Message string `json:"message" binding:"required"`
}

func HandleSendMessage(c *gin.Context) {
	var message Message
	if err := c.ShouldBindBodyWithJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if client.Client != nil && client.Client.IsConnected() {

		targetJID := types.JID{
			User:   message.To,
			Server: "s.whatsapp.net",
		}

		go client.Client.SendMessage(context.Background(), targetJID, &waE2E.Message{
			Conversation: &message.Message,
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Client is not connected"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Message sent"})
}
