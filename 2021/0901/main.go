package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	input, rowLen := readFile("input.txt")

	lowPoints := make([]int, 0)

	for index, element := range input {
		if isLowPoint(input, index, rowLen) {
			lowPoints = append(lowPoints, element+1)
		}
	}

	fmt.Println("Sum of low points:", sum(lowPoints))
}

func sum(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	return arr[0] + sum(arr[1:])
}

func isLowPoint(matrix []int, index, rowLen int) bool {
	indexes := [4]int{upIndex(matrix, index, rowLen), rightIndex(matrix, index, rowLen), leftIndex(matrix, index, rowLen), downIndex(matrix, index, rowLen)}
	for _, i := range indexes {
		if !(i < 0 || matrix[i] > matrix[index]) {
			return false
		}
	}
	return true
}

func upIndex(matrix []int, index, rowLen int) int {
	if index/rowLen == 0 {
		return -1
	}
	return ((index/rowLen)-1)*rowLen + index%rowLen
}

func downIndex(matrix []int, index, rowLen int) int {
	row := index / rowLen
	rows := len(matrix) / rowLen
	if row == rows-1 {
		return -1
	}
	return ((index/rowLen)+1)*rowLen + index%rowLen
}

func leftIndex(matrix []int, index, rowLen int) int {
	if index%rowLen == 0 {
		return -1
	}
	return index - 1
}

func rightIndex(matrix []int, index, rowLen int) int {
	if index%rowLen == rowLen-1 {
		return -1
	}
	return index + 1
}

func readFile(input string) ([]int, int) {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	arr := make([]int, 0)
	rows := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := scanner.Text()
		for _, char := range row {
			element := int(char) - int('0')
			arr = append(arr, element)
		}
		rows++
	}

	return arr, len(arr) / rows
}
