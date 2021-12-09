package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const DAYS = 256

func main() {
	ages := readFile("input.txt")

	status := make([]int, 9)
	fishes := 0

	for _, age := range ages {
		status[age]++
		fishes++
	}

	for day := DAYS; day > 0; day-- {
		fishes += status[0]
		status = append(status, status[0])
		status[7] += status[0]
		status = status[1:]

	}

	fmt.Println("At day", DAYS, "there will be", fishes, "lanterfishes.")
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
