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
	Shard []int        `json:"shard"`
	Txns  []MempoolTxn `json:"txns"`
}

type Proof struct {
	Height   int   `json:"Height"`
	Shard    []int `json:"Shard"`
	Position []int `json:"Position"`
}

type Input struct {
	Amount string `json:"Amount"`
}

type Invocation struct {
	Status    string `json:"Status"`
	Return    []int  `json:"Return"`
	CoinDelta string `json:"CoinDelta"`
	GasFee    string `json:"GasFee"`
}

type Runtime struct {
	ConfirmedBy []interface{} `json:"ConfirmedBy"` // 空数组类型不确定，使用interface{}
}

type Transaction struct {
	Version      int        `json:"Version"`
	Packing      string     `json:"Packing"`
	Timestamp    int64      `json:"Timestamp"`
	RSN          int        `json:"RSN"`
	Proof        Proof      `json:"Proof"`
	Size         int        `json:"Size"`
	BuildNum     int        `json:"BuildNum"`
	GasOffered   int        `json:"GasOffered"`
	GasPrice     string     `json:"GasPrice"`
	Grouped      bool       `json:"Grouped"`
	UTxnSize     int        `json:"uTxnSize"`
	Mode         string     `json:"Mode"`
	Initiator    string     `json:"Initiator"`
	OrigTxHash   string     `json:"OrigTxHash"`
	OrigExecIdx  int        `json:"OrigExecIdx"`
	Function     string     `json:"Function"`
	Input        Input      `json:"Input"`
	Invocation   Invocation `json:"Invocation"`
	Hash         string     `json:"Hash"`
	State        string     `json:"State"`
	ExecStage    string     `json:"ExecStage"`
	ConfirmState string     `json:"ConfirmState"`
	Runtime      Runtime    `json:"Runtime"`
}

type ConfirmedTxDataResp struct {
	Shard []int         `json:"shard"`
	Txns  []Transaction `json:"txns"`
}
