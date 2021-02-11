package parser

type txIn struct {
	outHash  []byte
	outIndex []byte
	sciptLen int
	script   []byte
	sequence []byte
}

type txOut struct {
	value     uint64
	scriptLen int
	script    []byte
}

type transaction struct {
	version  uint32
	inCount  int
	outCount int
	inputs   []txIn
	outputs  []txOut
	lockTime uint32
}
