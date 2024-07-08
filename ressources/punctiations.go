package goreloaded

// import "strings"

func Punctuations(s string) string {
	result := []rune(s)

	for i := 0; i < len(result)-1; i++ {
		if i > 0 && IsPunc(result[i-1]) && IsAlpha(result[i]) {
			// result[i], result[i+1] = result[i+1], result[i]
			result = append(result[:i], append([]rune{' '}, result[i:]...)...)
			// i++
		} else if IsPunc(result[i+1]) && result[i] == ' ' {
			result[i+1], result[i] = result[i], result[i+1]
		}
	}
	return string(result)
	// return strings.Join(strings.Fields(string(result)), " ")
	// result := []rune(s)
	// for i := 0; i < len(result)-1; i++ {
	// 	if IsPunc(result[i+1]) && result[i] == ' ' {
	// 		result[i+1], result[i] = result[i], result[i+1]
	// 	}
	// }
	// return strings.Join(strings.Fields(string(result)), " ")
}

func IsPunc(r rune) bool {
	if r == '.' || r == ',' || r == '!' || r == ':' || r == ';' || r == '?' {
		return true
	}
	return false
}

func IsAlpha(r rune) bool {
	if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') && r == ' ' {
		return false
	}
	return true
}
