package main

import (
	aes "kaushik1766/AES/AES"
	"os"
)

func main() {
	data, err := os.ReadFile("plaintext.txt")
	if err != nil {
		panic("file not found")
	}

	plaintext := string(data)
	// plaintext := "hello my name is kaushik"
	key := "abcdefghijklmnop"
	aes := aes.NewAES(key)
	// fmt.Println(aes.Encrypt(plaintext))
	err = os.WriteFile("cipher.txt", []byte(aes.Encrypt(plaintext)), 0666)
	if err != nil {
		panic("os error")
	}
}
