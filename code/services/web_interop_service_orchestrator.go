package services

import (
	"fmt"
	"github.com/OntoLedgy/web_interop_services_backend/code/services/websocket"
	"net/http"
)

func OrchestrateWebInteropServices() {

	//TODO parameterise address
	fmt.Println("Web Interop Backend services v0.01")

	SetupRoutes()

	http.ListenAndServe(
		":8080",
		nil)
}

func SetupRoutes() {
	messagePool :=
		websocket.CreateMessagePool()

	go messagePool.
		Start()

	http.HandleFunc(
		"/ws",
		func(
			httpResponseWriter http.ResponseWriter,
			httpRequest *http.Request) {

			websocket.ServeWebSocketConnection(
				messagePool,
				httpResponseWriter,
				httpRequest)

		})
}
