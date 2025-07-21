package main

import (
	"fmt"
	aes "kaushik1766/AES/AES"
)

func main() {
	plaintext := "hello my name is kaushik"
	key := "abcdefghijklmnop"
	aes := aes.NewAES(key)
	fmt.Println(aes.Encrypt(plaintext))
}
