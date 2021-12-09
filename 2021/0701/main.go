package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	positions := readFile("input.txt")
	median := getMedian(positions)

	movements := 0
	for _, pos := range positions {
		movements += abs(pos - median)
	}

	fmt.Println("Minimun movements are:", movements)
}

func getMedian(arr []int) int {
	sort.Ints(arr)
	n := len(arr)

	if n%2 == 0 {
		return arr[n/2]
	} else {
		return arr[(n+1)/2]
	}
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
