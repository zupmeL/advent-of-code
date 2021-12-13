package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const STEPS = 100

func main() {
	rows := readFile("input.txt")

	flashes := 0
	for i := 0; i < STEPS; i++ {
		flashes += step(rows)
	}

	fmt.Println("Flashes after 100 steps are:", flashes)
}

func step(matrix [][]int) int {
	flashes := 0
	for row := range matrix {
		for col := range matrix[row] {
			matrix[row][col]++
			if matrix[row][col] == 10 {
				matrix[row][col]++
				flashes += flash(matrix, row, col)
			}
		}
	}
	for row := range matrix {
		for col := range matrix[row] {
			if matrix[row][col] > 9 {
				matrix[row][col] = 0
			}
		}
	}
	return flashes
}

func flash(matrix [][]int, row, col int) int {
	flashes := 1
	for r := row - 1; r <= row+1; r++ {
		for c := col - 1; c <= col+1; c++ {
			if r == row && c == col {
				continue
			}
			if 0 <= r && r < len(matrix) && 0 <= c && c < len(matrix[r]) {
				matrix[r][c]++
				if matrix[r][c] == 10 {
					matrix[r][c]++
					flashes += flash(matrix, r, c)
				}
			}
		}
	}
	return flashes
}

func readFile(input string) [][]int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	rows := make([][]int, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := make([]int, 0)
		for _, char := range scanner.Text() {
			digit := int(char) - int('0')
			row = append(row, digit)
		}
		rows = append(rows, row)
	}

	return rows
}
