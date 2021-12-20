package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	input := readFile("input.txt")

	binary := ""
	for _, c := range string(input) {
		ui, _ := strconv.ParseUint(string(c), 16, 64)
		b := fmt.Sprintf("%04b", ui)
		binary += b
	}

	_, result := parseBits(binary, 0)
	fmt.Println("Result: ", result)
}

func parseBits(binary string, i int) (int, int) {
	typeId := convert(binary[i+3 : i+6])

	i += 6

	switch typeId {
	case 4:
		value := 0
		for {
			num := convert(binary[i+1 : i+5])
			value = value*16 + int(num)
			if binary[i] == '0' {
				return i + 5, value
			}
			i += 5
		}
	default:
		values := make([]int, 0)
		typeLength := binary[i]
		i++
		if typeLength == '0' {
			lenBits := convert(binary[i : i+15])
			i += 15
			startIndex := i
			for {
				nextIndex, value := parseBits(binary, i)
				values = append(values, value)
				i = nextIndex
				if nextIndex-startIndex == int(lenBits) {
					break
				}
			}
		} else {
			numberOfPackets := convert(binary[i : i+11])
			i += 11
			for n := 0; n < int(numberOfPackets); n++ {
				nextIndex, value := parseBits(binary, i)
				values = append(values, value)
				i = nextIndex
			}
		}
		value := performOp(typeId, values)
		return i, value
	}
}

func sum(values []int) int {
	res := 0
	for _, v := range values {
		res += v
	}
	return res
}

func mul(values []int) int {
	res := 1
	for _, v := range values {
		res *= v
	}
	return res
}

func min(values []int) int {
	res := math.MaxInt
	for _, v := range values {
		if v < res {
			res = v
		}
	}
	return res
}

func max(values []int) int {
	res := 0
	for _, v := range values {
		if v > res {
			res = v
		}
	}
	return res
}

func performOp(typeId int64, values []int) int {
	switch typeId {
	case 0:
		return sum(values)
	case 1:
		return mul(values)
	case 2:
		return min(values)
	case 3:
		return max(values)
	case 5:
		if values[0] > values[1] {
			return 1
		}
		return 0
	case 6:
		if values[0] < values[1] {
			return 1
		}
		return 0
	case 7:
		if values[0] == values[1] {
			return 1
		}
		return 0
	default:
		return 0
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
