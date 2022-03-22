package block

import (
	"time"
)

type Block struct {
	TimeStamp     int64
	Data          string
	PrevBlockHask []byte
	Hash          []byte
	Nonce         int
}

type BlockChain struct {
	Blocks []*Block
}

func CreateBlock(data string, prevBlockHask []byte) *Block {

	block := &Block{
		TimeStamp:     time.Now().Unix(),
		Data:          data,
		PrevBlockHask: prevBlockHask,
		Hash:          []byte{},
	}

	pow := NewProofOfWork(block)
	nonce, hash := pow.Proof()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

func (blockChain *BlockChain) AddBlock(data []byte) {
	prevBlock := blockChain.Blocks[len(blockChain.Blocks)-1]
	newBlock := CreateBlock(string(data), prevBlock.Hash)
	blockChain.Blocks = append(blockChain.Blocks, newBlock)
}

func CreateInitBlock() *Block {
	return CreateBlock("firstBlock", []byte{})
}

func CreateBlockChain() *BlockChain {
	return &BlockChain{[]*Block{CreateInitBlock()}}
}
