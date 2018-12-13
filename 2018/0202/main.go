package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	input := readFile("input.txt")

	line1, line2 := linesAtDistanceOne(input)
	if line1 != "" {
		fmt.Println("The result string at distance one is: ", equalChars(line1, line2))
	} else {
		fmt.Println("No lines at distance one.")
	}
}

func equalChars(str1, str2 string) string {
	var res string
	for index := 0; index < len(str1); index++ {
		if str1[index] == str2[index] {
			res += string(str1[index])
		}
	}
	return res
}

func linesAtDistanceOne(lines []string) (string, string) {
	rest := 1
	for _, line := range lines {
		for index := rest; index < len(lines); index++ {
			if distanceOne(line, lines[index]) {
				fmt.Println(line, " ", lines[index])
				return line, lines[index]
			}
		}
		rest++
	}
	return "", ""
}

func distanceOne(str1, str2 string) bool {
	distance := 0
	for index := 0; index < len(str1) && distance <= 1; index++ {
		if str1[index] != str2[index] {
			distance++
		}
	}
	if distance != 1 {
		return false
	}
	return true
}

func readFile(input string) []string {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var content []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		content = append(content, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return content
}
