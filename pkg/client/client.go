package client

import (
	"context"
	"fmt"
	"os"

	"github.com/mdp/qrterminal"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/proto/waE2E"
	"go.mau.fi/whatsmeow/store/sqlstore"
	"go.mau.fi/whatsmeow/types/events"
	waLog "go.mau.fi/whatsmeow/util/log"
)

var (
	Client *whatsmeow.Client
)

func InitClient() {
	dbLog := waLog.Stdout("Database", "", true)

	container, err := sqlstore.New("sqlite3", "file:client.db?_foreign_keys=on", dbLog)
	if err != nil {
		panic(err)
	}

	deviceStore, err := container.GetFirstDevice()
	if err != nil {
		panic(err)
	}

	clientLog := waLog.Stdout("Client", "", true)
	Client = whatsmeow.NewClient(deviceStore, clientLog)

	Client.AddEventHandler(eventHandler)

	if Client.Store.ID == nil {
		qrChan, _ := Client.GetQRChannel(context.Background())
		err = Client.Connect()
		if err != nil {
			panic(err)
		}
		for evt := range qrChan {
			if evt.Event == "code" {
				qrterminal.GenerateHalfBlock(evt.Code, qrterminal.L, os.Stdout)
			} else {
				fmt.Println("Login event:", evt.Event)
			}
		}
	} else {
		err = Client.Connect()
		if err != nil {
			panic(err)
		}
	}
}

func eventHandler(evt interface{}) {
	switch v := evt.(type) {
	case *events.Message:
		handleMessage(v)
	}
}

func handleMessage(messageEvent *events.Message) {

	reply := `Compi amigo, por esta línea no es posible gestionar tus dudas, si tienes alguna solicitud te puedes comunicar a los siguientes canales:

Línea telefónica: 3176672305
WhatsApp: 3157817119`

	targetJID := messageEvent.Info.Chat

	if messageEvent.Info.Chat.User == "573108147655" && !messageEvent.Info.IsFromMe {
		go Client.SendMessage(context.Background(), targetJID, &waE2E.Message{
			Conversation: &reply,
		})
	}
}
