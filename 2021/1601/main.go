package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var sum int

func main() {
	input := readFile("input.txt")

	binary := ""
	for _, c := range string(input) {
		ui, _ := strconv.ParseUint(string(c), 16, 64)
		b := fmt.Sprintf("%04b", ui)
		binary += b
	}

	parseBits(binary, 0)

	fmt.Println("sum: ", sum)
}

func parseBits(binary string, i int) int {
	version := convert(binary[i : i+3])
	typeId := convert(binary[i+3 : i+6])

	sum += int(version)
	i += 6

	switch typeId {
	case 4:
		for {
			if binary[i] == '0' {
				return i + 5
			}
			i += 5
		}
	default:
		typeLength := binary[i]
		i++
		if typeLength == '0' {
			lenBits := convert(binary[i : i+15])
			i += 15
			startIndex := i
			for {
				nextIndex := parseBits(binary, i)
				i = nextIndex
				if nextIndex-startIndex == int(lenBits) {
					break
				}
			}
		} else {
			numberOfPackets := convert(binary[i : i+11])
			i += 11
			for n := 0; n < int(numberOfPackets); n++ {
				nextIndex := parseBits(binary, i)
				i = nextIndex
			}
		}
		return i
	}
}

func convert(str string) int64 {
	i, err := strconv.ParseInt(str, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	return i
}

func readFile(input string) string {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	hex := scanner.Text()

	return hex
}
