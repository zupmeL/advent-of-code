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
	points, folds := readFile("input.txt")

	if folds[0].direction == "x" {
		points = foldX(points, folds[0].value)
	} else {
		points = foldY(points, folds[0].value)
	}

	fmt.Println("Visible dots are:", len(points.points))
}

type point struct {
	x, y int
}

type matrix struct {
	points     []point
	maxX, maxY int
}

func (matrix *matrix) add(newPoint point) {
	if len(matrix.points) > 0 {
		found := false
		for _, p := range matrix.points {
			if p.x == newPoint.x && p.y == newPoint.y {
				found = true
				break
			}
		}
		if !found {
			matrix.points = append(matrix.points, newPoint)
		}
	} else {
		matrix.points = append(matrix.points, newPoint)
	}
}

type fold struct {
	direction string
	value     int
}

func foldX(m matrix, num int) matrix {
	var matrix matrix
	matrix.maxX = num - 1
	matrix.maxY = m.maxY
	for _, point := range m.points {
		if point.x >= num {
			point.x = num - (point.x - num)
		}
		matrix.add(point)
	}
	return matrix
}

func foldY(m matrix, num int) matrix {
	var matrix matrix
	matrix.maxX = m.maxX
	matrix.maxY = num - 1
	for _, point := range m.points {
		if point.y >= num {
			point.y = num - (point.y - num)
		}
		matrix.add(point)
	}
	return matrix
}

func readFile(input string) (matrix, []fold) {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	m := matrix{}
	folds := make([]fold, 0)

	scanner := bufio.NewScanner(file)
	f := 0

	for scanner.Scan() {
		row := scanner.Text()

		digits := strings.Split(row, ",")
		words := strings.Split(row, " ")
		if len(digits) == 2 {
			x, _ := strconv.Atoi(digits[0])
			y, _ := strconv.Atoi(digits[1])
			m.points = append(m.points, point{x, y})
			if m.maxX < x {
				m.maxX = x + 1
			}
			if m.maxY < y {
				m.maxY = y + 1
			}
		}
		if len(words) == 3 {
			ff := strings.Split(words[2], "=")
			value, _ := strconv.Atoi(ff[1])
			folds = append(folds, fold{ff[0], value})
			f++
		}
	}

	return m, folds
}
