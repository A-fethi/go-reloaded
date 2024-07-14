package goreloaded

func Punctuations(el string) string {
	result := []rune(el)
	i := 0

	for i < len(result) {
		if i > 0 && IsPunc(result[i]) {
			// Move punctuation left as long as there are spaces before it
			j := i
			for j > 0 && result[j-1] == ' ' {
				result[j-1], result[j] = result[j], result[j-1]
				j--
			}
		}
		if i > 0 && i < len(result)-1 && IsPunc(result[i]) && result[i-1] != ' ' && result[i+1] != ' ' {
			result = append(result[:i+1], append([]rune{' '}, result[i+1:]...)...)
			i++
		}
		i++
	}
	return string(result)
}

func IsPunc(el rune) bool {
	if el == '.' || el == ',' || el == ':' || el == '?' || el == '!' || el == ';' {
		return true
	}
	return false
}
