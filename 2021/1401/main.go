package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

const STEPS = 10

func main() {
	sequence, instructions := readFile("input.txt")
	occurrences := make(map[rune]int)

	for _, char := range sequence {
		occurrences[char]++
	}

	for step := 0; step < STEPS; step++ {
		newSequence := ""
		for i := 0; i < len(sequence)-1; i++ {
			replacement := instructions[sequence[i:i+2]]
			if i == 0 {
				newSequence = replacement
			} else {
				newSequence = newSequence[:len(newSequence)-1] + replacement
			}
			occurrences[rune(replacement[1])]++
		}
		sequence = newSequence
	}

	maxOcc, minOcc := 0, math.MaxInt
	for _, n := range occurrences {
		if n > maxOcc {
			maxOcc = n
		}
		if n < minOcc {
			minOcc = n
		}
	}

	fmt.Println("MAX - MIN( ", maxOcc, " - ", minOcc, " ): ", maxOcc-minOcc)
}

func readFile(input string) (string, map[string]string) {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	sequence := scanner.Text()
	scanner.Scan()

	instructions := make(map[string]string, 0)
	for scanner.Scan() {
		row := scanner.Text()

		s := strings.Split(row, " -> ")
		replacement := string(s[0][0]) + string(s[1][0]) + string(s[0][1])
		instructions[s[0]] = replacement
	}

	return sequence, instructions
}
