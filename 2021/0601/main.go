package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const DAYS = 80

func main() {
	ages := readFile("input.txt")

	day := 0

	for day < DAYS {
		currentFishes := len(ages)
		for index := 0; index < currentFishes; index++ {
			if ages[index] == 0 {
				ages = append(ages, 8)
				ages[index] = 6
			} else {
				ages[index]--
			}
		}
		day++
	}

	fmt.Println("At day", day, "there will be", len(ages), "lanterfishes.")
}

func readFile(input string) []int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	ages := make([]int, 0)
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	row := scanner.Text()
	for _, item := range strings.Split(row, ",") {
		age, err := strconv.Atoi(item)
		if err != nil {
			log.Fatal(err)
		}
		ages = append(ages, age)
	}

	return ages
}
