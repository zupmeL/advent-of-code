package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type point struct {
	x, y int
}

type fabric struct {
	points map[point]int
}

func (f *fabric) add(x, y, width, height int) {
	if f.points == nil {
		f.points = make(map[point]int)
	}
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			f.points[point{x + i, y + j}]++
		}
	}
}

func (f *fabric) stackedPoints() int {
	stackedPoints := 0
	for _, counter := range f.points {
		if counter > 1 {
			stackedPoints++
		}
	}
	return stackedPoints
}

func main() {
	input := readFile("input.txt")

	var fabric fabric

	for _, claim := range input {
		var id, x, y, w, h int
		_, err := fmt.Sscanf(claim, "#%d @ %d,%d: %dx%d", &id, &x, &y, &w, &h)
		if err != nil {
			log.Fatal(err)
		}
		fabric.add(x, y, w, h)
	}

	fmt.Println("The number of stacked points in the fabric is: ", fabric.stackedPoints())

}

func readFile(input string) []string {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var content []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		content = append(content, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return content
}
