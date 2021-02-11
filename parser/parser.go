package parser

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
)

// ReverseBytes falan
func ReverseBytes(arr []byte) []byte {
	length := len(arr)
	for i := 0; i < length/2; i++ {
		arr[i], arr[length-1-i] = arr[length-1-i], arr[i]
	}

	return arr
}

// ReadVarInt is a function to read var_int values (1-9 bytes)
func ReadVarInt(file *os.File) int {
	data := make([]byte, 1)
	file.Read(data) // reading var_int first byte

	vi := 0
	switch data[0] {
	case 253:
		// uint16
	case 254:
		// uint32
	case 255:
		// uint64
	default:
		// uint8
		vi = int(data[0])
	}
	return vi
}

// ParseBlock falan
func ParseBlock(file *os.File, magicBytes []byte) (block, error) {
	var b block
	// src: https://en.bitcoin.it/wiki/Protocol_documentation //

	//	 [Block Parts]                    [Size]
	//	├ magic bytes                   - 4 bytes
	//	├ Block Size                    - 4 bytes
	//	├ block header                  - 80 bytes
	//	│   └┬ version                  - 4 bytes
	//	│	 ├ hash previous block      - 32 bytes
	//	│    ├ hash merkle root         - 32 bytes
	//	│    ├ time                     - 4 bytes
	//	│    ├ bits                     - 4 bytes
	//	│    └ nonce                    - 4 bytes
	//	├ for tx count                  - var_int
	//	     ├ tx data                  - remainder
	// 	     │	 └┬ version             - 4 bytes
	// 	     │	  ├ for input count     - var_int
	// 	     │	  │    ├ tx out hash    - 32 bytes
	// 	     │	  │    ├ tx out index   - 4 bytes
	// 	     │	  │    ├ script length  - var_int
	// 	     │	  │    └ sigscript      - script length
	// 	     │	  ├ for output count    - var_int
	// 	     │	  │    ├ value          - 8 bytes
	// 	     │	  │    ├ script length  - var_int
	// 	     │	  │    └ pkscript       - script length
	// 	     │	  └ locktime            - 4 bytes

	fmt.Print("Reading Block...\n")

	data := make([]byte, 4)
	file.Read(data) // reading magic bytes

	if !bytes.Equal(data, magicBytes) {
		return b, nil
	}

	file.Read(data) // reading block size
	b.blockSize = binary.LittleEndian.Uint32(data)

	// -----------  BLOCK HEADER  ----------- //

	file.Read(data) // reading version
	b.version = binary.LittleEndian.Uint32(data)

	b.hashPrevBlock = make([]byte, 32)
	file.Read(b.hashPrevBlock) // reading prev block hash

	b.hashMerkleRoot = make([]byte, 32)
	file.Read(b.hashMerkleRoot) // reading merkle root

	file.Read(data) // reading Timestamp
	b.time = binary.LittleEndian.Uint32(data)

	file.Read(data) // reading difficulty target
	b.bits = binary.LittleEndian.Uint32(data)

	file.Read(data) // reading nonce
	b.nonce = binary.LittleEndian.Uint32(data)

	b.tCount = ReadVarInt(file)

	// -----------  TRANSACTOINS  ----------- //

	for i := 0; i < b.tCount; i++ {

		var t transaction

		data = make([]byte, 4)
		file.Read(data) // reading version
		t.version = binary.LittleEndian.Uint32(data)

		t.inCount = ReadVarInt(file)

		// -----------  INPUTS  ----------- //

		for i := 0; i < t.inCount; i++ {

			var i txIn

			i.outHash = make([]byte, 32)
			file.Read(i.outHash) // reading tx out hash

			i.outIndex = make([]byte, 4)
			file.Read(i.outIndex) // Tx out index

			i.sciptLen = ReadVarInt(file)

			i.script = make([]byte, i.sciptLen)
			file.Read(i.script) // reading script

			i.sequence = make([]byte, 4)
			file.Read(i.sequence) // reading sequence

			t.inputs = append(t.inputs, i)
		}

		t.outCount = ReadVarInt(file)

		// -----------  OUTPUTS  ----------- //

		for i := 0; i < t.outCount; i++ {

			var o txOut

			data = make([]byte, 8)
			file.Read(data)                            // Tx out index
			o.value = binary.LittleEndian.Uint64(data) // 1e-8 (satoshi)

			o.scriptLen = ReadVarInt(file)

			data = make([]byte, o.scriptLen)
			file.Read(data) // reading script
			o.script = data

			t.outputs = append(t.outputs, o)
		}

		data = make([]byte, 4)
		file.Read(data) // reading locktime
		t.lockTime = binary.LittleEndian.Uint32(data)

		b.transactions = append(b.transactions, t)
	}
	return b, nil
}
