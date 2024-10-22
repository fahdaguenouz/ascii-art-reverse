package reverse

import (
	"fmt"
	"strings"
)

func Reverse(args []string) {
	input := args[0]
	filename := strings.TrimPrefix(input, "--reverse=")
	if err := validateFilename(filename); err != nil {
		fmt.Println(err)
		return
	}
	// Read the ASCII art from the file
	asciiArt, err := ReadASCIIArtFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	
	// Load the ASCII alphabet
	asciiAlphabet, err := LoadASCIIAlphabet("./art/standard.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Convert ASCII art back to text
	result := ConvertArtToText(asciiArt, asciiAlphabet)	
	fmt.Println(result)

}
