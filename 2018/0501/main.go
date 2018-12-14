package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

func main() {
	input := readFile("input.txt")

	length := len(input)
	for index := 1; index < length; index++ {
		isCurrLowerCase := unicode.IsLower(rune(input[index]))
		isLastLowerCase := unicode.IsLower(rune(input[index-1]))

		if isCurrLowerCase == !isLastLowerCase && strings.ToUpper(string(input[index-1])) == strings.ToUpper(string(input[index])) {
			input = removeFromTo(input, index-1, index)
			length = len(input)
			if index <= 2 {
				index = 0
			} else {
				index = index - 2
			}
		}
	}
	fmt.Printf("Only %d units remaining\n", len(input))
}

func removeFromTo(str string, startIndex, endIndex int) string {
	slices := []byte(str)
	slices = append(slices[:startIndex], slices[endIndex+1:]...)
	return string(slices)
}

func readFile(input string) string {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var content string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		content = scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return content
}
