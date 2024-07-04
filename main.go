package main

import (
	"fmt"
	"os"

	goreloaded "goreloaded/ressources"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Invalid Arguments")
	} else {
		content, _ := os.ReadFile(os.Args[1])
		str := string(content)
		SplitedStr := goreloaded.SplitWhiteSpaces(str)
		fmt.Println(SplitedStr)
		// fmt.Println(len(SplitedStr))
		// outputFile, _ := os.Create(os.Args[2])
		// defer outputFile.Close()
		output := os.WriteFile(os.Args[2], content, 0o666)
		if output != nil {
			fmt.Println(output)
			return
		 }
	}
}
