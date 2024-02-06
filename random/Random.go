package random

import (
	"math/rand"
)

func RandomInt32(cardinality int32) int32 {
	return rand.Int31() % cardinality
}

func RandomPercents(percents []int32) int32 {
	sum := (int32)(0)
	for i := 0; i < len(percents); i++ {
		sum += percents[i]
	}
	prefixSum := RandomInt32(sum)
	for i := 0; i < len(percents); i++ {
		prefixSum -= percents[i]
		if prefixSum < 0 {
			return (int32)(i)
		}
	}
	return -1
}

func RandomString(LEN int) string {
	ALPHABET_SIZE := 26
	ans := ""
	for i := 0; i < LEN; i++ {
		ind := RandomInt32(int32(ALPHABET_SIZE))
		ans = ans + string(ind+'a')
	}
	return ans
}

func RandomDouble() float64 {
	return rand.Float64()
}

func RandomRangeInt64(min int64, max int64) int64 {
	return rand.Int63n(max-min) + min
}
