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

	screen := NewScreen()

	for scanner.Scan() {
		line := scanner.Text()

		if strings.Index(line, "rect ") != -1 {
			gate1 := strings.Index(line, " ")
			gate2 := strings.Index(line, "x")

			x, _ := strconv.Atoi(line[gate1+1 : gate2])
			y, _ := strconv.Atoi(line[gate2+1:])

			screen.TurnOnPixels(x, y)
		}

		if strings.Index(line, "rotate row y") != -1 {
			gate1 := strings.Index(line, "=")
			gate2 := gate1 + strings.Index(line[gate1:], " ")
			gate3 := gate1 + strings.LastIndex(line[gate1:], " ")

			rowID, _ := strconv.Atoi(line[gate1+1 : gate2])
			count, _ := strconv.Atoi(line[gate3+1:])

			screen.RotateRow(rowID, count)
		}

		if strings.Index(line, "rotate column x") != -1 {
			gate1 := strings.Index(line, "=")
			gate2 := gate1 + strings.Index(line[gate1:], " ")
			gate3 := gate1 + strings.LastIndex(line[gate1:], " ")

			columnID, _ := strconv.Atoi(line[gate1+1 : gate2])
			count, _ := strconv.Atoi(line[gate3+1:])

			screen.RotateColumn(columnID, count)
		}
	}

	println("Pixels should be lit: " + strconv.Itoa(screen.GetLitPixelsCount())) //116

	println("Pixels #2")

	screen.PrintPixels() // UPOJFLBCEZ
}
