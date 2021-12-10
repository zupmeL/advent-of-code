package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	rows := readFile("input.txt")
	total := 0
	for _, row := range rows {
		parts := strings.Split(row, " | ")
		dictionary := newDictionary(strings.Split(parts[0], " "))
		number := 0
		for _, encode := range strings.Split(parts[1], " ") {
			decode := dictionary.convert(encode)
			number = number*10 + decode
		}
		total += number
	}

	fmt.Println("Total sum is:", total)

}

type dictionary struct {
	conversionChart [10][]int
}

func newDictionary(arr []string) dictionary {

	var conversionChart [10][]int
	remain := make([][]int, 0)
	for _, str := range arr {
		runes := sortRunes(str)
		switch len(str) {
		case 2:
			conversionChart[1] = runes
		case 3:
			conversionChart[7] = runes
		case 4:
			conversionChart[4] = runes
		case 7:
			conversionChart[8] = runes
		default:
			remain = append(remain, runes)
		}
	}
	twoOrFive := make([][]int, 0)
	for _, runes := range remain {
		switch len(runes) {
		case 5: //2-3-5
			if contains(runes, conversionChart[7]) {
				conversionChart[3] = runes
			} else {
				twoOrFive = append(twoOrFive, runes)
			}
		case 6: //0-6-9
			if contains(runes, conversionChart[4]) {
				conversionChart[9] = runes
			} else {
				if contains(runes, conversionChart[7]) {
					conversionChart[0] = runes
				} else {
					conversionChart[6] = runes
				}
			}
		}
	}

	for _, runes := range twoOrFive {
		if contains(conversionChart[6], runes) {
			conversionChart[5] = runes
		} else {
			conversionChart[2] = runes
		}
	}

	dictionary := dictionary{conversionChart}

	return dictionary
}

func (dictionary *dictionary) convert(str string) int {
	test := sortRunes(str)
	for index, conversion := range dictionary.conversionChart {
		if len(conversion) == len(test) {
			equal := true
			for i, v := range conversion {
				if v != test[i] {
					equal = false
					break
				}
			}
			if equal {
				return index
			}
		}
	}
	log.Fatal("Couldn't convert")
	return -1
}

func sortRunes(str string) []int {
	arr := make([]int, 0)
	for _, char := range str {
		arr = append(arr, int(char))
	}
	sort.Ints(arr)
	return arr
}

func contains(word, search []int) bool {
	isPresent := true
	for _, x := range search {
		i := sort.SearchInts(word, x)
		if !(i < len(word) && word[i] == x) {
			isPresent = false
			break
		}
	}
	return isPresent
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
