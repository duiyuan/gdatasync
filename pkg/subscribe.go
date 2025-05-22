package pkg

import (
	"log"

	"github.com/gorilla/websocket"
)

type TxnSubscriber struct {
	ServerUrl string
}

func NewTxnSubscriber() *TxnSubscriber {
	return &TxnSubscriber{
		ServerUrl: WSServerUrl,
	}
}

func (t *TxnSubscriber) Connect() {
	conn, _, err := websocket.DefaultDialer.Dial(t.ServerUrl, nil)

	if err != nil {
		log.Fatalf("连接 Websocket 失败 %v", err)
		return
	}

	defer conn.Close()
}
