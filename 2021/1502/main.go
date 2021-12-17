package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/RyanCarrier/dijkstra"
)

func main() {
	input, rows := readFile("input.txt")

	rowlen := len(input) / rows

	graph := dijkstra.NewGraph()

	for index, cost := range input {
		graph.AddVertex(index)

		if index == 0 {
			continue
		}

		if index%rowlen != 0 {
			graph.AddArc(index-1, index, int64(cost))
			graph.AddArc(index, index-1, int64(input[index-1]))
		}
		if index >= rowlen {
			graph.AddArc(index-rowlen, index, int64(cost))
			graph.AddArc(index, index-rowlen, int64(input[index-rowlen]))
		}

	}

	best, err := graph.Shortest(0, len(input)-1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Shortest distance ", best.Distance)
}

func readFile(input string) ([]int, int) {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	arr := make([]int, 0)
	rows := 0
	for scanner.Scan() {
		str := scanner.Text()
		row := make([]int, 0)
		for _, char := range str {
			row = append(row, int(char-'0'))
		}

		row = plus(row)
		arr = append(arr, row...)
		rows++
	}

	arr = plus(arr)
	rows *= 5
	return arr, rows
}

func plus(arr []int) []int {
	delta := make([]int, 0)
	for i := 1; i < 5; i++ {
		for _, el := range arr {
			el = el + i
			if el > 9 {
				delta = append(delta, el%9)
			} else {
				delta = append(delta, el)
			}
		}
	}
	return append(arr, delta...)
}
