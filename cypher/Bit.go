package cypher

func OnBit64(a int64, pos int) int64 {
	temp := int64(1) << pos
	return a | temp
}

func OnBit32(a int32, pos int) int32 {
	temp := int32(1) << pos
	return a | temp
}

func GetBit32(a int32, pos int) int32 {
	a = a >> pos
	return a & 1
}

func GetBit64(a int64, pos int) int64 {
	a = a >> pos
	return a & int64(1)
}
