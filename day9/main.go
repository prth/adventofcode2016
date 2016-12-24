package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("./input")
	scanner := bufio.NewScanner(file)

	// read the first line
	scanner.Scan()

	// get the input
	input := scanner.Text()
	input = strings.Replace(input, " ", "", -1)
	input2 := input

	ans1 := 0

	for strings.Index(input, "(") != -1 {
		gate1 := strings.Index(input, "(")
		gate2 := strings.Index(input, ")")

		marker := parseMarker(input[gate1+1 : gate2])

		ans1 += gate1

		charsCount := marker.charsCount

		if len(input)-gate2 < marker.charsCount {
			charsCount = len(input) - gate2
		}

		ans1 += charsCount * marker.repeatCount

		input = input[gate2+charsCount+1:]
	}

	ans1 += len(input)

	println("Answer #1: " + strconv.Itoa(ans1))

	ans2 := 0

	//println(time.Now().String())
	for strings.Index(input2, "(") != -1 {
		gate1 := strings.Index(input2, "(")
		gate2 := strings.Index(input2, ")")

		marker := parseMarker(input2[gate1+1 : gate2])

		charsCount := marker.charsCount

		if len(input2)-gate2 < marker.charsCount {
			charsCount = len(input2) - gate2
		}

		ans2 += gate1

		input2 = strings.Repeat(input2[gate2+1:gate2+charsCount+1], marker.repeatCount) +
			input2[gate2+charsCount+1:]

		newGate := strings.Index(input2, "(")

		if newGate != -1 {
			ans2 += newGate
			input2 = input2[newGate:]
		}
	}

	ans2 += len(input2)
	//println(time.Now().String())
	println("Answer #2: " + strconv.Itoa(ans2))
}

type Marker struct {
	charsCount  int
	repeatCount int
}

func parseMarker(str string) Marker {
	m := strings.Split(str, "x")

	charsCount, _ := strconv.Atoi(m[0])
	repeatCount, _ := strconv.Atoi(m[1])

	return Marker{
		charsCount:  charsCount,
		repeatCount: repeatCount,
	}
}
