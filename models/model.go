package models

type Block struct {
	BlockNumber      int64         `json:"block_number"`
	Timestamp        uint64        `json:"timestamp"`
	Hash             string        `json:"hash"`
	TransactionCount int           `json:"transaction_count"`
	Transactions     []Transaction `json:"transactions"`
}

type Transaction struct {
	Hash     string `json:"hash"`
	Value    string `json:"value"`
	Gas      uint64 `json:"gas"`
	Data     string `json:"data"`
	GasPrice uint64 `json:"gas_price"`
}

type StoreContractArgument struct {
	Name         string `json:"name"`
	Company      string `json:"company"`
	IdentityCard string `json:"identity_card"`
	HashDocs     string `json:"hash_docs"`
}

type GetContractArgument struct {
	IdentityCard string `json:"identity_card"`
	Company      string `json:"company"`
}

type GetContractResponse struct {
	HashDoc string `json:"hash_doc"`
}
