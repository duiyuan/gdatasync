package main

import (
	"encoding/json"
	"fmt"

	. "gitbub.com/duiyuan/gdatasync/pkg"
	subscriber "gitbub.com/duiyuan/gdatasync/pkg/subscriber"
)

var txnSubscriber *subscriber.Subscriber
var memSubscriber *subscriber.Subscriber

func handleTxMsg(msg []byte) {
	txn := &TxnSum{}
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

}

func main() {
	txnSubscriber = subscriber.NewSubscriber("txn_confirm_on_head")
	txnSubscriber.SetHandler(handleTxMsg)
	txnSubscriber.Connect()

	// memSubscriber = subscriber.NewSubscriber("mempool_insert")
	// memSubscriber.SetHandler(handleMemMsg)
	// memSubscriber.Connect()
}
