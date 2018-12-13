package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	numbers := readInput("input.txt")
	sum := 0
	watcher := map[int]bool{0: true}
	for {
		for _, number := range numbers {
			sum = sum + number
			if watcher[sum] {
				fmt.Printf("Reached %d twice.\n", sum)
				return
			}
			watcher[sum] = true
		}
	}
}

func readInput(input string) []int {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var numbers []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal("Couldn't convert scan to number")
		}
		numbers = append(numbers, number)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return numbers
}
