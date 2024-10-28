package reverse

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var models map[byte][]string

func Reverse(args []string) {
	models = ExtractArt() // Initialize the ASCII models
	filename := strings.TrimPrefix(args[0], "--reverse=")
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	scanner := bufio.NewScanner(file)

	var input []string

	for i := 0; i < 8; i++ {
		scanner.Scan()
		input = append(input, scanner.Text())
	}
	result := ReversingArt(input)
	fmt.Println(result)
}

func ExtractArt() map[byte][]string {
	artFile, err := os.Open("art/standard.txt")
	if err != nil {
		fmt.Println("Can't open art file: ", err)
		os.Exit(0)
	}

	// lettersMap stock le modèle de chacune des lettres
	lettersMap := make(map[byte][]string)
	scanner := bufio.NewScanner(artFile)

	currentLetter := ' '
	scanner.Scan()

	// construction de lettersMap.
	for scanner.Scan() {
		var model []string
		for i := 0; i < 8; i++ {
			model = append(model, scanner.Text())
			scanner.Scan()
		}
		lettersMap[byte(currentLetter)] = model
		currentLetter++
	}
	artFile.Close()
	return lettersMap
}

func ReversingArt(input []string) string {
	if len(input[0]) == 0 {
		return ""
	}

	letter := ' '

	for letter <= '~' {
		flag := true
		letterSize := len(models[byte(letter)][0])
		for i := 0; i < 8; i++ {
			if len(input[0]) < LetterSize(byte(letter)) || models[byte(letter)][i] != input[i][:letterSize] {
				flag = false
				break
			}
		}
		if flag {
			var nextInput []string
			for i := 0; i < 8; i++ {
				nextInput = append(nextInput, input[i][letterSize:])
			}
			return string(letter) + ReversingArt(nextInput)
		}

		letter++
	}
	return "Error in the provided art file enter a valid standard art  "
}

func LetterSize(c byte) int {
	return len(models[c][0])
}