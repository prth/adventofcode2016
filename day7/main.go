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

	countTCP := 0
	countSSL := 0

	for scanner.Scan() {
		line := scanner.Text()

		ip := parseIP(line)

		if isIPValidTCP(ip) {
			countTCP++
		}

		if isIPValidSSL(ip) {
			countSSL++
		}
	}

	println("Valid TCP IPs: " + strconv.Itoa(countTCP))
	println("Valid SSL IPs: " + strconv.Itoa(countSSL))
}

// IP supernet/hypernet sequences
type IP struct {
	hypernetSequences list.List
	supernetSequences list.List
}

func parseIP(str string) IP {
	ip := IP{}
	ipStr := str

	for len(ipStr) > 0 {
		gate1 := strings.Index(ipStr, "[")

		if gate1 == -1 {
			ip.supernetSequences.PushBack(ipStr)
			break
		}

		gate2 := strings.Index(ipStr, "]")

		ip.supernetSequences.PushBack(ipStr[0:gate1])
		ip.hypernetSequences.PushBack(ipStr[gate1+1 : gate2])

		ipStr = ipStr[gate2+1:]
	}

	return ip
}

func isIPValidTCP(ip IP) bool {
	isValid := false

	for e := ip.supernetSequences.Front(); e != nil; e = e.Next() {
		if hasABBA(e.Value.(string)) {
			isValid = true
			break
		}
	}

	if !isValid {
		return false
	}

	for e := ip.hypernetSequences.Front(); e != nil; e = e.Next() {
		if hasABBA(e.Value.(string)) {
			isValid = false
			break
		}
	}

	return isValid
}

func isIPValidSSL(ip IP) bool {
	for e := ip.supernetSequences.Front(); e != nil; e = e.Next() {
		str := e.Value.(string)

		for i := 0; i < len(str)-2; i++ {

			if !isABA(str[i : i+3]) {
				continue
			}

			expectedHyperStr := string(str[i+1]) + string(str[i]) + string(str[i+1])

			for h := ip.hypernetSequences.Front(); h != nil; h = h.Next() {
				if strings.Index(h.Value.(string), expectedHyperStr) > -1 {
					return true
				}
			}
		}
	}

	return false
}

func hasABBA(str string) bool {
	if len(str) < 4 {
		return false
	}

	for i := 0; i < len(str)-3; i++ {
		if isABBA(str[i : i+4]) {
			return true
		}
	}

	return false
}

func isABBA(str string) bool {
	if len(str) != 4 {
		return false
	}

	if str[0] == str[1] {
		return false
	}

	if str[0:2] == string(str[3])+string(str[2]) {
		return true
	}

	return false
}

func isABA(str string) bool {
	if len(str) != 3 {
		return false
	}

	if str[0] == str[1] {
		return false
	}

	if str[0] != str[2] {
		return false
	}

	return true
}
