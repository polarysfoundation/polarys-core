package block

import (
	"encoding/json"
	"log"

	pm256 "github.com/polarysfoundation/pm-256"
	"github.com/polarysfoundation/polarys-core/modules/common"
	"github.com/polarysfoundation/polarys-core/modules/core/transaction"
)

type Block struct {
	header       Header
	transactions []*transaction.Transaction

	hash common.Hash
}

func NewBlock(header *Header, tx []*transaction.Transaction) *Block {
	b := &Block{
		transactions: make([]*transaction.Transaction, 0),
	}

	b.header = *copyHeader(header)

	if len(tx) > 0 {
		b.transactions = append(b.transactions, tx...)
	}

	b.calculateHash()

	return b
}

func (b *Block) UnmarshalJSON(data []byte) error {
	aux := &struct {
		Header       Header                     `json:"header"`
		Transactions []*transaction.Transaction `json:"transactions"`
		Hash         common.Hash                `json:"hash"`
	}{}

	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}

	b.header = aux.Header
	b.transactions = aux.Transactions
	b.hash = aux.Hash
	return nil
}

func (b *Block) MarshalJSON() ([]byte, error) {
	aux := &struct {
		Header       Header                     `json:"header"`
		Transactions []*transaction.Transaction `json:"transactions"`
		Hash         common.Hash                `json:"hash"`
	}{
		Header:       b.header,
		Transactions: b.transactions,
		Hash:         b.hash,
	}

	return json.Marshal(aux)
}

func (b *Block) calculateHash() {
	header := &b.header
	var buf common.Hash

	d, err := header.marshal()
	if err != nil {
		log.Printf("error: %v", err)
		panic("error marshaling header block")
	}

	h := pm256.New256()
	h.Write(d)
	h.Sum(buf[:0])
	b.hash = buf
}

func (b *Block) Print() {
	log.Printf("Hash: %s\n", b.hash.String())
	log.Printf("PrevHash: %s\n", b.header.PrevHash.String())
	log.Printf("Height: %d\n", b.header.Height)
	log.Printf("Nonce: %d\n", b.header.Nonce)
	log.Printf("Coinbase: %s\n", b.header.Coinbase.String())
	log.Printf("GasPrice: %d\n", b.header.GasPrice)
	log.Printf("GasLimit: %d\n", b.header.GasLimit)
	log.Printf("Difficulty: %d\n", b.header.Difficulty)
	log.Printf("Timestamp: %d\n", b.header.Timestamp)
	log.Printf("Signature: %x\n", b.header.Signature)
	log.Printf("TotalDifficulty: %d\n", b.header.TotalDifficulty)
	log.Printf("Data: %s\n", b.header.Data)
	log.Printf("ExtraData: %s\n", b.header.ExtraData)
}

func (b *Block) Hash() common.Hash        { return b.hash }
func (b *Block) PrevHash() common.Hash    { return b.header.PrevHash }
func (b *Block) Height() uint64           { return b.header.Height }
func (b *Block) Nonce() uint64            { return b.header.Nonce }
func (b *Block) Coinbase() common.Address { return b.header.Coinbase }
func (b *Block) GasPrice() uint64         { return b.header.GasPrice }
func (b *Block) GasLimit() uint64         { return b.header.GasLimit }
func (b *Block) Difficulty() uint64       { return b.header.Difficulty }
func (b *Block) Timestamp() uint64        { return b.header.Timestamp }
func (b *Block) Signature() []byte        { return b.header.Signature }
func (b *Block) TotalDifficulty() uint64  { return b.header.TotalDifficulty }
func (b *Block) Data() []byte             { return b.header.Data }
func (b *Block) ExtraData() []byte        { return b.header.ExtraData }
