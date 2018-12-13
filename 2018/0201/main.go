package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var doubles, triples int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		chars := map[string]int{}
		doubleFound := true
		tripleFound := true

		for _, char := range scanner.Text() {
			chars[string(char)]++
		}

		for _, counter := range chars {
			if counter == 2 {
				doubleFound = true
			} else if counter == 3 {
				tripleFound = true
			}
		}

		if doubleFound {
			doubles++
		}
		if tripleFound {
			triples++
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Doubles: ", doubles)
	fmt.Println("Triples: ", triples)
	fmt.Println("CRC: ", doubles*triples)
}
