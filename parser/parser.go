package parser

import (
	"bytes"
	"encoding/binary"
	"os"
)

// ReverseBytes
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

// ParseBlock
func ParseBlock(file *os.File, magicBytes []byte) (Block, error) {

	data := make([]byte, 4)
	file.Read(data) // reading magic bytes

	var b Block

	if !bytes.Equal(data, magicBytes) {
		return b, nil
	}

	file.Read(data) // reading block size
	b.BlockSize = binary.LittleEndian.Uint32(data)

	// -----------  BLOCK HEADER  ----------- //

	file.Read(data) // reading version
	b.Version = binary.LittleEndian.Uint32(data)

	b.HashPrevBlock = make([]byte, 32)
	file.Read(b.HashPrevBlock) // reading prev block hash

	b.HashMerkleRoot = make([]byte, 32)
	file.Read(b.HashMerkleRoot) // reading merkle root

	file.Read(data) // reading Timestamp
	b.Time = binary.LittleEndian.Uint32(data)

	file.Read(data) // reading difficulty target
	b.Bits = binary.LittleEndian.Uint32(data)

	file.Read(data) // reading nonce
	b.Nonce = binary.LittleEndian.Uint32(data)

	b.TCount = ReadVarInt(file)

	// -----------  TRANSACTOINS  ----------- //

	for i := 0; i < b.TCount; i++ {

		var t Transaction

		data = make([]byte, 4)
		file.Read(data) // reading version
		t.Version = binary.LittleEndian.Uint32(data)

		t.InCount = ReadVarInt(file)

		// -----------  INPUTS  ----------- //

		for i := 0; i < t.InCount; i++ {

			var i TxIn

			i.OutHash = make([]byte, 32)
			file.Read(i.OutHash) // reading tx out hash

			i.OutIndex = make([]byte, 4)
			file.Read(i.OutIndex) // Tx out index

			i.SciptLen = ReadVarInt(file)

			i.Script = make([]byte, i.SciptLen)
			file.Read(i.Script) // reading script

			i.Sequence = make([]byte, 4)
			file.Read(i.Sequence) // reading sequence

			t.Inputs = append(t.Inputs, i)
		}

		t.OutCount = ReadVarInt(file)

		// -----------  OUTPUTS  ----------- //

		for i := 0; i < t.OutCount; i++ {

			var o TxOut

			data = make([]byte, 8)
			file.Read(data)                            // Tx out index
			o.Value = binary.LittleEndian.Uint64(data) // 1e-8 (satoshi)

			o.ScriptLen = ReadVarInt(file)

			data = make([]byte, o.ScriptLen)
			file.Read(data) // reading script
			o.Script = data

			t.Outputs = append(t.Outputs, o)
		}

		data = make([]byte, 4)
		file.Read(data) // reading locktime
		t.LockTime = binary.LittleEndian.Uint32(data)

		b.Transactions = append(b.Transactions, t)
	}
	return b, nil
}
