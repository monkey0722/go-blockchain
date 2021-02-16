package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"
)

// Block -> Type Definition.
type Block struct {
	nonce        int
	previousHash [32]byte
	timestamp    int64
	transactions []string
}

// NewBlock -> Create a new block.
func NewBlock(nonce int, previousHash [32]byte) *Block {
	b := new(Block)
	b.timestamp = time.Now().UnixNano()
	b.nonce = nonce
	b.previousHash = previousHash
	return b
}

// Print Make the output easy to read.
func (b *Block) Print() {
	fmt.Printf("timestamp     %d\n", b.timestamp)
	fmt.Printf("nonce     		%d\n", b.nonce)
	fmt.Printf("previousHash	%x\n", b.previousHash)
	fmt.Printf("transactions	%s\n", b.transactions)
}

// Hash Create a 32-byte hash
func (b *Block) Hash() [32]byte {
	m, _ := json.Marshal(b)
	fmt.Println(string(m))
	return sha256.Sum256([]byte(m))
}

// MarshalJSON Create MarshalJSON
func (b *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Timestamp    int64    `json:"timestamp"`
		Nonce        int      `json:"nonce"`
		PreviousHash [32]byte `json:"previous_hash"`
		Transactions []string `json:"transactions"`
	}{
		Timestamp:    b.timestamp,
		Nonce:        b.nonce,
		PreviousHash: b.previousHash,
		Transactions: b.transactions,
	})
}

// Blockchain -> Type Definition
type Blockchain struct {
	transactionPool []string
	chain           []*Block
}

// NewBlockchain -> Create a new blockchain.
func NewBlockchain() *Blockchain {
	b := &Block{}
	bc := new(Blockchain)
	bc.CreateBlock(0, b.Hash())
	return bc
}

// CreateBlock -> Create a block
func (bc *Blockchain) CreateBlock(nonce int, previousHash [32]byte) *Block {
	b := NewBlock(nonce, previousHash)
	bc.chain = append(bc.chain, b)
	return b
}

// LastBlock Create the LastBlock
func (bc *Blockchain) LastBlock() *Block {
	return bc.chain[len(bc.chain)-1]
}

// Print Make the output easy to read.
func (bc *Blockchain) Print() {
	for index, block := range bc.chain {
		fmt.Printf("%s Chain %d %s\n", strings.Repeat("=", 25), index, strings.Repeat("=", 25))
		block.Print()
	}
	fmt.Printf("%s\n", strings.Repeat("*", 25))
}

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	blockchain := NewBlockchain()
	blockchain.Print()

	previousHash := blockchain.LastBlock().Hash()
	blockchain.CreateBlock(10, previousHash)
	blockchain.Print()

	previousHash = blockchain.LastBlock().Hash()
	blockchain.CreateBlock(20, previousHash)
	blockchain.Print()
}
