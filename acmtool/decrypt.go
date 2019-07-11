package main


import (
	"crypto/cipher"
	"crypto/aes"
	"crypto/md5"
	"encoding/hex"
)


const (
	passphrase = "vvkidtbcjujhtglivdjtlkgtetbtdejlivgukincfhdt"
)



func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}


func decrypt(data []byte) []byte {
	//passphrase := os.Getenv("ACMDIAGNOSTICTOOLTOKEN")
	key := []byte(createHash(passphrase))
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}
	return plaintext
}