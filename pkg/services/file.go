package services

import (
	"compi-whatsapp/pkg/client"
	"compi-whatsapp/pkg/meow"
	"compi-whatsapp/pkg/queue"
	"encoding/json"
	"net/http"
)

func HandleSendFile(w http.ResponseWriter, r *http.Request) {
	var message meow.File

	w.WriteHeader(http.StatusOK)

	err := json.NewDecoder(r.Body).Decode(&message)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	err = validate.Struct(message)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	if client.Client == nil || !client.Client.IsConnected() {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Client is not connected"))
		return
	}

	go queue.AddFileToQueue(&queue.FileQueue{
		To:      message.To,
		URL:     message.URL,
		Caption: message.Caption,
	})

	w.Write([]byte("File added to queue"))
}
