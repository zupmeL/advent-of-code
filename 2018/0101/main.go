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
	sum := 0
	for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal("Couldn't convert scan to number")
		}
		sum = sum + number
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("The total sum is: ", sum)
}
