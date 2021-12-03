package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	input := readFile("input.txt")

	gamma, epsilon := getGammaAndEpsilon(input)

	fmt.Println("Gamma is:", gamma)
	fmt.Println("Epsilon is:", epsilon)
	fmt.Println("Base 10 Multiplication is:", binaryToInt64(gamma)*binaryToInt64(epsilon))
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

func getGammaAndEpsilon(input []string) (string, string) {
	gamma := ""
	epsilon := ""
	zeroes := make([]int, len(input[0]))

	for _, row := range input {
		for i, rune := range row {
			if rune == '0' {
				zeroes[i]++
			}
		}
	}

	for i := 0; i < len(zeroes); i++ {
		if zeroes[i] > len(input)/2 {
			gamma += "0"
			epsilon += "1"
		} else {
			gamma += "1"
			epsilon += "0"
		}
	}

	return gamma, epsilon
}

func binaryToInt64(binary string) int64 {
	i, err := strconv.ParseInt(binary, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	return i
}
