package parser

type block struct {
	// magicBytes     []byte
	blockSize      uint32
	version        uint32 //
	hashPrevBlock  []byte //
	hashMerkleRoot []byte // Block
	time           uint32 // Header
	bits           uint32 //
	nonce          uint32 //
	tCount         int
	transactions   []transaction
}
