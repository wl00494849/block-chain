package block

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
	"strconv"
)

type ProofOfWork struct {
	target *big.Int
	block  *Block
}

const targetBit = 12

func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBit))
	pow := &ProofOfWork{target, b}
	return pow
}

func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			[]byte(pow.block.Data),
			[]byte(strconv.FormatInt(pow.block.TimeStamp, 16)),
			[]byte(strconv.FormatInt(int64(targetBit), 16)),
			[]byte(strconv.FormatInt(int64(nonce), 16)),
		},
		[]byte{},
	)

	return data
}

func (pow *ProofOfWork) Proof() (int, []byte) {
	var hasInt big.Int
	var hash [32]byte
	n := 0

	fmt.Printf("Mining : %s \n", pow.block.Data)
	var i = 1
	for n < math.MaxInt64 {
		data := pow.prepareData(n)
		hash = sha256.Sum256(data)
		fmt.Println("第", i, "次Hash")
		fmt.Printf(" %x \n", hash)
		hasInt.SetBytes(hash[:])
		if hasInt.Cmp(pow.target) == -1 {
			break
		} else {
			n++
		}
		i++
	}

	fmt.Println()

	return n, hash[:]
}

func (pow *ProofOfWork) Validate() bool {

	var hasInt big.Int
	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hasInt.SetBytes(hash[:])

	isValid := hasInt.Cmp(pow.target) == -1

	return isValid
}
