package main

import (
	"fmt"
	"strconv"
	"bytes"
	"crypto/sha256"
	"time"
)


// 블록 구조체
type Block struct{
	Timestamp	int64		// 현재 시간의 타임스탬프
	Data		[]byte		// 트랜잭션 정보
	PrevBlockHash	[]byte		// 이전 블록의 해시값
	Hash		[]byte		// 현재 블록의 해시값
}

//블록을 구성하는 문자열에 대해 SHA-256 해시 계산 메서드
func (b *Block) SetHash(){
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}

// 블록을 생성하는 함수
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}}
	block.SetHash()
	return block
}

// 블록체인 구조체
type BlockChain struct{
	blocks []*Block
}

// 블록체인에 블록 추가 함수
func (chain *BlockChain) AddBlock(data string){
	prevBlock := chain.blocks[len(chain.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, newBlock)
}

//초기 Genesis 블록 생성 함수
func NewGenesisBlock() *Block{
	return NewBlock("Genesis Block", []byte{})
}

//초기 제네시스 블록을 블록체인에 생성하는 함수
func NewBlockChain() *BlockChain{
	return &BlockChain{[]*Block{NewGenesisBlock()}}
}

func main(){
	blockchain := NewBlockChain()
	blockchain.AddBlock("Send 1 BTC to Hippo")
	blockchain.AddBlock("Send 100 BTC to HippoMans")

	for _, block := range blockchain.blocks{
		fmt.Printf("Prev.Hash : %x\n", block.PrevBlockHash)
		fmt.Printf("Data : %s\n", block.Data)
		fmt.Printf("Current Hash : %x\n", block.Hash)
		fmt.Println("\n")
	}
}
