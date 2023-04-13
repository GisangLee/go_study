package jay

func reverse(str string) string {
	runes := []rune(str)
	strLen := len(runes)
	reversedRunes := make([]rune, strLen)
	for i := 0; i < strLen; i++ {
		reversedRunes[(strLen-1)-i] = runes[i]
	}
	return string(reversedRunes)
}
