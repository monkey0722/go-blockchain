package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)

// Block -> Type Definition.
type Block struct {
	nonce        int
	previousHash string
	timestamp    int64
	transactions []string
}

// NewBlock -> Create a new block.
func NewBlock(nonce int, previousHash string) *Block {
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
	fmt.Printf("previousHash	%s\n", b.previousHash)
	fmt.Printf("transactions	%s\n", b.transactions)
}

// Blockchain -> Type Definition
type Blockchain struct {
	transactionPool []string
	chain           []*Block
}

// NewBlockchain -> Create a new blockchain.
func NewBlockchain() *Blockchain {
	bc := new(Blockchain)
	bc.CreateBlock(0, "Init hash")
	return bc
}

// CreateBlock -> Create a block
func (bc *Blockchain) CreateBlock(nonce int, previousHash string) *Block {
	b := NewBlock(nonce, previousHash)
	bc.chain = append(bc.chain, b)
	return b
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

	blockchain.CreateBlock(10, "hash 1")
	blockchain.Print()
}
