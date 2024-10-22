package reverse

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Function to get the appropriate art file based on the banner type
func GetArtFile(banner string) (string, error) {
	standard := "./art/standard.txt"
	shadow := "./art/shadow.txt"
	thinkertoy := "./art/thinkertoy.txt"

	switch banner {
	case "standard":
		return standard, nil
	case "shadow":
		return shadow, nil
	case "thinkertoy":
		return thinkertoy, nil
	default:
		return "", fmt.Errorf("invalid banner. Please choose from: standard, shadow, thinkertoy\n ")
	}
}

func ReadArtFile(art string) ([]string, error) {
	file, err := os.Open(art)
	if err != nil {
		return nil, fmt.Errorf("error opening the file: %v\n ", err)
	}
	defer file.Close()

	var asciiArt []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		asciiArt = append(asciiArt, scanner.Text())
	}
	if len(asciiArt) != 855 {
		return nil, fmt.Errorf("file error: expected 855 lines, but got %d lines", len(asciiArt))
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading the file: %v\n ", err)
	}

	return asciiArt, nil
}

// Function to generate ASCII art from the input string
func GenerateASCIIArt(input string, asciiArt []string) string {
	var result string
	lines := strings.Split(input, "\\n")
count:=0
	for _, line := range lines {
		if line == "" {
			result += "\n"
			count++
			continue
		}
		for i := 1; i <= 8; i++ {
			for _, r := range line {
				if r < 32 || r > 126 {
					fmt.Println("invalid character: please enter a valid character between ASCII code 32 and 126")
					return ""
				}
				index := 9*(int(r)-32) + i
				result += asciiArt[index]
			}
			result += "\n"

		}
	}
	if count == len(result) {
		result = result[:len(result)-1]
	}

	return result
}
func GenerateASCIIArtLetter(input string, asciiArt []string, targetWord string, colorCode string, resetCode string) string {
	var result string
	lines := strings.Split(input, "\\n")
count:=0

	// Process each line
	for _, line := range lines {
		if line == "" {
			result += "\n"
			count++

			continue
		}

		// Loop through each line of ASCII art (8 lines per character)
		for i := 1; i <= 8; i++ {
			for _, r := range line {
				if r < 32 || r > 126 {
					fmt.Printf("invalid character: please enter a valid character between ASCII code 32 and 126")
					return ""
				}
				index := 9*(int(r)-32) + i // Get the index of the ASCII art line for the character

				// Check if the current character matches part of the targetWord
				if strings.ContainsRune(targetWord, r) {
					// Apply the color to the matched character
					result += colorCode + asciiArt[index] + resetCode
				} else {
					// No match, just add the ASCII art without coloring
					result += asciiArt[index]
				}
			}
			result += "\n" // Add newline after each set of ASCII lines
		}
	}
	if count == len(result) {
		result = result[:len(result)-1]
	}

	return result
}

// Read the ASCII Art from a file
func ReadASCIIArtFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening the file: %v", err)
	}
	defer file.Close()

	var asciiArt []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		asciiArt = append(asciiArt, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading the file: %v", err)
	}

	return asciiArt, nil
}

// LoadASCIIAlphabet reads an ASCII art file and returns a map of letters to their ASCII representations
func LoadASCIIAlphabet(filename string) (map[string][]string, error) {
	asciiAlphabet := make(map[string][]string)
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var letter string
	var artLines []string

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("Reading line:", line)

		// If we hit a new letter, store the previous one
		if len(line) == 0 && letter != "" {
			if len(artLines) > 0 {
				asciiAlphabet[letter] = append(asciiAlphabet[letter], artLines...)
				fmt.Println("Adding to asciiAlphabet:", letter, artLines)
			}
			letter = ""
			artLines = nil
		} else if len(line) > 0 {
			// Store the letter when it's encountered
			if letter == "" {
				letter = line // Assume the first line is the letter
			} else {
				artLines = append(artLines, line) // Add the art lines
			}
		}
	}

	// Don't forget to add the last character if the file doesn't end with a blank line
	if letter != "" && len(artLines) > 0 {
		asciiAlphabet[letter] = append(asciiAlphabet[letter], artLines...)
		fmt.Println("Adding last to asciiAlphabet:", letter, artLines)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return asciiAlphabet, nil
}

// ConvertArtToText converts ASCII art lines back to the original string
func ConvertArtToText(asciiArt []string, asciiAlphabet map[string][]string) string {
	var result string
	numLines := len(asciiArt)

	if numLines == 0 || numLines%8 != 0 {
		fmt.Println(fmt.Errorf(" Invalid ASCII art input. "))
		return ""
	}

	// Process each character in chunks of 8 lines
	for i := 0; i < numLines; i += 8 {
		if i+8 > numLines {
			fmt.Println(fmt.Errorf(" Incomplete character art for last segment. "))
			return ""
		}
		
		characterArt := asciiArt[i : i+8]
		fmt.Println("Processing character art:", characterArt)

		found := false
		for letter, art := range asciiAlphabet {
			match := true
			for line := 0; line < 8; line++ {
				fmt.Printf("Comparing with letter: %s, line: %s with %s\n", letter, characterArt[line], art[line])
				if characterArt[line] != art[line] {
					match = false
					break
				}
			}
			if match {
				result += letter
				found = true
				break
			}
		}

		// If no match is found, return an error
		if !found {
			fmt.Println("Character Art:", characterArt)
			fmt.Println("ASCII Alphabet:", asciiAlphabet)
			fmt.Println(fmt.Errorf(" Error: Letter not found in art file. The art must be in the standard format. "))
			return ""
		}
	}

	return result
}