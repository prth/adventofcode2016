package main

import (
	"bufio"
	"os"
	"strings"
)

func main() {
	file, _ := os.Open("./input")
	scanner := bufio.NewScanner(file)

	charMap := make(map[int]map[rune]int)

	maxLength := 0

	for scanner.Scan() {
		line := scanner.Text()

		for index, char := range line {
			if charMap[index] == nil {
				maxLength++
				charMap[index] = make(map[rune]int)
			}

			charMap[index][char]++
		}
	}

	var ans1 = make([]string, maxLength)
	var ans2 = make([]string, maxLength)
	var maxChar, minChar rune
	var maxCount, minCount int

	for i, m := range charMap {
		maxCount = 0
		minCount = -1

		for char, count := range m {
			if maxCount < count {
				maxChar = char
				maxCount = count
			}

			if minCount == -1 || minCount > count {
				minChar = char
				minCount = count
			}
		}

		ans1[i] = string(maxChar)
		ans2[i] = string(minChar)
	}

	println("Answer #1: " + strings.Join(ans1[:], ""))
	println("Answer #2: " + strings.Join(ans2[:], ""))
}
