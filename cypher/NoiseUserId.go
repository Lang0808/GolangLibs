package cypher

import "errors"

var CodeBoard = [][]int{
	{8, 26}, {36, 54}, {25, 38}, {56, 57}, {10, 41}, {48, 40},
	{13, 51}, {37, 7}, {3, 47}, {2, 62}, {44, 59}, {27, 50},
	{12, 15}, {9, 19}, {52, 11}, {63, 17}, {42, 60}, {5, 31}, {55, 21},
	{24, 28}, {46, 6}, {45, 30}, {23, 53}, {14, 32}, {39, 34}, {58, 22}, {43, 33},
	{16, 49}, {4, 1}, {20, 29}, {35, 18}, {0, 61},
}

func NoiseUserId(userId int32) int64 {
	ans := int64(0)
	for i := 0; i < 32; i = i + 1 {
		bit := GetBit32(userId, i)
		if bit == 1 {
			indexes := CodeBoard[i]
			for j := 0; j < len(indexes); j = j + 1 {
				ans = OnBit64(ans, indexes[j])
			}
		}
	}
	return ans
}

func DenoiseUserId(noisedUserId int64) (int32, error) {
	ans := int32(0)
	for i := 0; i < 32; i = i + 1 {
		indexes := CodeBoard[i]
		index := indexes[0]
		bit := GetBit64(noisedUserId, index)
		for j := 0; j < len(indexes); j = j + 1 {
			bit2 := GetBit64(noisedUserId, indexes[j])
			if bit2 != bit {
				return -1, errors.New("Error occurs when denoise user id")
			}
		}
		if bit == 1 {
			ans = OnBit32(ans, i)
		}
	}
	return ans, nil
}
