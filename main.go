package main

import (
	"fmt"
	bf "bitcoin_golang/blockFunction"
)

func main(){
	blockchain := bf.NewBlockChain()

	blockchain.AddBlock("Send 1 BTC to Hippo")
	blockchain.AddBlock("Send 100 BTC to HippoMans")

	for _, block := range blockchain.BlockArray {
		fmt.Printf("Prev. hash : %x\n", block.PrevBlockHash)
		fmt.Printf("Data : %s\n", block.Data)
		fmt.Printf("Current hash : %x\n", block.Hash)
		fmt.Println()
	}
}
