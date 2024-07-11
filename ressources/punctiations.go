package goreloaded

func Punctuations(el string) string {
	result := []rune(el)

	for i := 0; i < len(result)-1; i++ {
		if IsPunc(result[i+1]) && result[i] == ' ' {
			result[i+1], result[i] = result[i], result[i+1]
		}
	}
	return string(result)
}

// func Punctuations(el string) string {
//     result := []rune(el)
//     i := 0

//     for i < len(result) {
//         if i > 0 && IsPunc(result[i]) {
//             // Move punctuation left as long as there are spaces before it
//             j := i
//             for j > 0 && result[j-1] == ' ' {
//                 result[j-1], result[j] = result[j], result[j-1]
//                 j--
//             }
//         }
//         i++
//     }

//     return string(result)
// }

func IsPunc(el rune) bool {
	if el == '.' || el == ',' || el == ':' || el == '?' || el == '!' || el == ';' {
		return true
	}
	return false
}