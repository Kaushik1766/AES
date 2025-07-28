// Package aes
package aes

import (
	"fmt"
	"sync"
)

type AES struct {
	key []byte
}

func NewAES(key string) AES {
	return AES{[]byte(key)}
}

func (a *AES) Encrypt(plaintext string) string {
	plaintextBytes := []byte(plaintext)
	brokenPlaintext := BreakPlaintext(plaintextBytes)

	var cipher [][4][4]byte = make([][4][4]byte, len(brokenPlaintext))

	var wg sync.WaitGroup

	for i, block := range brokenPlaintext {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			cipher[i] = EncryptBlock(block, a.key)
		}(i)
		// cipher = append(cipher, EncryptBlock(block, a.key))
	}
	wg.Wait()

	result := ""

	for _, val := range cipher {
		for i := range 4 {
			for j := range 4 {
				result += fmt.Sprintf("%02x", val[i][j])
			}
		}
	}
	return result
}

func (a *AES) Decrypt(cipher string) string {
	return ""
}
