package blockFunction

import (
	"github.com/boltdb/bolt"
	"log"
)

type BlockchainIterator struct {
	currentHash	[]byte
	db		*bolt.DB
}

func (bci *BlockchainIterator) Next() *Block{
	var block *Block

	err := bci.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		encodedBlock := b.Get(bci.currentHash)
		block = DeserializeBlock(encodedBlock)
		return nil
	})

	if err != nil {
		log.Panic(err)
	}

	bci.currentHash = block.PrevBlockHash
	return block
}
