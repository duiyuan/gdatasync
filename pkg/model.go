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
