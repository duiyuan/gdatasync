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
	"github.com/duiyuan/godemo/internal/datasync/pkg/connection"
	"github.com/duiyuan/godemo/internal/datasync/pkg/subscriber"
)

var txnSubscriber *connection.SubscriberConn
var memSubscriber *connection.SubscriberConn
var confdMemSubscriber *connection.SubscriberConn

func handleTxMsg(msg []byte) {
	txn := &pkg.TxnSum{}
	// str := string(msg)
	if err := json.Unmarshal(msg, &txn); err != nil {
		txnSubscriber.Logger.Error(err)
		return
	}
	// fmt.Println(str)
	hash, ts, function, height := txn.Hash, txn.Timestamp, txn.Function, txn.Height
	txnSubscriber.Logger.Infof("%s,%s,%d,%d\n", hash, function, height, ts)
}

func handleMemMsg(msg []byte) {
	var data pkg.InsertMempoolRep

	if err := json.Unmarshal(msg, &data); err != nil {
		memSubscriber.Logger.Print(err)
		return
	}

	txns := data.Txns
	for _, item := range txns {
		Hash, time, funcName, pack := item.Hash, item.Timestamp, item.Function, item.Packing
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

	txnSubscriber = subscriber.MakeSubscriber(opts, "txn_confirm_on_head", &wg, handleTxMsg)
	// memSubscriber = subscriber.MakeSubscriber(opts, "mempool_insert", &wg, handleMemMsg)
	// confdMemSubscriber = subscriber.MakeSubscriber(opts, "mempool_confirm", &wg, handleComfdMemMsg)

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
