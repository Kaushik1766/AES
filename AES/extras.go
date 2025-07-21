package aes

import "kaushik1766/AES/utils"

func EncryptBlock(plaintextBlock [4][4]byte, keyBytes []byte) [4][4]byte {
	roundResult := InitialRound(plaintextBlock, keyBytes)
	roundKey := utils.InitialRoundKey(keyBytes)

	for roundNumber := range 9 {
		roundKey = utils.NextKey(roundKey, roundNumber+1)
		roundResult = Round(roundResult, roundKey)
	}

	// last round
	roundKey = utils.NextKey(roundKey, 10)
	roundResult = LastRound(roundResult, roundKey)

	return roundResult
}

func BreakPlaintext(plaintext []byte) [][4][4]byte {
	var ans [][4][4]byte

	var cur []byte

	for i, val := range plaintext {
		if i%16 == 0 && i != 0 {
			ans = append(ans, utils.ConvertToGrid(cur))
			cur = []byte{val}
		} else {
			cur = append(cur, val)
		}
	}

	remLength := 16 - len(cur)
	paddingChar := byte(remLength)

	for range remLength {
		cur = append(cur, paddingChar)
	}

	ans = append(ans, utils.ConvertToGrid(cur))
	return ans
}

func InitialRound(plaintext [4][4]byte, key []byte) [4][4]byte {
	initialKey := utils.InitialRoundKey(key)

	addKey := utils.AddKey(plaintext, initialKey)

	return addKey
}

func Round(cip [4][4]byte, roundKey [4][4]byte) [4][4]byte {
	subBytes := utils.SubstituteMatrix(cip)
	shiftRows := utils.ShiftRows(subBytes)
	mixColumns := utils.MixColumns(shiftRows)
	addKey := utils.AddKey(mixColumns, roundKey)

	return addKey
}

func LastRound(cip [4][4]byte, roundKey [4][4]byte) [4][4]byte {
	cip = utils.SubstituteMatrix(cip)
	cip = utils.ShiftRows(cip)
	cip = utils.AddKey(cip, roundKey)
	return cip
}
