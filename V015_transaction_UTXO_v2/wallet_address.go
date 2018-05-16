package main

import (
	"bytes"
	"crypto/sha256"
	"log"

	"golang.org/x/crypto/ripemd160"
)

const version = byte(0x00)
const addressChecksumLen = 4

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

	// return secondSHA[:4]
	return secondSHA[len(secondSHA)-addressChecksumLen:]
}

// PublicKeyBytes get PublicKey []byte
func (wallet Wallet) PublicKeyBytes() []byte {
	pubKey := append(wallet.PublicKey.X.Bytes(), wallet.PublicKey.Y.Bytes()...)
	return pubKey[:]
}

// GetAddress returns wallet address
// Version  Public key hash                           Checksum
// 00       62E907B15CBF27D5425399EBF6F0FB50EBB88F18  C29B7D93
func (wallet Wallet) GetAddress() []byte {
	//1. RIPEMD160(SHA256(PubKey))
	pubKeyHash := HashPubKey(wallet.PublicKeyBytes())

	//2. checksum = SHA256(SHA256(version + pubKeyHash)) 前四个字节
	versionedPayload := append([]byte{version}, pubKeyHash...)
	checksum := checksum(versionedPayload)

	//3. Base58Encode(version + pubKeyHash + checksum)
	fullPayload := append(versionedPayload, checksum...)
	address := Base58Encode(fullPayload)

	return address
}

// ValidateAddress check if address if valid
func ValidateAddress(address string) bool {
	pubKeyHash := Base58Decode([]byte(address))
	actualChecksum := pubKeyHash[len(pubKeyHash)-addressChecksumLen:]
	version := pubKeyHash[0]
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-addressChecksumLen]
	targetChecksum := checksum(append([]byte{version}, pubKeyHash...))

	return bytes.Compare(actualChecksum, targetChecksum) == 0
}

// GetHashPubKey get hashPubKey
func GetHashPubKey(address string) []byte {
	hashPubKey := Base58Decode([]byte(address))
	hashPubKey = hashPubKey[1 : len(hashPubKey)-addressChecksumLen]
	return hashPubKey
}
