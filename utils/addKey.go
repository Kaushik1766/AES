package utils

func AddKey(block [4][4]byte, roundKey [4][4]byte) [4][4]byte {
	var ans [4][4]byte
	for i := range 4 {
		for j := range 4 {
			ans[i][j] = block[i][j] ^ roundKey[i][j]
		}
	}
	return ans
}
