package services

import (
	"compi-whatsapp/pkg/client"
	"compi-whatsapp/pkg/meow"
	"encoding/json"
	"net/http"
)

func HandleSendMessage(w http.ResponseWriter, r *http.Request) {
	var message meow.Message

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

	go meow.SendMessage(&meow.Message{
		To:      message.To,
		Message: message.Message,
	})

	w.Write([]byte("Message sent"))

}
