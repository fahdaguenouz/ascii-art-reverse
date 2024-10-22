package reverse

import (
	"fmt"
	"strings"
)

func AsciiColor(args []string) string {
	// Map of available colors
	colors := map[string]string{
		"red":     "\033[31m",
		"green":   "\033[32m",
		"blue":    "\033[34m",
		"yellow":  "\033[33m",
		"cyan":    "\033[36m",
		"magenta": "\033[35m",
		"white":   "\033[37m",
		"black":   "\033[30m",
		"gray":    "\033[90m",
		"purple":  "\033[35m",
		"orange":  "\033[38;5;214m",
		"reset":   "\033[0m",
	}

	// Default to standard banner
	banner := "standard"
	input := ""
	colorname := ""
	word := ""

	// Handle argument cases
	switch len(args) {
	case 1:
		// Case: Only input provided
		input = args[0]
	case 2:
		// Case: Input with color, default banner
		color := args[0]
		input = args[1]

		colorname = strings.TrimPrefix(color, "--color=")

		// Validate color name
		if _, exists := colors[colorname]; !exists {
			return "Invalid color name.\n "
		}
	case 3:
		// Case: Input with color and specific banner
		color := args[0]
		input = args[1]
		banner = args[2]

		if banner != "standard" && banner != "shadow" && banner != "thinkertoy" {
			banner = "standard"
			input=args[2]
			word=args[1]
			colorname = strings.TrimPrefix(color, "--color=")

			// Validate color name
			if _, exists := colors[colorname]; !exists {
				return "Invalid color name.\n "
			}
			artFile, err := GetArtFile(banner)
			if err != nil {
				fmt.Println(err)
				return ""
			}

			asciiArt, err := ReadArtFile(artFile)
			if err != nil {
				fmt.Println(err)
				return ""
			}

			// Generate the colored ASCII art for the specific word
			result := GenerateASCIIArtLetter(input, asciiArt, word, colors[colorname], colors["reset"])
			return result
		} else {
			colorname = strings.TrimPrefix(color, "--color=")

			// Validate color name
			if _, exists := colors[colorname]; !exists {
				return "Invalid color name.\n "
			}

			// Validate banner
			if banner != "standard" && banner != "thinkertoy" && banner != "shadow" {
				return "Invalid banner name.\n "
			}

		}

	case 4:
		// Case: Input with color, word to be colored, and specific banner
		color := args[0]
		word = args[1]
		input = args[2]
		banner = args[3]

		// Extract and validate color name
		colorname = strings.TrimPrefix(color, "--color=")
		if _, exists := colors[colorname]; !exists {
			return "Invalid color name.\n "
		}

		// Validate banner
		if banner != "standard" && banner != "thinkertoy" && banner != "shadow" {
			return "Invalid banner name.\n "
		}

		artFile, err := GetArtFile(banner)
		if err != nil {
			fmt.Println(err)
			return ""
		}

		asciiArt, err := ReadArtFile(artFile)
		if err != nil {
			fmt.Println(err)
			return ""
		}

		// Generate the colored ASCII art for the specific word
		result := GenerateASCIIArtLetter(input, asciiArt, word, colors[colorname], colors["reset"])
		return result
	default:
		return "Usage: go run . [OPTION] [STRING]\nEX: go run . --color=<color> <substring to be colored> 'something'\n "
	}


	artFile, err := GetArtFile(banner)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	asciiArt, err := ReadArtFile(artFile)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	// Generate ASCII art
	result := GenerateASCIIArt(input, asciiArt)

	// Apply color if a valid color is provided
	if colorname != "" {
		return colors[colorname] + result + colors["reset"]
	}

	return result
}
