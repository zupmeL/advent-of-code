package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const maxInt = 9223372036854775807

type point struct {
	x, y int
}

type coverage struct {
	id, dist int
}

func (c1 coverage) compare(c2 coverage) coverage {
	if c1.dist > c2.dist {
		return c2
	}
	if c1.dist < c2.dist {
		return c1
	}
	return coverage{-1, c1.dist}
}

func main() {
	points, edge := readFile("input.txt")

	matrix := initMatrix(edge)

	for id, center := range points {
		matrix[center] = coverage{id, 0}
		for x := 0; x < edge; x++ {
			for y := 0; y < edge; y++ {
				currPoint := point{x, y}
				if currPoint == center {
					continue
				}
				currCoverage := coverage{id, getDistance(currPoint, center)}
				matrix[point{x, y}] = currCoverage.compare(matrix[point{x, y}])
			}
		}
	}

	largestArea := 0
	for id := range points {
		touchEdges := false
		currArea := 0
		for x := 0; x < edge; x++ {
			for y := 0; y < edge; y++ {
				if matrix[point{x, y}].id == id {
					if x == 0 || x == edge-1 || y == 0 || y == edge-1 {
						touchEdges = true
						break
					}
					currArea++
				}
			}
		}
		if !touchEdges && currArea > largestArea {
			largestArea = currArea
		}
	}

	fmt.Println("The largest area is: ", largestArea)
}

func initMatrix(edge int) map[point]coverage {
	matrix := make(map[point]coverage, edge)

	for x := 0; x < edge; x++ {
		for y := 0; y < edge; y++ {
			matrix[point{x, y}] = coverage{-1, maxInt}
		}
	}
	return matrix
}

func getDistance(a, b point) int {
	return abs(a.x-b.x) + abs(a.y-b.y)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func readFile(input string) ([]point, int) {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	max := 0
	var points []point
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var point point
		fmt.Sscanf(scanner.Text(), "%d, %d", &point.x, &point.y)
		points = append(points, point)
		if point.x > max {
			max = point.x
		}
		if point.y > max {
			max = point.y
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return points, max + 1
}
