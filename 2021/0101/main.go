package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	firstItem := getFirstItem(scanner)
	increments := getIncrements(scanner, firstItem)

	fmt.Println("INCREMENTS: ", increments)
}

func getFirstItem(scanner *bufio.Scanner) int {
	scanner.Scan()
	item, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal("Couldn't convert scanned row to number")
	}
	return item
}

func getIncrements(scanner *bufio.Scanner, previousItem int) int {
	if scanner.Scan() {
		item, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal("Couldn't convert scanned row to number")
		}
		if previousItem < item {
			return getIncrements(scanner, item) + 1
		} else {
			return getIncrements(scanner, item)
		}
	}
	return 0
}
