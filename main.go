package main

import (
	"fmt"
	"strconv"
	"bytes"
	"crypto/sha256"
	"time"
	"math/big"
	"math"
)

// 채굴 난이도
const targetBits = 10

// 최대 nonce 크기
const MaxNonce = math.MaxInt64

// 블록 구조체
type Block struct{
	Timestamp	int64		// 현재 시간의 타임스탬프
	Data		[]byte		// 트랜잭션 정보
	PrevBlockHash	[]byte		// 이전 블록의 해시값
	Hash		[]byte		// 현재 블록의 해시값
	Nonce		int		// 증명 검증에 nonce가 필요
}


// 블록을 생성하는 함수
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, 0}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce
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

//작업 증명 구조체
type ProofOfWork struct{
	block *Block			//블록 포인터
	target *big.Int			//타겟 포인터
}

//초기 작업 증명을 위한 타겟값을 결정 -> NewProofOfWork() 함수를 통해 target 결정 
func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)				//bit.Int를 1로 초기화
	target.Lsh(target, uint(256-targetBits))	//256 - 24 비트 만큼 좌측 시프트 연산을 취한다. 

	pow := &ProofOfWork{b, target}
	return pow
}

//10진수 int를 16진수 hex로 변환
func IntToHex(n int64) []byte {
    return []byte(strconv.FormatInt(n, 16))
}

//데이터 준비를 위한 preData() 함수 코드
func (pow *ProofOfWork) preData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.Data,
			IntToHex(pow.block.Timestamp),
			IntToHex(int64(targetBits)),
			IntToHex(int64(nonce)),
		},
		[]byte{},
	)
	return data
}

//POW 알고리즘 실행 Run() 함수 코드
func (pow *ProofOfWork) Run() (int, []byte){
	var hashInt big.Int				//hashInt는 hash의 정수 표현값으로 nonce 카운터
	var hash [32]byte
	nonce := 0

	fmt.Printf("Mining the block containing \"%s\"\n", pow.block.Data)

	for nonce < MaxNonce {
		data := pow.preData(nonce)
		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])
		if hashInt.Cmp(pow.target) == -1{
			break
		} else {
			nonce++
		}
	}
	fmt.Print("\n\n")
	return nonce, hash[:]
}

//작업증명을 검증하는 Validate() 함수 코드
func (pow *ProofOfWork) Validate() bool{
	var hashInt big.Int

	data := pow.preData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	isValid := hashInt.Cmp(pow.target) == -1
	return isValid
}

func main(){
	empty := []byte{}
	blockchain := NewBlockChain()
	blockchain.AddBlock("Send 1 BTC to Hippo")
	blockchain.AddBlock("Send 100 BTC to HippoMans")
	blockchain.AddBlock("Send 10000 BTC to Mans")
	
	fmt.Println("\n블록의 내용 확인")
	for _, block := range blockchain.blocks{
		pow := NewProofOfWork(block)
		if bytes.Compare(block.PrevBlockHash, empty) != 0{
			fmt.Printf("Prev.Hash : %x\n", block.PrevBlockHash)
		}else{
			fmt.Println("It is Genesisblock. so, PrevBlockHash is not exist")
	}
		fmt.Printf("Data : %s\n", block.Data)
		fmt.Printf("Current Hash : %x\n", block.Hash)
		fmt.Printf("Pow : %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
