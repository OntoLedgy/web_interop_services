package websocket

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

//TODO Add struct to wrap this

func ServeWebSocketConnection(
	pool *MessagePools,
	httpResponseWriter http.ResponseWriter,
	httpRequest *http.Request) {

	fmt.Println(
		"WebSocket Endpoint Hit")

	conn, err := Upgrade(
		httpResponseWriter,
		httpRequest)

	if err != nil {
		fmt.Fprintf(httpResponseWriter, "%+v\n", err)
	}

	client := &WebSocketClients{
		Conn: conn,
		Pool: pool,
	}

	pool.Register <- client

	client.Read()
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(httpRequest *http.Request) bool { return true },
}

func Upgrade(
	httpResponseWriter http.ResponseWriter,
	httpRequest *http.Request) (
	*websocket.Conn,
	error) {

	conn, err :=
		upgrader.Upgrade(
			httpResponseWriter,
			httpRequest,
			nil)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return conn, nil
}

func CreateMessagePool() *MessagePools {
	return &MessagePools{
		Register:   make(chan *WebSocketClients),
		Unregister: make(chan *WebSocketClients),
		Clients:    make(map[*WebSocketClients]bool),
		Broadcast:  make(chan Messages),
	}
}
