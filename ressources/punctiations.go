package goreloaded

// import "strings"

func Punctuations(s string) string {
	result := []rune(s)

	for i := 0; i < len(result)-1; i++ {
		if IsPunc(result[i+1]) && result[i] == ' ' {
			result[i+1], result[i] = result[i], result[i+1]
		}
	}
	return string(result)
	//return strings.Join(strings.Fields(string(result)), " ")
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
