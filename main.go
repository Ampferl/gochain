package main

import (
	"fmt"
)

func main() {
	bc := NewBlockchain()

	bc.AddBlock("Send 3 Coins to Jonas")
	bc.AddBlock("Send 6 more Coins to Jonas")

	for _, block := range bc.blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
