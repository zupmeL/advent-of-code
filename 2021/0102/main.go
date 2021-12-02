package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const SLICE_LENGTH = 3

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	arr := getAllItems(scanner)

	if len(arr) > SLICE_LENGTH {
		firstSum := getSliceSum(arr[0:SLICE_LENGTH])
		increments := getIncrements(arr[1:], firstSum)
		fmt.Println("INCREMENTS", increments)
	} else {
		log.Fatal("Not enough items")
	}

}

func getAllItems(scanner *bufio.Scanner) []int {
	items := make([]int, 0)

	for scanner.Scan() {
		item, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal("Couldn't convert scanned row to number")
		}
		items = append(items, item)
	}
	return items
}

func getSliceSum(slice []int) int {
	sum := 0
	for _, value := range slice {
		sum += value
	}
	return sum
}

func getIncrements(arr []int, previousSum int) int {
	if len(arr) >= SLICE_LENGTH {
		sum := getSliceSum(arr[0:SLICE_LENGTH])
		if sum > previousSum {
			return getIncrements(arr[1:], sum) + 1
		} else {
			return getIncrements(arr[1:], sum)
		}
	}
	return 0
}
