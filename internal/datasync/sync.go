package datasync

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/duiyuan/godemo/internal/datasync/options"
	"github.com/duiyuan/godemo/internal/datasync/pkg"
	"github.com/duiyuan/godemo/internal/datasync/pkg/subscriber"
)

var txnSubscriber *subscriber.Subscriber
var memSubscriber *subscriber.Subscriber
var confdMemSubscriber *subscriber.Subscriber

func handleTxMsg(msg []byte) {
	txn := &pkg.TxnSum{}
	str := string(msg)
	if err := json.Unmarshal(msg, &txn); err != nil {
		txnSubscriber.Logger.Print(err)
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
	var data pkg.InsertMempoolRep

	if err := json.Unmarshal(msg, &data); err != nil {
		memSubscriber.Logger.Print(err)
		return
	}

	txns := data.Txns
	for _, item := range txns {
		Hash := item.Hash
		time := item.Timestamp
		funcName := item.Function
		pack := item.Packing
		memSubscriber.Logger.Printf("%s,%s,%s,%d\n", Hash, funcName, pack, time)
	}
}

func handleComfdMemMsg(msg []byte) {
	var data interface{}

	if err := json.Unmarshal(msg, &data); err != nil {
		confdMemSubscriber.Logger.Print(err)
		return
	}

	confdMemSubscriber.Logger.Println(string(msg))
}

func Start(opts *options.Options) error {
	var wg sync.WaitGroup

	go func() {
		txnSubscriber = subscriber.NewSubscriber(opts.RuntimeOption.WSS, "txn_confirm_on_head", &wg)
		txnSubscriber.SetHandler(handleTxMsg)
		txnSubscriber.Connect()
	}()

	go func() {
		memSubscriber = subscriber.NewSubscriber(opts.RuntimeOption.WSS, "mempool_insert", &wg)
		memSubscriber.SetHandler(handleMemMsg)
		memSubscriber.Connect()
	}()

	go func() {
		confdMemSubscriber = subscriber.NewSubscriber(opts.RuntimeOption.WSS, "mempool_confirm", &wg)
		confdMemSubscriber.SetHandler(handleComfdMemMsg)
		confdMemSubscriber.Connect()
	}()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGTERM, syscall.SIGINT)

	select {
	case <-interrupt:
		txnSubscriber.Cancel()
		memSubscriber.Cancel()
		confdMemSubscriber.Cancel()
	case <-Wait(&wg):
		fmt.Println("all subscribers down")
	}

	return nil
}

func Wait(wg *sync.WaitGroup) <-chan bool {
	done := make(chan bool)

	go func() {
		wg.Wait()
		done <- true
	}()

	return done
}
