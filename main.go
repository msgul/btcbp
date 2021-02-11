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
	file, err := os.Open("blocks/blk00000.dat")
	if err != nil {
		panic(err)
	}

	a, err := parser.ParseBlock(file, magicBytes)
	a, err = parser.ParseBlock(file, magicBytes)
	a, err = parser.ParseBlock(file, magicBytes)
	a, err = parser.ParseBlock(file, magicBytes)
	fmt.Print(a)

	parser.ReverseBytes(magicBytes)
	defer file.Close()
}
