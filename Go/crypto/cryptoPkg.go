package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"os"
	"golang.org/x/crypto/argon2"
)

func RestCryptor(phrase []byte) {
	salt := make([]byte, 32)
	rand.Read(salt)
	uk := argon2.IDKey(phrase, salt, 1, 2*1024*1024, 4, 32)
	mk := make([]byte, 32)
	rand.Read(mk)
	cipherBlock, err := aes.NewCipher(uk)
	if err != nil {
		panic(err)
	}
	gcm, err := cipher.NewGCMWithRandomNonce(cipherBlock)
	if err != nil {
		panic(err)
	}
	sealed := gcm.Seal(nil, nil, mk, nil)
	data, err := os.ReadFile(~/.biovault/disk_key.key)
	if err != nil {
		panic(err)
		}
	if len(data) == 0 {
		diskKey := make([]byte, 32)
		rand.Read(diskKey)
		os.WriteFile("~/.biovault/disk_key.key", 
	}
}



func main() {
	RestCryptor([]byte("my-test-passphrase"))
}


