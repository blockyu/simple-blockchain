package core

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"simple-blockchain/consensus"
	"simple-blockchain/types"
	"time"
)

type Blockchain struct {
	tip []byte
	DB  *bolt.DB
}

func (bc *Blockchain) AddBlock(data string) {
	var lastHash []byte

	err := bc.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("blocksBucket"))
		lastHash = b.Get([]byte("l"))
		return nil
	})
	if err != nil {
		fmt.Printf(err.Error())
	}
	newBlock := NewBlock(data, lastHash)

	err = bc.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("blocksBucket"))
		err := b.Put(newBlock.Hash, newBlock.Serialize())
		if err != nil {
			fmt.Printf(err.Error())
		}
		err = b.Put([]byte("l"), newBlock.Hash)
		bc.tip = newBlock.Hash
		return nil
	})
	if err != nil {
		fmt.Printf(err.Error())
	}
}

type BlockchainIterator struct {
	currentHash 	[]byte
	db				*bolt.DB
}

func (bc *Blockchain) Iterator() *BlockchainIterator {
	bci := &BlockchainIterator{bc.tip, bc.DB}
	return bci
}

func (i *BlockchainIterator) Next() *types.Block {
	var block *types.Block

	err := i.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("blocksBucket"))
		encodedBlock := b.Get(i.currentHash)
		block = types.DeserializeBlock(encodedBlock)
		return nil
	})
	if err != nil {
		fmt.Printf(err.Error())
	}
	i.currentHash = block.PrevBlockHash

	return block
}

func NewBlockchain() *Blockchain {
	var tip []byte
	db, err := bolt.Open("my.DB", 0600, nil)
	if err != nil {
		fmt.Printf(err.Error())
	}
	err = db.Update(func(tx *bolt.Tx) error {
		//b, err := tx.CreateBucketIfNotExists([]byte("blocksBucket"))
		b := tx.Bucket([]byte("blocksBucket"))
		if b == nil {
			genesis := NewGenesisBlock()
			b, err := tx.CreateBucket([]byte("blocksBucket"))
			if err != nil {
				log.Fatal(err.Error())
			}
			err = b.Put(genesis.Hash, genesis.Serialize())
			err = b.Put([]byte("l"), genesis.Hash)
			tip = genesis.Hash
		} else {
			tip = b.Get([]byte("l"))
		}
		return nil
	})

	bc := Blockchain{tip, db}
	return &bc
}

func NewGenesisBlock() *types.Block {
	return NewBlock("Genesis Block generate !! ", []byte{})
}

func NewBlock(data string, prevBlockHash []byte) *types.Block {
	block := &types.Block{Timestamp: time.Now().Unix(), Data: []byte(data), PrevBlockHash: prevBlockHash, Hash: []byte{}}
	pow := consensus.NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash
	block.Nonce = nonce

	return block
}

