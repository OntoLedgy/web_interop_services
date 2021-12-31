package websocket

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
)

type WebSocketClients struct {
	ID   string
	Conn *websocket.Conn
	Pool *MessagePools
}

func (webSocketClient *WebSocketClients) Read() {
	defer func() {
		webSocketClient.Pool.Unregister <- webSocketClient

		webSocketClient.Conn.Close()

	}()

	for {
		messageType, messageBody, err :=
			webSocketClient.Conn.ReadMessage()

		if err != nil {
			log.Println(err)
			return
		}
		message := Messages{
			Type: messageType,
			Body: string(messageBody)}

		webSocketClient.Pool.Broadcast <- message

		fmt.Printf("Messages Received: %+v\n", message)
	}
}
