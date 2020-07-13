package types

type Template struct {
	TemplateName  string `json:"templateName"`
	ABI           string `json:"abi"`
	StorageLayout string `json:"storageLayout"`
}

// received from eth_getBlockByNumber
type RawBlock struct {
	Hash         string   `json:"hash"`
	ParentHash   string   `json:"parentHash"`
	StateRoot    string   `json:"stateRoot"`
	TxRoot       string   `json:"transactionsRoot"`
	ReceiptRoot  string   `json:"receiptsRoot"`
	Number       string   `json:"number"`
	GasLimit     string   `json:"gasLimit"`
	GasUsed      string   `json:"gasUsed"`
	Timestamp    string   `json:"timestamp"`
	ExtraData    string   `json:"extraData"`
	Transactions []string `json:"transactions"`
}

type Block struct {
	Hash         Hash   `json:"hash"`
	ParentHash   Hash   `json:"parentHash"`
	StateRoot    Hash   `json:"stateRoot"`
	TxRoot       Hash   `json:"txRoot"`
	ReceiptRoot  Hash   `json:"receiptRoot"`
	Number       uint64 `json:"number"`
	GasLimit     uint64 `json:"gasLimit"`
	GasUsed      uint64 `json:"gasUsed"`
	Timestamp    uint64 `json:"timestamp"`
	ExtraData    string `json:"extraData"`
	Transactions []Hash `json:"transactions"`
}

type Transaction struct {
	Hash              Hash            `json:"hash"`
	Status            bool            `json:"status"`
	BlockNumber       uint64          `json:"blockNumber"`
	BlockHash         Hash            `json:"blockHash"`
	Index             uint64          `json:"index"`
	Nonce             uint64          `json:"nonce"`
	From              Address         `json:"from"`
	To                Address         `json:"to"`
	Value             uint64          `json:"value"`
	Gas               uint64          `json:"gas"`
	GasPrice          uint64          `json:"gasPrice"`
	GasUsed           uint64          `json:"gasUsed"`
	CumulativeGasUsed uint64          `json:"cumulativeGasUsed"`
	CreatedContract   Address         `json:"createdContract"`
	Data              HexData         `json:"data"`
	PrivateData       HexData         `json:"privateData"`
	IsPrivate         bool            `json:"isPrivate"`
	Timestamp         uint64          `json:"timestamp"`
	Events            []*Event        `json:"events"`
	InternalCalls     []*InternalCall `json:"internalCalls"`
}

type InternalCall struct {
	From    Address `json:"from"`
	To      Address `json:"to"`
	Gas     uint64  `json:"gas"`
	GasUsed uint64  `json:"gasUsed"`
	Value   uint64  `json:"value"`
	Input   HexData `json:"input"`
	Output  HexData `json:"output"`
	Type    string  `json:"type"`
}

type Event struct {
	Index            uint64  `json:"index"`
	Address          Address `json:"address"`
	Topics           []Hash  `json:"topics"`
	Data             HexData `json:"data"`
	BlockNumber      uint64  `json:"blockNumber"`
	BlockHash        Hash    `json:"blockHash"`
	TransactionHash  Hash    `json:"transactionHash"`
	TransactionIndex uint64  `json:"transactionIndex"`
	Timestamp        uint64  `json:"timestamp"`
}
