package wss

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
)

var (
	IsUnexpectedCloseError = websocket.IsUnexpectedCloseError
	CloseAbnormalClosure   = websocket.CloseAbnormalClosure
	CloseGoingAway         = websocket.CloseGoingAway
	PongMessage            = websocket.PongMessage
	PingMessage            = websocket.PingMessage
)

type WSServer struct {
	Port int
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func NewWSServer() *WSServer {
	return &WSServer{
		Port: 9081,
	}
}

func (ws *WSServer) Serve() {
	http.HandleFunc("/api", ws.HandleWebSocet)
	port := ":" + strconv.Itoa(ws.Port)
	log.Printf("WebSocket server start on %s \n", port)
	log.Printf("Endpoint ws://127.0.0.1%s/ws\n", port)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func (ws *WSServer) HandleWebSocet(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}

	defer conn.Close()

	conn.SetPongHandler(func(appData string) error {
		log.Println("Received Ping")
		err := conn.WriteControl(PongMessage, []byte(appData), time.Now().Add(time.Second))
		if err != nil {
			log.Println("Error sending Pong:", err)
			return err
		}
		return nil
	})

	go ws.heartbeat(conn)

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			if IsUnexpectedCloseError(err, CloseGoingAway, CloseAbnormalClosure) {
				log.Printf("Error: %v", err)
			}
			break
		}

		log.Printf("Received: %s", message)
		err = conn.WriteMessage(messageType, message)
		if err != nil {
			log.Println("Write error:", err)
			break
		}
	}
}

func (ws *WSServer) heartbeat(conn *websocket.Conn) {
	ticker := time.NewTicker(time.Second * 30)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			err := conn.WriteControl(websocket.PongMessage, []byte("heartbeat"), time.Now().Add(time.Second))
			if err != nil {
				log.Println("Failed to send Ping:", err)
				conn.Close()
				return
			}
			log.Println("Sent Ping")
		}
	}
}
