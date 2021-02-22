package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
)

// Wallet -> Type Definition.
type Wallet struct {
	privateKey *ecdsa.PrivateKey
	publickKey *ecdsa.PublicKey
}

// NewWallet ->
func NewWallet() *Wallet {
	w := new(Wallet)
	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	w.privateKey = privateKey
	w.publickKey = &w.privateKey.PublicKey
	return w
}

// PrivateKey -> Return privateKey.
func (w *Wallet) PrivateKey() *ecdsa.PrivateKey {
	return w.privateKey
}

// PrivateKeyStr -> Return privateKey as a string.
func (w *Wallet) PrivateKeyStr() string {
	return fmt.Sprintf("%x", w.privateKey.D.Bytes())
}

// PublickKey -> Return publickKey.
func (w *Wallet) PublickKey() *ecdsa.PublicKey {
	return w.publickKey
}

// PublickKeyStr -> > Return publickKey as a string.
func (w *Wallet) PublickKeyStr() string {
	return fmt.Sprintf("%x%x", w.publickKey.X.Bytes(), w.publickKey.Y.Bytes())
}
