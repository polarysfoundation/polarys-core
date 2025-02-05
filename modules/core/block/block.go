package block

import (
	"log"

	pm256 "github.com/polarysfoundation/pm-256"
	"github.com/polarysfoundation/polarys-core/modules/common"
)

type Block struct {
	header Header
	hash   common.Hash
}

func NewBlock(header *Header) *Block {
	b := &Block{}

	b.header = *copyHeader(header)

	b.calculateHash()

	return b
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
