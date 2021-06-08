package websocket

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"

	"brsrker.com/diamond/proxyserver/internal/logger"
)

const TAG = "websocket"

var Upgrader = websocket.Upgrader{CheckOrigin: func(*http.Request) bool { return true }}
var Client *websocket.Conn

func socketConnection(w http.ResponseWriter, r *http.Request) {
	var err error
	client, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	Client = client
	ReceiveMessage()
}

func SendMessage(message string) {
	logger.Info(TAG, fmt.Sprintf("Sending Message to socket clientapp, message: %v", message))
	err := Client.WriteJSON(message)
	if err != nil {
		logger.Error(TAG, err)
		Client.Close()
	}
}

func ReceiveMessage() {
	for {
		var msg []byte
		// Read in a new message as JSON and map it to a Message object
		_, msg, err := Client.ReadMessage()
		if err != nil {
			logger.Error(TAG, err)
			// delete(clientapp, ws)
			break
		}
		logger.Info(TAG, string(msg))
	}
}
