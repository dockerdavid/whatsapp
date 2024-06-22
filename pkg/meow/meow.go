package meow

import (
	"compi-whatsapp/pkg/client"
	"context"
	"io"
	"net/http"

	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/proto/waE2E"
	"go.mau.fi/whatsmeow/types"
	"google.golang.org/protobuf/proto"
)

type File struct {
	To      string `json:"to" validate:"required"`
	URL     string `json:"url" validate:"required"`
	Caption string `json:"caption" validate:"required"`
}

type Message struct {
	To      string `json:"to" validate:"required"`
	Message string `json:"message" validate:"required"`
}

func SendFile(message *File) {
	resp, err := http.Get(message.URL)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	uploaded, err := client.Client.Upload(context.Background(), data, whatsmeow.MediaImage)

	if err != nil {
		return
	}

	targetJID := types.JID{
		User:   message.To,
		Server: "s.whatsapp.net",
	}

	msg := &waE2E.Message{
		DocumentMessage: &waE2E.DocumentMessage{
			Caption:       proto.String(message.Caption),
			URL:           proto.String(uploaded.URL),
			DirectPath:    proto.String(uploaded.DirectPath),
			MediaKey:      uploaded.MediaKey,
			Mimetype:      proto.String(http.DetectContentType(data)),
			FileEncSHA256: uploaded.FileEncSHA256,
			FileSHA256:    uploaded.FileSHA256,
			FileLength:    proto.Uint64(uint64(len(data))),
		},
	}

	client.Client.SendMessage(context.Background(), targetJID, msg)
}

func SendMessage(message *Message) {
	targetJID := types.JID{
		User:   message.To,
		Server: "s.whatsapp.net",
	}

	msg := &waE2E.Message{
		Conversation: &message.Message,
	}

	client.Client.SendMessage(context.Background(), targetJID, msg)
}
