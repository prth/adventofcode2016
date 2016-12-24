package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var ans1 string
var ans2 int

func main() {
	file, _ := os.Open("./input")
	scanner := bufio.NewScanner(file)

	botMap := make(map[string]*Bot)

	var rootBot string
	ans2 = 1

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "bot ") {
			botID := line[4 : strings.Index(line, "gives")-1]

			if botMap[botID] == nil {
				botMap[botID] = &Bot{id: botID}
			}

			gate1 := strings.Index(line, "low to")
			gate2 := strings.Index(line, "high to")

			if strings.HasPrefix(line[gate1+7:], "bot") {
				botMap[botID].low = line[gate1+11 : strings.Index(line, "and")-1]
			}

			if strings.HasPrefix(line[gate1+7:], "output") {
				botMap[botID].outputLow = line[gate1+14 : strings.Index(line, "and")-1]
			}

			if strings.HasPrefix(line[gate2+8:], "bot") {
				botMap[botID].high = line[gate2+12:]
			}

			if strings.HasPrefix(line[gate2+8:], "output") {
				botMap[botID].outputHigh = line[gate2+15:]
			}
		} else {
			value, _ := strconv.Atoi(line[6 : strings.Index(line, "goes")-1])
			botID := line[strings.LastIndex(line, " ")+1:]

			if botMap[botID] == nil {
				botMap[botID] = &Bot{id: botID}
			}

			if botMap[botID].value1 == 0 {
				botMap[botID].value1 = value
			} else {
				botMap[botID].value2 = value
				rootBot = botID
			}
		}
	}

	dfs(*botMap[rootBot], botMap)

	println("Answer #1: " + ans1)
	println("Answer #2: " + strconv.Itoa(ans2))
}

func dfs(bot Bot, botMap map[string]*Bot) {
	highValue, lowValue := bot.value1, bot.value2

	if bot.value2 > bot.value1 {
		highValue, lowValue = bot.value2, bot.value1
	}

	bot.value1 = 0
	bot.value2 = 0

	if highValue == 61 && lowValue == 17 {
		ans1 = bot.id
	}

	if bot.high != "" && botMap[bot.high].value2 == 0 {
		if botMap[bot.high].value1 == 0 {
			botMap[bot.high].value1 = highValue
		} else {
			botMap[bot.high].value2 = highValue
		}

		if botMap[bot.high].value2 > 0 {
			dfs(*botMap[bot.high], botMap)
		}
	} else if bot.outputHigh != "" {
		if bot.outputHigh == "0" || bot.outputHigh == "1" || bot.outputHigh == "2" {
			ans2 = ans2 * highValue
		}
	}

	if bot.low != "" && botMap[bot.low].value2 == 0 {
		if botMap[bot.low].value1 == 0 {
			botMap[bot.low].value1 = lowValue
		} else {
			botMap[bot.low].value2 = lowValue
		}

		if botMap[bot.low].value2 > 0 {
			dfs(*botMap[bot.low], botMap)
		}
	} else if bot.outputLow != "" {
		if bot.outputLow == "0" || bot.outputLow == "1" || bot.outputLow == "2" {
			ans2 = ans2 * lowValue
		}
	}
}

type Bot struct {
	id         string
	high       string
	low        string
	value1     int
	value2     int
	outputHigh string
	outputLow  string
}
