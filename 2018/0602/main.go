package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const maxDistance = 10000

type point struct {
	x, y int
}

func (a point) distance(b point) int {
	return abs(a.x-b.x) + abs(a.y-b.y)
}

func main() {
	points, edge := readFile("input.txt")

	total := 0
	for x := 0; x < edge; x++ {
		for y := 0; y < edge; y++ {
			under := checkDistance(point{x, y}, points)
			if under {
				total++
			}
		}
	}

	fmt.Println("The region size is: ", total)
}

func checkDistance(center point, points []point) bool {
	currDistance := 0
	for _, currPoint := range points {
		currDistance += center.distance(currPoint)
		if currDistance < maxDistance {
			continue
		}
		return false
	}
	return true
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
