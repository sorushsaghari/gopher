package controllers

import (
	"../services"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"log"
	"fmt"
)

type ChatController struct {
	 upgrader  websocket.Upgrader
	 clients  map[*websocket.Conn]bool
	 broadcast  chan services.MessageDto
	cs services.MessageService
}
func NewChatController(cs services.MessageService) *ChatController{
	return  &ChatController{
		cs: cs,
		upgrader: websocket.Upgrader{CheckOrigin:func(r *http.Request) bool {
			return true
		},},
		clients: make(map[*websocket.Conn]bool),
		broadcast: make(chan services.MessageDto),
	}
}

func (cs* ChatController) HandleConnection(c *gin.Context)  {
	ws, err := cs.upgrader.Upgrade(c.Writer, c.Request, nil)
	defer ws.Close()
	if err != nil {
		 c.AbortWithError(http.StatusInternalServerError,  err)
	}

	cs.clients[ws] = true

	for {
		var msg services.MessageDto

		err := ws.ReadJSON(&msg)
		if err != nil {
			c.AbortWithError(300, err)
			delete(cs.clients, ws)
			break
		}

		cs.broadcast <- msg
		fmt.Println(cs)
	}
	fmt.Println(cs)

}

func (cs* ChatController) HandleMessages() {
	for {
		// Grab the next message from the broadcast channel
		msg := <- cs.broadcast
		// Send it out to every client that is currently connected
		for client := range cs.clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(cs.clients, client)
			}
		}
		fmt.Println(msg)
	}
}