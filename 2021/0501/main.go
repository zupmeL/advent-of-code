package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, maxValue := readFile("input.txt")

	diagram := newDiagram(maxValue + 1)

	for _, segment := range input {
		diagram.addSegment(segment)
	}
	overlaps := diagram.getOverlaps()
	fmt.Println("The overlaps are:", overlaps)
}

type point struct {
	x, y int
}

func PointFromString(str string) point {

	axes := strings.Split(strings.Trim(str, " "), ",")
	if len(axes) != 2 {
		log.Fatal("Error creating point")
	}
	coordinates := make([]int, 0)
	for _, value := range axes {
		axe, err := strconv.Atoi(value)
		if err != nil {
			log.Fatal(err)
		}
		coordinates = append(coordinates, axe)
	}
	p := point{coordinates[0], coordinates[1]}
	return p

}

type segment struct {
	pointA, pointB point
}

func (segment *segment) isVertical() bool {
	return segment.pointA.x == segment.pointB.x
}

func (segment *segment) isHorizontal() bool {
	return segment.pointA.y == segment.pointB.y
}

func (segment *segment) maxValue() int {
	maxX := 0
	if segment.pointA.x > segment.pointB.x {
		maxX = segment.pointA.x
	} else {
		maxX = segment.pointB.x
	}
	maxY := 0
	if segment.pointA.y > segment.pointB.y {
		maxY = segment.pointA.y
	} else {
		maxY = segment.pointB.y
	}
	if maxX > maxY {
		return maxX
	}
	return maxY
}

func SegmentFromString(str string) segment {
	points := strings.Split(str, " -> ")
	return segment{PointFromString(points[0]), PointFromString(points[1])}
}

type diagram struct {
	matrix [][]int
}

func newDiagram(len int) diagram {
	diagram := diagram{make([][]int, len)}
	for i := 0; i < len; i++ {
		diagram.matrix[i] = make([]int, len)
	}
	return diagram
}

func (diagram *diagram) addSegment(segment segment) {
	if segment.isHorizontal() {
		if segment.pointA.x < segment.pointB.x {
			for i := segment.pointA.x; i <= segment.pointB.x; i++ {
				diagram.matrix[segment.pointA.y][i]++
			}
		} else {
			for i := segment.pointB.x; i <= segment.pointA.x; i++ {
				diagram.matrix[segment.pointA.y][i]++
			}
		}
	}
	if segment.isVertical() {
		if segment.pointA.y < segment.pointB.y {
			for i := segment.pointA.y; i <= segment.pointB.y; i++ {
				diagram.matrix[i][segment.pointA.x]++
			}
		} else {
			for i := segment.pointB.y; i <= segment.pointA.y; i++ {
				diagram.matrix[i][segment.pointA.x]++
			}
		}
	}
}

func (diagram *diagram) getOverlaps() int {
	overlaps := 0
	for i := 0; i < len(diagram.matrix); i++ {
		for j := 0; j < len(diagram.matrix); j++ {
			if diagram.matrix[i][j] > 1 {
				overlaps++
			}
		}
	}
	return overlaps
}

func readFile(input string) ([]segment, int) {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	segments := make([]segment, 0)
	maxValue := 0

	for scanner.Scan() {
		segment := SegmentFromString(scanner.Text())
		if segment.isHorizontal() || segment.isVertical() {
			segments = append(segments, segment)
			currMaxValue := segment.maxValue()
			if currMaxValue > maxValue {
				maxValue = currMaxValue
			}
		}
	}

	return segments, maxValue
}
