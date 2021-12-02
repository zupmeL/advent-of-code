package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type position struct {
	depth      int
	horizontal int
}

func (position *position) add(p position) {
	position.depth += p.depth
	position.horizontal += p.horizontal
}

type instruction struct {
	action string
	unit   int
}

type submarine struct {
	position position
	aim      int
}

func (submarine *submarine) execute(instruction instruction) {
	switch instruction.action {
	case "forward":
		movement := position{instruction.unit * submarine.aim, instruction.unit}
		submarine.position.add(movement)
	case "down":
		submarine.aim += instruction.unit
	case "up":
		submarine.aim -= instruction.unit
	}
}

func main() {
	input := readFile("input.txt")

	instructions := getInstructionsFromInput(input)

	submarine := submarine{position{0, 0}, 0}

	for _, instruction := range instructions {
		submarine.execute(instruction)
	}

	fmt.Println("The submarine is at position", submarine.position.depth, "depth and", submarine.position.horizontal, "horizontal")
	fmt.Println("Multiplication", submarine.position.depth*submarine.position.horizontal)
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

func getInstructionsFromInput(input []string) []instruction {
	instructions := make([]instruction, len(input))

	for _, row := range input {
		words := strings.Fields(row)
		item, err := strconv.Atoi(words[1])
		if err != nil {
			log.Fatal("Couldn't convert scanned unit to number")
		}
		instructions = append(instructions, instruction{words[0], item})
	}

	return instructions
}
