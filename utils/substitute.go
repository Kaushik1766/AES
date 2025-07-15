package utils

func SubstituteMatrix(block [4][4]byte) [4][4]byte {
	for i := range 4 {
		for j := range 4 {
			block[i][j] = sbox[block[i][j]]
		}
	}
	return block
}
