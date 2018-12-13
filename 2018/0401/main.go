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
	var sleepDuration = map[int]int{}
	var sleepSlots = map[int]map[int]int{}
	for index := 0; index < len(actions); index++ {
		if actions[index].actionType == actionAwake {
			if sleepSlots[actions[index].guardID] == nil {
				sleepSlots[actions[index].guardID] = make(map[int]int)
			}
			for min := actions[index-1].time.Minute(); min < actions[index].time.Minute(); min++ {
				sleepSlots[actions[index].guardID][min]++
				sleepDuration[actions[index].guardID]++
			}
		}
	}

	var sleepyGuard int
	maxSleepyMinutes := 0
	for guardID, duration := range sleepDuration {
		if duration > maxSleepyMinutes {
			maxSleepyMinutes = duration
			sleepyGuard = guardID
		}
	}

	bestMinute := 0
	maxCounter := 0
	for minute, counter := range sleepSlots[sleepyGuard] {
		if counter > maxCounter {
			bestMinute = minute
			maxCounter = counter
		}
	}
	return sleepyGuard, bestMinute
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
