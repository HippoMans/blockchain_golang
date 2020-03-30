package blockFunction

type Blockchain struct {
	BlockArray []*Block
}

func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.BlockArray[len(bc.BlockArray) -1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.BlockArray = append(bc.BlockArray, newBlock)
}

func NewBlockChain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}
