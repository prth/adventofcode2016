package main

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"strings"
)

func main() {
	input := "ugkcyxxp"

	ans1 := ""
	ans2 := [8]string{}

	ans2Counter := 0

	index := 0
	hash := ""

	var positionIndex int
	var err error

	for len(ans1) < 8 || ans2Counter < 8 {
		index++
		hash = getMd5(input + strconv.Itoa(index))

		for !strings.HasPrefix(hash, "00000") {
			index++
			hash = getMd5(input + strconv.Itoa(index))
		}

		if len(ans1) < 8 {
			ans1 += string(hash[5])
		}

		positionIndex, err = strconv.Atoi(string(hash[5]))

		if err == nil && positionIndex < 8 && ans2[positionIndex] == "" {
			ans2[positionIndex] = string(hash[6])
			ans2Counter++
		}
	}

	println("Answer #1: " + ans1)
	println("Answer #2: " + strings.Join(ans2[:], ""))
}

func getMd5(str string) string {
	hasher := md5.New()
	hasher.Write([]byte(str))

	return hex.EncodeToString(hasher.Sum(nil))
}
