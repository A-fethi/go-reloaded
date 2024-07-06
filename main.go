package main

import (
	"fmt"
	"os"
	"strings"

	goreloaded "goreloaded/ressources"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Invalid Arguments")
	} else {
		content, _ := os.ReadFile(os.Args[1])
		str := string(content)
		SplitedStr := strings.Fields(str)
		fmt.Println(SplitedStr)
		for i := 0; i < len(SplitedStr); i++ {
			if SplitedStr[i] == "(hex)" {
				SplitedStr[i-1] = goreloaded.HexConv(SplitedStr[i-1])
				SplitedStr = append(SplitedStr[:i], SplitedStr[i+1:]...)
			} else if SplitedStr[i] == "(bin)" {
				SplitedStr[i-1] = goreloaded.BinConv(SplitedStr[i-1])
				SplitedStr = append(SplitedStr[:i], SplitedStr[i+1:]...)
			} else if SplitedStr[i] == "(up)" {
				SplitedStr[i-1] = strings.ToUpper(SplitedStr[i-1])
				SplitedStr = append(SplitedStr[:i], SplitedStr[i+1:]...)
			} else if SplitedStr[i] == "(up," {
				nb := 0
				nb = IsNumeric(SplitedStr[i+1])
				if i < nb {
					nb = i
					for j := i - 1; j >= nb-i; j-- {
						SplitedStr[j] = strings.ToUpper(SplitedStr[j])
					}
				} else {
					for j := i - 1; j >= i-nb; j-- {
						SplitedStr[j] = strings.ToUpper(SplitedStr[j])
					}
				}
				SplitedStr = append(SplitedStr[:i], SplitedStr[i+2:]...)
			} else if SplitedStr[i] == "(low)" {
				SplitedStr[i-1] = strings.ToLower(SplitedStr[i-1])
				SplitedStr = append(SplitedStr[:i], SplitedStr[i+1:]...)
			} else if SplitedStr[i] == "(cap)" {
				SplitedStr[i-1] = goreloaded.Capitalize(SplitedStr[i-1])
				SplitedStr = append(SplitedStr[:i], SplitedStr[i+1:]...)
			}
		}
		fmt.Println(SplitedStr)
		// fmt.Println(len(SplitedStr))
		// outputFile, _ := os.Create(os.Args[2])
		// defer outputFile.Close()
		// output := os.WriteFile(os.Args[2], content, 0o666)
		// if output != nil {
		// 	fmt.Println(output)
		// 	return
		// }
	}
}

func IsNumeric(s string) int {
	var n int
	for _, i := range s {
		if i >= '0' && i <= '9' {
			n += int(i - '0')
		} else {
			continue
		}
	}
	return n
}

// fmt.Println(splitedStr)

// output := strings.Join(splitedStr, " ")
// err = os.WriteFile(os.Args[2], []byte(output), 0644)
// if err != nil {
// 	fmt.Println("Error writing to file:", err)
// 	return
// }
