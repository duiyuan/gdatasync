package pkg

type TxnSum struct {
	Hash     string `json:"Hash"`
	Height   int64  `json:"Height"`
	Function string `json:"Function"`
	// Timestamp int64  `json:"Timestamp"`

	ISN       int    `json:"ISN"`
	TTL       int    `json:"TTL"`
	BuildNum  int    `json:"BuildNum"`
	GasPrice  string `json:"GasPrice"`
	UTxnSize  int    `json:"uTxnSize"`
	BlockTime int64  `json:"BlockTime"`
	ExecStage string `json:"ExecStage"`
	Timestamp int64  `json:"Timestamp"`
}

type MempoolTxn struct {
	Version      int    `json:"Version"`
	Packing      string `json:"Packing"`
	Timestamp    int64  `json:"Timestamp"`
	TTL          int    `json:"TTL"`
	ISN          int16  `json:"ISN"`
	Size         int32  `json:"Size"`
	Mode         string `json:"Mode"`
	Function     string `json:"Function"`
	Hash         string `json:"Hash"`
	State        string `json:"State"`
	ExecStage    string `json:"ExecStage"`
	ConfirmState string `json:"ConfirmState"`
}

type InsertMempoolRep struct {
	Shard []int16      `json:"shard"`
	Txns  []MempoolTxn `json:"txns"`
}
