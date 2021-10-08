package parser

type TxIn struct {
	OutHash  []byte `json:"hash"`
	OutIndex []byte `json:"index"`
	SciptLen int    `json:"scr_len"`
	Script   []byte `json:"script"`
	Sequence []byte `json:"seq"`
}

type TxOut struct {
	Value     uint64 `json:"value"`
	ScriptLen int    `json:"scr_len"`
	Script    []byte `json:"script"`
}

type Transaction struct {
	Version  uint32  `json:"ver"`
	InCount  int     `json:"vin_sz"`
	OutCount int     `json:"vout_sz"`
	Inputs   []TxIn  `json:"inputs"`
	Outputs  []TxOut `json:"outs"`
	LockTime uint32  `json:"lock_time"`
}
