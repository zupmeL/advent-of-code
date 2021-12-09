package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	positions := readFile("input.txt")
	floor, ceil := getMeans(positions)

	ceilMovs := 0
	floorMovs := 0
	for _, pos := range positions {
		ceilMovs += price(abs(pos - ceil))
		floorMovs += price(abs(pos - floor))
	}
	minMovs := ceilMovs
	if ceilMovs > floorMovs {
		minMovs = floorMovs
	}
	fmt.Println("Minimun movements are:", minMovs)
}

func getMeans(arr []int) (int, int) {
	total := 0
	n := len(arr)

	for _, element := range arr {
		total += (element)
	}
	floor := (total / n)
	return floor, floor + 1
}

func price(x int) int {
	return x * (x + 1) / 2
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
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
