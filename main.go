package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/msgul/btcbp/parser"
)

func main() {
	// bitcoin main network magic bytes
	magicBytes := []byte{249, 190, 180, 217}

	// opening blk*.dat file
	file, err := os.Open("blocks/blk00000.dat")
	if err != nil {
		panic(err)
	}

	block, err := parser.ParseBlock(file, magicBytes) // parsing first block

	b, err := json.MarshalIndent(block, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(b))

	defer file.Close()
}
