package subscriber

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"sync"
	"syscall"

	"gitbub.com/duiyuan/godemo/internal/datasync/pkg/conf"
	"gitbub.com/duiyuan/godemo/internal/pkg/filesystem"
	"github.com/gorilla/websocket"
)

var (
	CloseMessage         = websocket.CloseMessage
	CloseAbnormalClosure = websocket.CloseAbnormalClosure
	FormatCloseMessage   = websocket.FormatCloseMessage
)

type Handler func(msg []byte)

type Subscriber struct {
	Subscription string
	Logger       *log.Logger
	HandleMsg    func(msg []byte)
	ctx          context.Context
	Cancel       context.CancelFunc
	wg           *sync.WaitGroup
}

func NewSubscriber(subscription string, wg *sync.WaitGroup) *Subscriber {
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	return &Subscriber{
		Subscription: subscription,
		ctx:          ctx,
		Cancel:       cancel,
		wg:           wg,
	}
}

func (t *Subscriber) SetHandler(handler Handler) {
	t.HandleMsg = handler
}

func (t *Subscriber) Connect() {
	dirname, err := filesystem.SureLogDir("datasync")
	if err != nil {
		log.Fatal("fail to make dirname for datasync log")
	}

	logPath := filepath.Join(dirname, t.Subscription+".log")
	logFile, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("fail to create file %v, err: %v\n", t.Subscription, err)
		return
	}
	defer logFile.Close()

	t.Logger = log.New(logFile, "", log.LstdFlags)

	conn, _, err := websocket.DefaultDialer.Dial(conf.WSServerUrl, nil)
	if err != nil {
		log.Fatalf("websocket connection error: %v\n", err)
		return
	}
	defer func() {
		defer conn.Close()
		t.Logger.Printf("%s websocket closed \n", t.Subscription)
	}()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)

	msg := map[string]interface{}{
		"req": "subscribe." + t.Subscription,
		"arg": map[string]interface{}{},
	}
	submsg, _ := json.Marshal(msg)

	if err = conn.WriteMessage(websocket.TextMessage, submsg); err != nil {
		t.Logger.Printf("fail to send subscribe message %v\n", err)
		return
	}

	t.Logger.Println("sent subscribe message")

	done := make(chan bool)
	go func() {
		defer close(done)
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				t.Logger.Printf("received message err: %v \n", err)
				return
			}
			t.HandleMsg(message)
		}
	}()

	select {
	case <-t.ctx.Done():
		t.Logger.Println("got context.done, now exit")
	case <-done:
		t.Logger.Println("close server")
		t.wg.Done()
	}
	if err = conn.WriteMessage(CloseMessage, FormatCloseMessage(CloseAbnormalClosure, "")); err != nil {
		t.Logger.Printf("fail to close websocket %v\n", err)
	}

}
