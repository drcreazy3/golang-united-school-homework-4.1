package reverse_string

func ReverseString(input string) (output string) {
	bytes := []byte(input)
	var reverseBytes []byte
	for i := len(bytes); i > 0; i-- {
		reverseBytes = append(reverseBytes, bytes[i-1])
	}

	str := ""
	for _, r := range reverseBytes {
		str += string(r)
	}

	return str
}
