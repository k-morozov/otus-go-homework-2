package homework_2

import (
	"strconv"
	"strings"
	"unicode"
)

func Unpack(s string) string {
	var unpackString strings.Builder

	if 0 == len(s) {
		return unpackString.String()
	}
	var lastRune rune
	var startCounter int

	for index, value := range s {
		if unicode.IsDigit(rune(value)) {
			if 0 == startCounter {
				startCounter = index
			}
		} else {
			if 0 != startCounter {
				counter, _ := strconv.ParseInt(string(s[startCounter:index]), 10, 32)
				temp := strings.Repeat(string(lastRune), int(counter-1))
				unpackString.WriteString(temp)
				startCounter = 0
			}
			unpackString.WriteRune(rune(value))
			lastRune = rune(value)
		}
	}
	if 0 != startCounter {
		counter, _ := strconv.ParseInt(string(s[startCounter:]), 10, 32)
		temp := strings.Repeat(string(lastRune), int(counter-1))
		unpackString.WriteString(temp)
	}

	return unpackString.String()
}
