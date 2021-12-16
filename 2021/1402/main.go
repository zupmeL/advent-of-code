package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

const STEPS = 40

func main() {
	sequence, instructions := readFile("input.txt")

	tracking := make(map[string]int)
	for i := 0; i < len(sequence)-1; i++ {
		tracking[string(sequence[i])+string(sequence[i+1])]++
	}

	for step := 0; step < STEPS; step++ {
		newTracking := make(map[string]int)
		for pattern, amount := range tracking {
			newTracking[string(pattern[0])+instructions[pattern]] += amount
			newTracking[instructions[pattern]+string(pattern[1])] += amount
		}
		tracking = newTracking
	}

	occurrences := make(map[byte]int)
	occurrences[sequence[len(sequence)-1]]++ // last character
	for pattern, amount := range tracking {
		occurrences[pattern[0]] += amount
	}

	maxOcc, minOcc := 0, math.MaxInt
	for _, occ := range occurrences {
		if occ < minOcc {
			minOcc = occ
		}
		if occ > maxOcc {
			maxOcc = occ
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
		instructions[s[0]] = s[1]
	}

	return sequence, instructions
}
