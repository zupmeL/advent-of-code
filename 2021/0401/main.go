package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const MATRIX_LEN = 5

type board struct {
	matrix [MATRIX_LEN * MATRIX_LEN]int
	marked [MATRIX_LEN * MATRIX_LEN]bool
	bingo  bool
}

func (b *board) add(item, row_index, col_index int) {
	b.matrix[MATRIX_LEN*row_index+col_index] = item
}

func (b *board) mark(call int) {
	for index, value := range b.matrix {
		if value == call {
			b.marked[index] = true
		}
	}
}

func (b *board) checkBingo() bool {
	b.bingo = false
	for i := 0; i < MATRIX_LEN && !b.bingo; i++ {
		for j := 0; j < MATRIX_LEN; j++ {
			if !b.marked[i*MATRIX_LEN+j] {
				break
			}
			if j == MATRIX_LEN-1 {
				b.bingo = true
			}
		}
	}
	for j := 0; j < MATRIX_LEN && !b.bingo; j++ {
		for i := 0; i < MATRIX_LEN; i++ {
			if !b.marked[i*MATRIX_LEN+j] {
				break
			}
			if i == MATRIX_LEN-1 {
				b.bingo = true
			}
		}
	}
	return b.bingo
}

func (b *board) sumUnmarked() int {
	acc := 0
	for i := 0; i < MATRIX_LEN; i++ {
		for j := 0; j < MATRIX_LEN; j++ {
			if !b.marked[i*MATRIX_LEN+j] {
				acc += b.matrix[i*MATRIX_LEN+j]
			}
		}
	}
	return acc
}

func main() {
	calls, boards := readFile("input.txt")

	var call int
	bingoIndex := -1

	for _, call = range calls {
		for index := 0; index < len(boards); index++ {
			boards[index].mark(call)
			if boards[index].checkBingo() {
				bingoIndex = index
				break
			}
		}
		if bingoIndex >= 0 {
			break
		}
	}
	fmt.Print("Multiplication between last call and the sum of unmarked cells is:\n")
	fmt.Println(call * boards[bingoIndex].sumUnmarked())

}

func readFile(input string) ([]int, []board) {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var calls []int
	var boards []board
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	row := scanner.Text()
	for _, item := range strings.Split(row, ",") {
		call, err := strconv.Atoi(item)
		if err != nil {
			log.Fatal(err)
		}
		calls = append(calls, call)
	}
	for scanner.Scan() {
		scanner.Text() // blank row
		boards = append(boards, readBoard(scanner, MATRIX_LEN))
	}

	return calls, boards
}

func readBoard(scanner *bufio.Scanner, length int) board {
	board := board{}
	for i := 0; i < length; i++ {
		scanner.Scan()
		row := scanner.Text()
		j := 0
		for _, item := range strings.Split(strings.Trim(strings.ReplaceAll(row, "  ", " "), " "), " ") {
			cell, err := strconv.Atoi(item)
			if err != nil {
				log.Fatal(err)
			}
			board.add(cell, i, j)
			j++
		}
	}
	return board
}
