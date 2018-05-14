package main

import (
	"crypto/sha256"
	"log"

	"golang.org/x/crypto/ripemd160"
)

const version = byte(0x00)

// GetAddress returns wallet address
func (wallet Wallet) GetAddress() []byte {
	// add this line
	public := append(wallet.PublicKey.X.Bytes(), wallet.PublicKey.Y.Bytes()...)

	publicSHA256 := sha256.Sum256(public)

	RIPEMD160Hasher := ripemd160.New()
	_, err := RIPEMD160Hasher.Write(publicSHA256[:])
	if err != nil {
		log.Panic(err)
	}
	publicRIPEMD160 := RIPEMD160Hasher.Sum(nil)

	versionedPayload := append([]byte{version}, publicRIPEMD160...)
	checksum := checksum(versionedPayload)

	fullPayload := append(versionedPayload, checksum...)
	address := Base58Encode(fullPayload)

	return address
}

// Checksum ...
func checksum(payload []byte) []byte {
	firstSHA := sha256.Sum256(payload)
	secondSHA := sha256.Sum256(firstSHA[:])

	return secondSHA[:4]
}
