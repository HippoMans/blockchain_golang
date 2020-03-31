package blockFunction

import (
	"fmt"
	"strconv"
)

func (cli *CLI) printChain() {
	bcIterator := cli.BC.Iterator()

	for {
		block := bcIterator.Next()
		fmt.Printf("Prev. hash : %x\n", block.PrevBlockHash)
		fmt.Printf("Data : %s\n", block.Data)
		fmt.Printf("Current Hash : %x\n", block.Hash)

		pow := NewProofOfWork(block)
		fmt.Printf("PoW : %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
}
