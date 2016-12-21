package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("./input")

	scanner := bufio.NewScanner(file)

	countTrianglesPart1 := 0
	countTrianglesPart2 := 0

	// maintain buffer of triangles in set of 3
	// to evaluate valid triangles for stage 2
	var triangesStage2Buffer [3]Triange
	indexStage2Buffer := 0

	for scanner.Scan() {
		line := scanner.Text()
		line = sanitizeStr(line)

		var sides [3]int

		for i, side := range strings.Fields(line) {
			sides[i], _ = strconv.Atoi(side)
		}

		triange := Triange{sides: sides}

		if isValidTriange(triange) {
			countTrianglesPart1++
		}

		// prepare the buffer
		for i, side := range sides {
			triangesStage2Buffer[i].sides[indexStage2Buffer] = side
		}

		indexStage2Buffer++

		// evaluate the triangles once buffer is complete
		if indexStage2Buffer == 3 {
			for _, tri := range triangesStage2Buffer {
				if isValidTriange(tri) {
					countTrianglesPart2++
				}
			}

			// reset the stage2 buffer index
			indexStage2Buffer = 0
		}
	}

	println("Valid Triangles #1: " + strconv.Itoa(countTrianglesPart1))
	println("Valid Triangles #2: " + strconv.Itoa(countTrianglesPart2))
}

// Triange three sides
type Triange struct {
	sides [3]int
}

func sanitizeStr(str string) string {
	reg, _ := regexp.Compile("[ ]+")

	str = reg.ReplaceAllString(str, " ")

	return strings.TrimSpace(str)
}

func isValidTriange(triange Triange) bool {
	if triange.sides[0]+triange.sides[1] <= triange.sides[2] {
		return false
	}

	if triange.sides[0]+triange.sides[2] <= triange.sides[1] {
		return false
	}

	if triange.sides[1]+triange.sides[2] <= triange.sides[0] {
		return false
	}

	return true
}
