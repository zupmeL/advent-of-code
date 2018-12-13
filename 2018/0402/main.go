package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type action struct {
	guardID    int
	actionType actionType
	time       time.Time
}

type actionType byte

const (
	actionShift actionType = iota
	actionSleep
	actionAwake
)

func main() {
	input := readFile("input.txt")
	sort.Strings(input)

	var actions []action

	for _, actionText := range input {
		divider := strings.Index(actionText, "]")
		timeText := actionText[1:divider]
		time, err := time.Parse("2006-01-02 15:04", timeText)
		if err != nil {
			log.Fatal(err)
		}

		action := action{time: time}

		splits := strings.Split(actionText[divider+2:], " ")
		switch splits[0] {
		case "Guard":
			guardID, err := strconv.Atoi(splits[1][1:])
			if err != nil {
				log.Fatal(err)
			}
			action.guardID = guardID
			action.actionType = actionShift
		case "falls":
			action.guardID = actions[len(actions)-1].guardID
			action.actionType = actionSleep
		case "wakes":
			action.guardID = actions[len(actions)-1].guardID
			action.actionType = actionAwake
		}
		actions = append(actions, action)
	}
	guardID, minute := findOpening(actions)
	fmt.Println(guardID * minute)
}

func findOpening(actions []action) (int, int) {
	var sleepSlots = map[int]map[int]int{}
	for i := 0; i < 60; i++ {
		sleepSlots[i] = make(map[int]int)
	}
	for index := 0; index < len(actions); index++ {
		if actions[index].actionType == actionAwake {

			for min := actions[index-1].time.Minute(); min < actions[index].time.Minute(); min++ {
				sleepSlots[min][actions[index].guardID]++
			}
		}
	}

	bestGuard := 0
	bestMinute := 0
	maxCounter := 0
	for minute, guards := range sleepSlots {
		for guardID, counter := range guards {
			if counter > maxCounter {
				bestGuard = guardID
				bestMinute = minute
				maxCounter = counter
			}
		}
	}
	return bestGuard, bestMinute
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
