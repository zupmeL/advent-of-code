package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

const charA = 65
const diffToLower = 32

func main() {
	input := readFile("input.txt")

	charRemoved := ""
	minLength := len(input)
	for index := charA; index < charA+26; index++ {
		workingString := strings.Replace(input, string(index), "", -1)
		workingString = strings.Replace(workingString, string(index+diffToLower), "", -1)
		workingString = react(workingString)
		if minLength > len(workingString) {
			minLength = len(workingString)
			charRemoved = string(index)
		}
	}
	fmt.Printf("Best unit to remove is %s that result in %d units remaining.\n", charRemoved, minLength)
}

func react(input string) string {
	length := len(input)
	for index := 1; index < length; index++ {
		isCurrLowerCase := unicode.IsLower(rune(input[index]))
		isLastLowerCase := unicode.IsLower(rune(input[index-1]))
		currUpperCase := strings.ToUpper(string(input[index]))
		lastUpperCase := strings.ToUpper(string(input[index-1]))

		if isCurrLowerCase == !isLastLowerCase && currUpperCase == lastUpperCase {
			input = removeFromTo(input, index-1, index)
			length = len(input)
			if index <= 2 {
				index = 0
			} else {
				index = index - 2
			}
		}
	}
	return input
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
