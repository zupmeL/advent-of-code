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

	oxygen := getRating(input, 0, oxygenCheck)
	co2 := getRating(input, 0, co2Check)

	fmt.Println("Oxygen generation rating is:", oxygen)
	fmt.Println("CO2 scrubber rating is:", co2)
	fmt.Println("Base 10 Multiplication is:", binaryToInt64(oxygen)*binaryToInt64(co2))
}

func mostCommonValueAtIndex(input []string, index int) rune {
	zeroes := 0
	for _, row := range input {
		if row[index] == '0' {
			zeroes++
		}
	}
	if zeroes > len(input)/2 {
		return '0'
	}
	return '1'
}

func oxygenCheck(input byte, mcv byte) bool {
	return input == mcv
}

func co2Check(input byte, mcv byte) bool {
	return !(input == mcv)
}

func getRating(input []string, index int, check func(byte, byte) bool) string {
	mcv := mostCommonValueAtIndex(input, index)

	newInput := make([]string, 0)

	for _, row := range input {
		if check(row[index], byte(mcv)) {
			newInput = append(newInput, row)
		}
	}

	if len(newInput) > 1 {
		return getRating(newInput, index+1, check)
	}
	return newInput[0]
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

func binaryToInt64(binary string) int64 {
	i, err := strconv.ParseInt(binary, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	return i
}
