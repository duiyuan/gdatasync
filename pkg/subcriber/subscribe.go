package subscriber

import (
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"syscall"

	"gitbub.com/duiyuan/gdatasync/pkg/conf"
	"github.com/gorilla/websocket"
)

const (
	CloseMessage         = websocket.CloseMessage
	CloseAbnormalClosure = websocket.CloseAbnormalClosure
)

type TxnSubscriber struct {
	ServerUrl string
	Logger    *log.Logger
}

func NewTxnSubscriber() *TxnSubscriber {
	return &TxnSubscriber{
		ServerUrl: conf.WSServerUrl,
	}
}

func (t *TxnSubscriber) Connect() {
	logFile, err := os.OpenFile("txn_confirm_on_head.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("fail to create file log %v", err)
		return
	}
	defer logFile.Close()

	logger := log.New(logFile, "subscribe", log.LstdFlags)

	conn, _, err := websocket.DefaultDialer.Dial(t.ServerUrl, nil)
	if err != nil {
		log.Fatalf("连接 Websocket 失败 %v", err)
		return
	}
	defer func() {
		defer conn.Close()
		logger.Println("websocket closed")
	}()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)

	msg := map[string]string{
		"req": "subscribe.txn_confirm_on_head",
	}
	submsg, _ := json.Marshal(msg)

	if err = conn.WriteMessage(websocket.TextMessage, submsg); err != nil {
		logger.Printf("fail to send subscribe message %v", err)
		return
	}

	logger.Println("sent subscribe message")

	done := make(chan bool)

	go func() {
		defer close(done)
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				logger.Printf("received message err: %v", err)
				return
			}
			logger.Printf("received message: %s", message)
		}
	}()

	select {
	case <-interrupt:
		logger.Println("got interrupt signal, now exit")
	case <-done:
		logger.Println("close server")
	}

	if err = conn.WriteMessage(CloseMessage, websocket.FormatCloseMessage(CloseAbnormalClosure, "")); err != nil {
		logger.Printf("fail to close websocket %v", err)
	}

}
