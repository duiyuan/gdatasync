package wss

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"sync"
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

type Request struct {
	Req string                 `json:"req"`
	Arg map[string]interface{} `json:"arg"`
}

type WSServer struct {
	Port            int
	mempSubscribers []*websocket.Conn
	mempSubLocker   sync.Mutex
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
		Port:            9081,
		mempSubscribers: []*websocket.Conn{},
		mempSubLocker:   sync.Mutex{},
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

	defer func() {
		conn.Close()
	}()

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
	go ws.scheduleMempoolInsert()

	var firstMessage Request

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

		if err := json.Unmarshal(message, &firstMessage); err != nil {
			log.Fatal(err)
			continue
		}

		if firstMessage.Req == "subscribe.mempool_insert" {
			ws.mempSubLocker.Lock()
			ws.mempSubscribers = append(ws.mempSubscribers, conn)
			ws.mempSubLocker.Unlock()
		}
	}
	if firstMessage.Req == "mempool_insert" {
		ws.mempSubLocker.Lock()
		ws.mempSubscribers = append(ws.mempSubscribers, conn)
		ws.mempSubLocker.Unlock()
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

func (ws *WSServer) scheduleMempoolInsert() {
	ticker := time.NewTicker(time.Second * 5)

	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			for _, conn := range ws.mempSubscribers {
				data := map[string]interface{}{
					"Hash":      "123",
					"Timestamp": time.Now().UnixMilli(),
				}
				msg, err := json.Marshal(data)
				if err != nil {
					log.Fatalln(err)
					return
				}
				if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
					log.Fatal(err)
					return
				}
			}
		}
	}

}
