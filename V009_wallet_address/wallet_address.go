package main

import (
	"crypto/sha256"
	"log"

	"golang.org/x/crypto/ripemd160"
)

const version = byte(0x00)

// HashPubKey hashes public key
func HashPubKey(pubKey []byte) []byte {
	publicSHA256 := sha256.Sum256(pubKey)

	RIPEMD160Hasher := ripemd160.New()
	_, err := RIPEMD160Hasher.Write(publicSHA256[:])
	if err != nil {
		log.Panic(err)
	}
	publicRIPEMD160 := RIPEMD160Hasher.Sum(nil)

	return publicRIPEMD160
}

// Checksum ...
func checksum(payload []byte) []byte {
	firstSHA := sha256.Sum256(payload)
	secondSHA := sha256.Sum256(firstSHA[:])

	return secondSHA[:4]
}

// GetAddress returns wallet address
// Version  Public key hash                           Checksum
// 00       62E907B15CBF27D5425399EBF6F0FB50EBB88F18  C29B7D93
func (wallet Wallet) GetAddress() []byte {
	//1. RIPEMD160(SHA256(PubKey))
	pubKeyHash := HashPubKey(wallet.PublicKey)

	//2. checksum = SHA256(SHA256(version + pubKeyHash)) 前四个字节
	versionedPayload := append([]byte{version}, pubKeyHash...)
	checksum := checksum(versionedPayload)

	//3. Base58Encode(version + pubKeyHash + checksum)
	fullPayload := append(versionedPayload, checksum...)
	address := Base58Encode(fullPayload)

	return address
}
