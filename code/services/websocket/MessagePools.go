package websocket

import "fmt"

type MessagePools struct {
	Register   chan *WebSocketClients
	Unregister chan *WebSocketClients
	Broadcast  chan Messages
	Clients    map[*WebSocketClients]bool
}

func (pool *MessagePools) Start() {
	for {
		select {

		case client := <-pool.Register:
			pool.Clients[client] =
				true

			fmt.Println(
				"Size of Connection MessagePools: ",
				len(pool.Clients))

			for client, _ := range pool.Clients {
				fmt.Println(client)

				client.Conn.WriteJSON(
					Messages{
						Type: 1,
						Body: "New User Joined..."})
			}
			break

		case client := <-pool.Unregister:

			delete(
				pool.Clients,
				client)

			fmt.Println(
				"Size of Connection MessagePools: ",
				len(pool.Clients))

			for client, _ := range pool.Clients {
				client.Conn.WriteJSON(
					Messages{
						Type: 1,
						Body: "User Disconnected..."})
			}
			break

		case message := <-pool.Broadcast:

			fmt.Println(
				"Sending message to all clients in MessagePools")

			for client, _ := range pool.Clients {

				err :=
					client.Conn.WriteJSON(
						message)

				if err != nil {
					fmt.Println(err)
					return
				}
			}
		}
	}
}
