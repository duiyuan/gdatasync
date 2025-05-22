package main

import subscriber "gitbub.com/duiyuan/gdatasync/pkg/subcriber"

func main() {
	txn := subscriber.NewTxnSubscriber()

	txn.Connect()
}
