package unpack

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
		if unicode.IsDigit(value) {
			if 0 == startCounter {
				startCounter = index
			}
		} else if unicode.IsLetter(value) {
			if 0 != startCounter {
				temp := unpackIntervalImpl([]rune(s[startCounter:index]), lastRune)
				unpackString.WriteString(temp)
				startCounter = 0
			}
			_, _ = unpackString.WriteRune(value)
			lastRune = value
		}
	}
	if 0 != startCounter {
		temp := unpackIntervalImpl([]rune(s[startCounter:]), lastRune)
		unpackString.WriteString(temp)
	}

	return unpackString.String()
}

func unpackIntervalImpl(s []rune, lastRune rune) string {
	counter, _ := strconv.Atoi(string(s[:]))
	temp := strings.Repeat(string(lastRune), counter-1)
	return temp
}
