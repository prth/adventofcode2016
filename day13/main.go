package main

import "strconv"
import "strings"

var officeMap map[Coordinate]string
var visitedMap map[Coordinate]bool

var input = 1364

var start = Coordinate{
	x: 1,
	y: 1,
}

var destination = Coordinate{
	x: 31,
	y: 39,
}

func main() {
	officeMap = make(map[Coordinate]string)
	visitedMap = make(map[Coordinate]bool)

	path1 := getFewestNoOfSteps(start, destination)

	println("Answer #1: " + strconv.Itoa(len(path1)-1))
}

func getFewestNoOfSteps(start Coordinate, end Coordinate) []Coordinate {
	queue := []([]Coordinate){}
	queue = append(queue, []Coordinate{start})

	for len(queue) > 0 {
		path := queue[0]

		queue = queue[1:]

		if path[len(path)-1] == end {
			return path
		}

		newPaths := getPossibleNewPaths(path)

		for _, p := range newPaths {
			queue = append(queue, p)
		}
	}

	return []Coordinate{}
}

func getPossibleNewPaths(currentPath []Coordinate) []([]Coordinate) {
	node := currentPath[len(currentPath)-1]
	neighbours := getNeighbors(node)

	newPaths := []([]Coordinate){}

	for _, c := range neighbours {
		if !visitedMap[c] {
			newPath := []Coordinate{}

			for _, node := range currentPath {
				newPath = append(newPath, node)
			}

			newPath = append(newPath, c)

			newPaths = append(newPaths, newPath)

			visitedMap[c] = true
		}
	}

	return newPaths
}

func getNeighbors(coordinate Coordinate) []Coordinate {
	neighbors := []Coordinate{}

	if coordinate.x > 0 {
		l1 := Coordinate{
			x: coordinate.x - 1,
			y: coordinate.y,
		}

		if isOpenSpace(l1) {
			neighbors = append(neighbors, l1)
		}
	}

	if coordinate.y > 0 {
		d1 := Coordinate{
			x: coordinate.x,
			y: coordinate.y - 1,
		}

		if isOpenSpace(d1) {
			neighbors = append(neighbors, d1)
		}
	}

	if coordinate.x < destination.x+2 {
		r1 := Coordinate{
			x: coordinate.x + 1,
			y: coordinate.y,
		}

		if isOpenSpace(r1) {
			neighbors = append(neighbors, r1)
		}
	}

	if coordinate.y < destination.y+2 {
		u1 := Coordinate{
			x: coordinate.x,
			y: coordinate.y + 1,
		}

		if isOpenSpace(u1) {
			neighbors = append(neighbors, u1)
		}
	}

	return neighbors
}

func isOpenSpace(coordinate Coordinate) bool {
	if officeMap[coordinate] != "" {
		return officeMap[coordinate] == "."
	}

	input := input
	x := coordinate.x
	y := coordinate.y

	t := x*x + 3*x + 2*x*y + y + y*y + input
	b := strconv.FormatInt(int64(t), 2)

	isOpen := (strings.Count(b, "1") % 2) == 0

	if isOpen {
		officeMap[coordinate] = "."
	} else {
		officeMap[coordinate] = "#"
	}

	return isOpen
}

type Coordinate struct {
	x int
	y int
}
