package datasync

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"gitbub.com/duiyuan/godemo/internal/datasync/pkg"
	"gitbub.com/duiyuan/godemo/internal/datasync/pkg/subscriber"
)

var txnSubscriber *subscriber.Subscriber
var memSubscriber *subscriber.Subscriber
var wg sync.WaitGroup

func handleTxMsg(msg []byte) {
	txn := &pkg.TxnSum{}
	str := string(msg)
	if err := json.Unmarshal(msg, &txn); err != nil {
		txnSubscriber.Logger.Fatal(err)
		return
	}
	fmt.Println(str)
	hash := txn.Hash
	function := txn.Function
	height := txn.Height
	ts := txn.Timestamp
	txnSubscriber.Logger.Printf("%s,%s,%d,%d", hash, function, height, ts)
}

func handleMemMsg(msg []byte) {
	var data interface{}

	if err := json.Unmarshal(msg, &data); err != nil {
		memSubscriber.Logger.Fatal(err)
		return
	}

	memSubscriber.Logger.Println(string(msg))
}

func Start() {
	ch := make(chan bool, 2)
	txnSubscriber = subscriber.NewSubscriber("txn_confirm_on_head", ch)
	txnSubscriber.SetHandler(handleTxMsg)
	go txnSubscriber.Connect()

	// memSubscriber = subscriber.NewSubscriber("mempool_insert", ch)
	// memSubscriber.SetHandler(handleMemMsg)
	// go memSubscriber.Connect()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGTERM, syscall.SIGINT)

	select {
	case <-interrupt:
		txnSubscriber.Cancel()
		memSubscriber.Cancel()
		return
	}

}
