package goreloaded

func Atoan(s []string) []string {
	for i := 0; i < len(s)-1; i++ {
		if s[i] == "a" {
			if IsVowel(s[i+1]) {
				s[i] = "an"
			}
		} else if s[i] == "A" {
			if IsVowel(s[i+1]) {
				s[i] = "An"
			}
		}
	}
	return s
}

func IsVowel(s string) bool {
	if len(s) == 0 {
		return false
	}
	vowels := "aeiouhAEIOUH"
	for _, vowel := range vowels {
		if s[0] == byte(vowel) {
			return true
		}
	}
	return false
}
