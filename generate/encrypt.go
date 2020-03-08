package generate

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

func Encrypt(compressedString string) ([]byte, []byte) {

	// Create random default encryption key
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		fmt.Println("Error generating random AES encryption key")
	}

	c, err := aes.NewCipher(key)
	// if there are any errors, handle them
	if err != nil {
		fmt.Println(err)
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		fmt.Println(err)
	}

	// creates a new byte array the size of the nonce
	// which must be passed to Seal
	nonce := make([]byte, gcm.NonceSize())

	// random sequence
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Println(err)
	}

	encryptedString := gcm.Seal(nonce, nonce, []byte(compressedString), nil)

	return encryptedString, key
}
