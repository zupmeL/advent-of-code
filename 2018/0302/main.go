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
	nonOverlappingClaims []int
	points               map[point]int
}

func (f *fabric) add(id, x, y, width, height int) {
	if f.points == nil {
		f.points = make(map[point]int)
	}
	overlap := false
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			if _, ok := f.points[point{x + i, y + j}]; ok {
				f.removeOverlappingClaim(f.points[point{x + i, y + j}])
				overlap = true
			}
			f.points[point{x + i, y + j}] = id
		}
	}
	if !overlap {
		f.nonOverlappingClaims = append(f.nonOverlappingClaims, id)
	}
}

func (f *fabric) removeOverlappingClaim(id int) {
	for k, v := range f.nonOverlappingClaims {
		if id == v {
			f.nonOverlappingClaims[k] = f.nonOverlappingClaims[len(f.nonOverlappingClaims)-1]
			f.nonOverlappingClaims = f.nonOverlappingClaims[:len(f.nonOverlappingClaims)-1]
		}
	}
	return
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
		fabric.add(id, x, y, w, h)
	}

	fmt.Println("The claims that don't overlap in the fabric are: ", fabric.nonOverlappingClaims)

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
