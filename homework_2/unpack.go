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
	var lastRune rune = -1

	for i := 0; i < len(s); i++ {
		if unicode.IsDigit(rune(s[i])) {
			counter, _ := strconv.ParseInt(string(s[i]), 10, 32)
			if counter > 1 {
				temp := strings.Repeat(string(lastRune), int(counter-1))
				unpackString.WriteString(temp)
			}
		} else {
			unpackString.WriteRune(rune(s[i]))
			lastRune = rune(s[i])
		}
	}

	return unpackString.String()
}
