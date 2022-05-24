package main

import (
	"crypto/ecdsa"
	"errors"
	"fmt"
	"log"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil/hdkeychain"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/tyler-smith/go-bip39"
)

func DerivePrivateKey(path accounts.DerivatePath, masterKey *hdkeychain.ExtendedKey) (*ecdsa.PrivateKey, error) {
	var err error
	key := masterKey
	for _, n := range path {
		key, err = key.Child(n)
		if err != nil {
			return nil, err
		}
	}
	privateKey, err := key.ECPrivKey()
	privateKeyECDSA := privateKey.ToECDSA()
	if err != nil {
		return nil, err
	}
	return privateKeyECDSA, nil
}

func DerivePublicKey(privatekey *ecdsa.PrivateKey) (*ecdsa.PublicKey, error) {
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, error.New("Failed to get publicKey")
	}
	return publicKeyECDSA, nil
}

func DriveAddressFromMnemnoic() {
	path, err := accounts.ParseDerivationPath("m/44'/60'/0'/0/1")
	if err != nil {
		panic(err)
	}

	nm := "cargo emotion solt dentist client hint will penalty wrestle divide inform ranch"

	seed, err := bip39.NewSeedWithErrorChecking(nm, "")
	masterKey, err := hdkeychain.NewMaster(seed, &chaincfg.MainNetParams)
	if err != nil {
		fmt.Println("Failed to NewMaster",err)
		return
	}

	privateKey, err := DerivePrivateKey(path, masterKey)

	publicKey, err := DerivePublicKey(privateKey)

	address := crypto.PubkeyToAddress(*publicKey)
	fmt.Println(address.Hex())
}


func main() {}
	DeriveAddressFromMnemonic()
}
