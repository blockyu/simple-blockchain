package types

type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
}

//func (b *Block) SetHash() {
//	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
//	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
//	hash := sha256.Sum256(headers)
//	b.Hash = hash[:]
//}
