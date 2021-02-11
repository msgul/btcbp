package main

import (
	"block-parser/parser"
	"fmt"
	"os"
)

func main() {
	// main network magic bytes
	magicBytes := []byte{249, 190, 180, 217}
	fmt.Print(magicBytes)

	// opening blk*.dat file
	file, err := os.Open("blocks/blk00000.dat")
	if err != nil {
		panic(err)
	}

	firstBlock, err := parser.ParseBlock(file, magicBytes) // parsing first block
	fmt.Print(firstBlock)

	parser.ReverseBytes(magicBytes)
	defer file.Close()
}
