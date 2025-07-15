package utils

func ShiftRows(block [4][4]byte) [4][4]byte {
	var shiftedBlock [4][4]byte

	for i := range 4 {
		for j := range 4 {
			shiftedBlock[i][j] = block[i][(j+i)%4]
		}
	}

	return shiftedBlock
}
