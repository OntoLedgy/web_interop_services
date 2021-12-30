package main

import (
	"fmt"
	"net/http"

	//"github.com/gorilla/websocket"
	"github.com/OntoLedgy/web_interop_services/code/services/websocket"
)

func serveWs(w http.ResponseWriter, r *http.Request) {
	ws, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+V\n", err)
	}
	go websocket.Writer(ws)
	websocket.Reader(ws)
}

func setupRoutes() {
	http.HandleFunc("/ws", serveWs)
}

func main() {
	fmt.Println("Web Interop Backend services v0.01")
	setupRoutes()
	http.ListenAndServe(":8080", nil)
}
