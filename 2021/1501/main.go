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
	fmt.Println("Shortest distance ", best.Distance, " following path ", best.Path)
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
		row := scanner.Text()
		for _, char := range row {
			arr = append(arr, int(char-'0'))
		}
		rows++
	}

	return arr, rows
}
