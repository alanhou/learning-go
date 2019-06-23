package main

func lengthOfNonRepeatingSubStr(s string) int{
	//lastOccurred := make(map [byte] int)
	lastOccurred := make(map [rune] int)
	start := 0
	maxLength := 0

	//for i, ch := range []byte(s) {
	for i, ch := range []rune(s) {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i - start + 1 > maxLength{
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}

	return maxLength
}


