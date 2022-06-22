package reverse_string

func ReverseString(input string) (output string) {
	bytes := []rune(input)
	var reverseRunes []rune
	for i := len(bytes); i > 0; i-- {
		reverseRunes = append(reverseRunes, bytes[i-1])
	}

	return string(reverseRunes)
}
