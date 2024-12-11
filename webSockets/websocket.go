package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func websocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	for {

		//read message from client
		_, message, err := conn.ReadMessage()

		if err != nil {
			log.Println(err)
			break
		}
		// show message
		log.Printf("Received message: %s", message)
		// send message to client
		err = conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			log.Println(err)
			break
		}
		if string(message) == "test" {
			returnMessage := "testing testing..."
			conn.WriteMessage(websocket.TextMessage, []byte(returnMessage))

		}
	}
}

func main() {
	fmt.Println("init")
	http.HandleFunc("/websocket", websocketHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
