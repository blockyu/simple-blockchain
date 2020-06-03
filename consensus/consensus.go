package consensus

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
	"simple-blockchain/types"
	"simple-blockchain/utils"
)

const (
	TARGET_BITS = 1
	MAX_NONCE   = math.MaxInt64
)

type ProofOfWork struct {
	block  *types.Block
	target *big.Int
}

func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	fmt.Printf("Mining the block containing \"%s\" \n", pow.block.Data)

	for nonce < MAX_NONCE {
		data := pow.serializeData(nonce)
		hash = sha256.Sum256(data)
		fmt.Printf("\r%x %d", hash, nonce)

		hashInt.SetBytes(hash[:])
		if hashInt.Cmp(pow.target) == -1 {
			break
		} else {
			nonce++
		}
	}
	fmt.Print("\n\n")
	return nonce, hash[:]
}

func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int

	data := pow.serializeData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	isValid := hashInt.Cmp(pow.target) == -1
	return isValid
}

func (pow *ProofOfWork) serializeData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.Data,
			utils.IntToHex(pow.block.Timestamp),
			utils.IntToHex(int64(TARGET_BITS)),
			utils.IntToHex(int64(nonce)),
		},
		[]byte{},
	)
	return data
}

func NewProofOfWork(b *types.Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-TARGET_BITS))
	return &ProofOfWork{b, target}
}
