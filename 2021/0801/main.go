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
	occ1478 := 0
	for _, row := range rows {
		parts := strings.Split(row, " | ")
		numbers := strings.Split(parts[1], " ")
		for _, number := range numbers {
			switch len(number) {
			case 2, 3, 4, 7:
				occ1478++
			default:
			}
		}
	}

	fmt.Println("Occurrencies of digits 1, 4, 7, or 8:", occ1478)

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
