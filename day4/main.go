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

	sumPart1 := 0

	for scanner.Scan() {
		str := scanner.Text()

		room := parseRoom(str)

		if isValidRoom(room) {
			sumPart1 += room.sectorID

			println(decryptRoomName(room) + "  " + strconv.Itoa(room.sectorID))
		}
	}

	println("Sum of the sector IDs of the real rooms #1: " + strconv.Itoa(sumPart1)) // 361724

	//answer #2: 482
}

// Room containts encryptedName and checksum
type Room struct {
	encryptedName string
	sectorID      int
	checksum      string
}

func parseRoom(str string) Room {
	lastDashIndex := strings.LastIndex(str, "-")
	firstOpenBracketIndex := strings.Index(str, "[")

	encryptedName := str[0:lastDashIndex]

	checksum := str[firstOpenBracketIndex+1 : len(str)-1]
	sectorID, _ := strconv.Atoi(str[lastDashIndex+1 : firstOpenBracketIndex])

	return Room{
		encryptedName: encryptedName,
		sectorID:      sectorID,
		checksum:      checksum,
	}
}

func isValidRoom(room Room) bool {
	charCountMap := make(map[string]int)

	for _, r := range room.encryptedName {
		if string(r) == "-" {
			continue
		}

		charCountMap[string(r)]++
	}

	for i := 0; i < len(room.checksum)-1; i++ {
		c1 := string(room.checksum[i])
		c2 := string(room.checksum[i+1])

		if _, ok := charCountMap[c1]; !ok {
			return false
		}

		if _, ok := charCountMap[c2]; !ok {
			return false
		}

		if charCountMap[c1] < charCountMap[c2] {
			return false
		}

		if charCountMap[c1] == charCountMap[c2] && c1 > c2 {
			return false
		}
	}

	return true
}

const alphabets = "abcdefghijklmnopqrstuvwxyz"

func decryptRoomName(room Room) string {
	shift := room.sectorID % len(alphabets)

	decryptedName := ""

	for _, r := range room.encryptedName {
		if string(r) == "-" {
			decryptedName += " "
			continue
		}

		encryptedIndex := strings.Index(alphabets, string(r))

		decryptedIndex := encryptedIndex + shift

		if decryptedIndex >= len(alphabets) {
			decryptedIndex = decryptedIndex - len(alphabets)
		}

		decryptedName += string(alphabets[decryptedIndex])
	}

	return decryptedName
}
