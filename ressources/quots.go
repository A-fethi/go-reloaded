package goreloaded

func Quots(s string) string {
	myRune := []rune(s)
	for i, j := 0, len(myRune)-1; i < j; i, j = i+1, j-1 {
		if myRune[i] == '\'' && myRune[i+1] == ' ' {
			myRune = append(myRune[:i+1], myRune[i+2:]...)
		}
		if myRune[j] == '\'' && myRune[j-1] == ' ' {
			myRune = append(myRune[:j-1], myRune[j:]...)
		}
	}
	return string(myRune)
}
