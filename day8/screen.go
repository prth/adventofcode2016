package main

import "strconv"

// Screen - display tech
type Screen interface {
	RotateRow(rowID int, count int)
	RotateColumn(columnID int, count int)
	TurnOnPixels(width int, height int)
	GetLitPixelsCount() int
	PrintPixels()
}

type screen struct {
	pixels [6][50]int
}

// NewScreen - return instance of new screen
func NewScreen() Screen {
	return &screen{
		pixels: [6][50]int{},
	}
}

func (s *screen) RotateRow(rowID int, count int) {
	str := ""

	for i := 0; i < 50; i++ {
		str += strconv.Itoa(s.pixels[rowID][i])
	}

	str = rotateStr(str, count)

	for i := 0; i < 50; i++ {
		s.pixels[rowID][i], _ = strconv.Atoi(string(str[i]))
	}
}

func (s *screen) RotateColumn(columnID int, count int) {
	str := ""

	for j := 0; j < 6; j++ {
		str += strconv.Itoa(s.pixels[j][columnID])
	}

	str = rotateStr(str, count)

	for j := 0; j < 6; j++ {
		s.pixels[j][columnID], _ = strconv.Atoi(string(str[j]))
	}
}

func (s *screen) TurnOnPixels(width int, height int) {
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			s.pixels[j][i] = 1
		}
	}
}

func (s *screen) GetLitPixelsCount() int {
	count := 0

	for i := 0; i < 6; i++ {
		for j := 0; j < 50; j++ {
			if s.pixels[i][j] == 1 {
				count++
			}
		}
	}

	return count
}

func (s *screen) PrintPixels() {
	for i := 0; i < 6; i++ {
		for j := 0; j < 50; j++ {
			if s.pixels[i][j] == 1 {
				print("*")
			} else {
				print(" ")
			}
		}

		println("")
	}
}

func rotateStr(str string, shift int) string {
	shift = shift % len(str)
	rotatedStr := str[len(str)-shift:] + str[0:len(str)-shift]

	return rotatedStr
}
