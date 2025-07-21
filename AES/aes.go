// Package aes
package aes

import "fmt"

type AES struct {
	key []byte
}

func NewAES(key string) AES {
	return AES{[]byte(key)}
}

func (a *AES) Encrypt(plaintext string) string {
	plaintextBytes := []byte(plaintext)
	brokenPlaintext := BreakPlaintext(plaintextBytes)

	var cipher [][4][4]byte

	for _, block := range brokenPlaintext {
		cipher = append(cipher, EncryptBlock(block, a.key))
	}

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
