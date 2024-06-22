package client

import (
	"context"
	"fmt"
	"os"

	"github.com/mdp/qrterminal"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/store/sqlstore"
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
