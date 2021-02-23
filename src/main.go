package main

import (
	"fmt"
	"log"

	"./wallet"
)

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	w := wallet.NewWallet()
	fmt.Println(w.PrivateKeyStr())
	fmt.Println(w.PublickKeyStr())
	fmt.Println(w.BlockchainAddress())
}
