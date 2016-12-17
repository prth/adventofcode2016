package main

import (
	"math"
	"strconv"
	"strings"
)

func main() {
	instructions := "R4, R3, L3, L2, L1, R1, L1, R2, R3, L5, L5, R4, L4, R2, R4, L3, R3, L3, R3, R4, R2, L1, R2, L3, L2, L1, R3, R5, L1, L4, R2, L4, R3, R1, R2, L5, R2, L189, R5, L5, R52, R3, L1, R4, R5, R1, R4, L1, L3, R2, L2, L3, R4, R3, L2, L5, R4, R5, L2, R2, L1, L3, R3, L4, R4, R5, L1, L1, R3, L5, L2, R76, R2, R2, L1, L3, R189, L3, L4, L1, L3, R5, R4, L1, R1, L1, L1, R2, L4, R2, L5, L5, L5, R2, L4, L5, R4, R4, R5, L5, R3, L1, L3, L1, L1, L3, L4, R5, L3, R5, R3, R3, L5, L5, R3, R4, L3, R3, R1, R3, R2, R2, L1, R1, L3, L3, L3, L1, R2, L1, R4, R4, L1, L1, R3, R3, R4, R1, L5, L2, R2, R3, R2, L3, R4, L5, R1, R4, R5, R4, L4, R1, L3, R1, R3, L2, L3, R1, L2, R3, L3, L1, L3, R4, L4, L5, R3, R5, R4, R1, L2, R3, R5, L5, L4, L1, L1"

	instructions = sanitizeString(instructions)

	coordinates := Coordinates{x: 0, y: 0}

	var firstLocationToVisitTwice Coordinates
	movementMapToFirstSameLocation := make(map[Coordinates]int)
	movementMapToFirstSameLocation[coordinates] = 1

	compass := NewCompass()

	for _, instruction := range strings.Fields(instructions) {
		distance := getDistance(instruction)
		direction := getDirection(instruction)

		compass.Move(direction)

		if firstLocationToVisitTwice != (Coordinates{}) {
			switch compass.GetFacingCardinalDirection() {
			case "N":
				coordinates.y += distance
			case "E":
				coordinates.x += distance
			case "S":
				coordinates.y -= distance
			case "W":
				coordinates.x -= distance
			}
		} else {
			// move in steps of 1 till the first location vistied twice is found
			for i := 0; i < distance; i++ {
				switch compass.GetFacingCardinalDirection() {
				case "N":
					coordinates.y++
				case "E":
					coordinates.x++
				case "S":
					coordinates.y--
				case "W":
					coordinates.x--
				}

				movementMapToFirstSameLocation[coordinates]++

				if movementMapToFirstSameLocation[coordinates] > 1 {
					firstLocationToVisitTwice = coordinates
				}
			}
		}
	}

	distanceFromStartingPoint := math.Abs(float64(coordinates.x)) + math.Abs(float64(coordinates.y))
	println("Distance: " + strconv.FormatFloat(distanceFromStartingPoint, 'f', 6, 64)) //288.000000

	distanceToFirstLocationVisitedTwice :=
		math.Abs(float64(firstLocationToVisitTwice.x)) + math.Abs(float64(firstLocationToVisitTwice.y))
	println("Distance To First Location Visited Twice: " +
		strconv.FormatFloat(distanceToFirstLocationVisitedTwice, 'f', 6, 64)) //111.000000
}

func sanitizeString(str string) string {
	replacer := strings.NewReplacer(",", "")

	return replacer.Replace(str)
}

func getDistance(instruction string) int {
	distance, _ := strconv.Atoi(instruction[1:])

	return distance
}

func getDirection(instruction string) string {
	return string(instruction[0])
}

// Coordinates x:y
type Coordinates struct {
	x int
	y int
}
