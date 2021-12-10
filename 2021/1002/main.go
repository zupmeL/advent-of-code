package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	rows := readFile("input.txt")

	parenthesis := make(map[rune]rune)
	parenthesis['('] = ')'
	parenthesis['['] = ']'
	parenthesis['{'] = '}'
	parenthesis['<'] = '>'

	costs := make(map[rune]int)
	costs[')'] = 1
	costs[']'] = 2
	costs['}'] = 3
	costs['>'] = 4

	scores := make([]int, 0)
	for _, str := range rows {
		backlog := make([]rune, 0)
		uncompleted := true
		for _, char := range str {
			if strings.Contains("<{[(", string(char)) {
				backlog = append([]rune{parenthesis[char]}, backlog...)
			} else {
				if char == backlog[0] {
					backlog = backlog[1:]
				} else {
					uncompleted = false
					break
				}
			}
		}
		if uncompleted {
			score := 0
			for _, char := range backlog {
				score *= 5
				score += costs[char]
			}
			scores = append(scores, score)
		}
	}

	sort.Ints(scores)
	fmt.Println("Middle score is:", scores[len(scores)/2])

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
