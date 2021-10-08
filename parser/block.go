package parser

type Block struct {
	// magicBytes     []byte
	BlockSize      uint32        `json:"size"`
	Version        uint32        `json:"ver"`
	HashPrevBlock  []byte        `json:"prev_block"`
	HashMerkleRoot []byte        `json:"mrkl_root"`
	Time           uint32        `json:"time"`
	Bits           uint32        `json:"bits"`
	Nonce          uint32        `json:"nonce"`
	TCount         int           `json:"n_tx"`
	Transactions   []Transaction `json:"tx"`
}
