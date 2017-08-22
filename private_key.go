package main

import (
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
)

type PrivateKey struct {
	Mnemonic string `json:"mnemonic"`
	Key      string `json:"key"`
}

func getMnemonicPrivateKey(password string) (PrivateKey, error) {
	menmonic, error := getMnemonic()
	key, error := getPrivateKey(menmonic, password)
	return PrivateKey{menmonic, key}, error
}

// Mnemonic Generation
func getMnemonic() (string, error) {
	entropy, error := bip39.NewEntropy(256)
	mnemonic, error := bip39.NewMnemonic(entropy)
	return mnemonic, error
}

func getPrivateKey(mnemonic string, password string) (string, error) {
	// Seed for the Private key generation
	seed := bip39.NewSeed(mnemonic, password)
	masterKey, error := bip32.NewMasterKey(seed)

	//Private Key m/0
	bip32PrivateKey, error := masterKey.NewChildKey(0)

	// Private key for m/0/0
	firstDerivedKey, error := bip32PrivateKey.NewChildKey(0)

	return firstDerivedKey.String(), error
}
