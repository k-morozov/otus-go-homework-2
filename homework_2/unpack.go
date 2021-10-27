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
	var lastOp bool = true

	for i := 0; i < len(s); i++ {
		if unicode.IsDigit(rune(s[i])) {
			lastOp = true
			counter, _ := strconv.ParseInt(string(s[i]), 10, 32)
			if counter > 1 {
				temp := strings.Repeat(string(lastRune), int(counter-1))
				unpackString.WriteString(temp)
			} else {
				lastOp = false
			}
		} else {
			lastOp = false
			unpackString.WriteRune(rune(s[i]))
			lastRune = rune(s[i])
		}
	}

	if lastOp {
		unpackString.WriteRune(lastRune)
	}

	return unpackString.String()
}
