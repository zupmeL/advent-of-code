package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	rows := readFile("input.txt")

	res := 0

	parenthesis := make(map[rune]string)
	parenthesis['('] = ")"
	parenthesis['['] = "]"
	parenthesis['{'] = "}"
	parenthesis['<'] = ">"

	costs := make(map[rune]int)
	costs[')'] = 3
	costs[']'] = 57
	costs['}'] = 1197
	costs['>'] = 25137

	for _, str := range rows {
		backlog := ""
		for _, char := range str {
			if strings.Contains("<{[(", string(char)) {
				backlog += parenthesis[char]
			} else {
				if byte(char) == backlog[len(backlog)-1] {
					backlog = backlog[:len(backlog)-1]
				} else {
					res += costs[char]
					break
				}
			}
		}
	}

	fmt.Println("Result is:", res)

}

func readFile(input string) []string {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	rows := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rows = append(rows, scanner.Text())
	}

	return rows
}
