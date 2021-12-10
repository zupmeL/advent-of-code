package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	input, rowLen := readFile("input.txt")

	lowPoints := make([]int, 0)

	for index := range input {
		if isLowPoint(input, index, rowLen) {
			lowPoints = append(lowPoints, index)
		}
	}

	basins := make([]int, 0)
	for _, pointIndex := range lowPoints {
		basin := make([]int, 0)
		basins = append(basins, len(mapBasin(input, basin, pointIndex, rowLen)))
	}
	sort.Ints(basins)
	fmt.Println("Multiplication between 3 largest basins:", mul(basins[len(basins)-3:]))

}

func mul(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	if len(arr) == 1 {
		return arr[0]
	}
	return arr[0] * mul(arr[1:])
}

func mapBasin(matrix, basin []int, index, rowLen int) []int {
	if index >= 0 && matrix[index] != 9 && !exists(basin, index) {
		basin = append(basin, index)
		basin = mapBasin(matrix, basin, upIndex(matrix, index, rowLen), rowLen)
		basin = mapBasin(matrix, basin, downIndex(matrix, index, rowLen), rowLen)
		basin = mapBasin(matrix, basin, leftIndex(matrix, index, rowLen), rowLen)
		basin = mapBasin(matrix, basin, rightIndex(matrix, index, rowLen), rowLen)
	}
	return basin
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

func exists(arr []int, x int) bool {
	for _, el := range arr {
		if el == x {
			return true
		}
	}
	return false
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
