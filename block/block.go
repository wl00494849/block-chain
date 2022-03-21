package block

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

type Block struct {
	TimeStamp     int64
	Data          []byte
	PrevBlockHask []byte
	Hash          []byte
}

func (b *Block) SetHash() {

	timeStamp := []byte(strconv.FormatInt(b.TimeStamp, 10))
	payload := bytes.Join([][]byte{b.PrevBlockHask, b.Data, timeStamp}, []byte{})
	hashValue := sha256.Sum256(payload)
	b.Hash = hashValue[:]

}

func CreateBlock(data string, prevBlockHask []byte) *Block {

	block := &Block{
		TimeStamp:     time.Now().Unix(),
		Data:          []byte(data),
		PrevBlockHask: prevBlockHask,
		Hash:          []byte{},
	}

	block.SetHash()

	return block
}
