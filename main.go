package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

type Client struct {
	conn     *websocket.Conn
	username string
}

var clients = make(map[*websocket.Conn]Client)

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	username := r.URL.Query().Get("username")
	client := Client{conn: conn, username: username}
	clients[conn] = client

	log.Println("client", username, "connected")

	if err := conn.WriteMessage(1, []byte("Hi "+username+"!")); err != nil {
		log.Println(err)
		delete(clients, conn)
		conn.Close()
		return
	}

	go reader(client)
}

func reader(client Client) {
	conn := client.conn
	username := client.username

	defer func() {
		conn.Close()
		delete(clients, conn)
	}()

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		fmt.Println(username+":", string(message))

		for _, c := range clients {
			if c.conn != conn {
				if err := c.conn.WriteMessage(messageType, []byte(username+": "+string(message))); err != nil {
					log.Println(err)
					return
				}
			}
		}
	}
}

func homePage(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "index.html") // `index.html` faylni beradi
}

func setUpRoutes() {
    http.HandleFunc("/", homePage) // Asosiy sahifa
    http.HandleFunc("/ws", wsEndpoint) // WebSocket
}

func main() {
    setUpRoutes()
    fmt.Println("Server ishga tushdi: http://10.10.1.189:9999/")
    http.ListenAndServe("0.0.0.0:9999", nil) 
}

