package blockFunction

import (
	"bytes"
	"time"
	"strconv"
	"crypto/sha256"
)

type Block struct {
	Timestamp	int64		// 현재 시간의 타임스탬프
	Data		[]byte		// 트랜잭션 정보
	PrevBlockHash	[]byte		// 이전 블록의 해시값
	Hash		[]byte		// 현재 블록의 해시값
}

func (b *Block) SetHash(){
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}

func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}}
	block.SetHash()
	return block
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}
