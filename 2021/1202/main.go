package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	rows := readFile("input.txt")

	paths := initPaths(rows)
	visited := make([]string, 0)
	fmt.Println("Total paths to end:", navigate(paths, "start", visited, true))
}

func initPaths(rows []string) map[string][]string {
	paths := make(map[string][]string)

	for _, row := range rows {
		words := strings.Split(row, "-")
		paths[words[0]] = append(paths[words[0]], words[1])
		paths[words[1]] = append(paths[words[1]], words[0])
	}
	return paths
}

func navigate(paths map[string][]string, node string, visited []string, bonus bool) int {
	if node == "end" {
		return 1
	}
	total := 0
	for _, path := range paths[node] {
		if path == "start" {
			continue
		}
		if contains(visited, path) {
			if bonus {
				total += navigate(paths, path, visited, false)
			} else {
				continue
			}
		} else {
			if node == strings.ToLower(node) {
				visited = append(visited, node)
			}
			total += navigate(paths, path, visited, bonus)
		}
	}
	return total
}

func contains(arr []string, str string) bool {
	for _, el := range arr {
		if el == str {
			return true
		}
	}
	return false
}

func readFile(input string) []string {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	rows := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rows = append(rows, scanner.Text())
	}

	return rows
}
