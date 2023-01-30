package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Expected some text to asciize!")
		fmt.Println(`Usage: ./asciize "your-text-here"`)
		os.Exit(1)
	}

	var words []string

	for i := 1; i < len(os.Args); i++ {
		words = append(words, os.Args[i])
	}

	for i := 0; i < len(words); i++ {
		var chars []string
		for j := 0; j < len(words[i]); j++ {
			chars = append(chars, words[i][j:j+1])
		}
		for k := 0; k < len(chars); k++ {
			bytesRead, _ := os.ReadFile("letters/" + chars[k] + ".txt")
			fileContent := string(bytesRead)
			lines := strings.Split(fileContent, "\n")

			for l := 0; l < len(lines); l++ {
				fmt.Println(lines[l])
			}

			fmt.Print("\n\n")
		}
	}
}
