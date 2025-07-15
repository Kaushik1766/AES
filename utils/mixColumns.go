package utils

func MixColumns(block [4][4]byte) [4][4]byte {

	var ans [4][4]byte

	for col := range 4 {
		ans[0][col] = mul2[block[0][col]] ^ mul3[block[1][col]] ^ block[2][col] ^ block[3][col]
		ans[1][col] = block[0][col] ^ mul2[block[1][col]] ^ mul3[block[2][col]] ^ block[3][col]
		ans[2][col] = block[0][col] ^ block[1][col] ^ mul2[block[2][col]] ^ mul3[block[3][col]]
		ans[3][col] = mul3[block[0][col]] ^ block[1][col] ^ block[2][col] ^ mul2[block[3][col]]

	}
	return ans
}
