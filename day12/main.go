package main

import (
	"bufio"
	"container/list"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("./input")
	scanner := bufio.NewScanner(file)

	registerMap := make(map[string]int)
	registerMap["c"] = 1 //for second answer

	botID := ""

	var commands list.List

	for scanner.Scan() {
		line := scanner.Text()

		commands.PushBack(line)
	}

	for command := commands.Front(); command != nil; command = command.Next() {
		line := command.Value.(string)

		if strings.HasPrefix(line, "cpy") {
			split := strings.Split(line[4:], " ")
			source := split[0]
			targetBotID := split[1]

			if sourceInt, err := strconv.Atoi(source); err == nil {
				registerMap[targetBotID] = sourceInt
			} else {
				registerMap[targetBotID] = registerMap[source]
			}
		}

		if strings.HasPrefix(line, "inc") {
			botID = line[4:]
			registerMap[botID]++
		}

		if strings.HasPrefix(line, "dec") {
			botID = line[4:]
			registerMap[botID]--
		}

		if strings.HasPrefix(line, "jnz") {
			split := strings.Split(line[4:], " ")
			source := split[0]
			skipCount, _ := strconv.Atoi(split[1])

			sourceValue := 0

			if sourceInt, err := strconv.Atoi(source); err == nil {
				sourceValue = sourceInt
			} else {
				sourceValue = registerMap[source]
			}

			if sourceValue > 0 {
				if skipCount > 0 {
					for i := 0; i < skipCount-1; i++ {
						command = command.Next()
						if command == nil {
							break
						}
					}
				} else {
					for i := 0; i > skipCount-1; i-- {
						command = command.Prev()
						if command == nil {
							break
						}
					}
				}
			}
		}
	}

	println("Answer: " + strconv.Itoa(registerMap["a"]))
}
