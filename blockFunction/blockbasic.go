package blockFunction

import (
	"time"
)

type Block struct {
	Timestamp	int64		// 현재 시간의 타임스탬프
	Data		[]byte		// 트랜잭션 정보
	PrevBlockHash	[]byte		// 이전 블록의 해시값
	Hash		[]byte		// 현재 블록의 해시값
	Nonce		int		// 증명 검증
}

func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, 0}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}
