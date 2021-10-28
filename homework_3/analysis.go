package homework_3

import (
	"sort"
	"strings"
)

func frequencyAnalysis(text string) []string {
	data := make(map[string]int)

	for _, value := range strings.Split(text, " ") {
		value := strings.ToLower(value)
		data[value]++
	}

	keys := make([]string, 0, len(data))

	for key := range data {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(lhs int, rhs int) bool {
		if data[keys[lhs]] != data[keys[rhs]] {
			return data[keys[lhs]] > data[keys[rhs]]
		}
		return keys[lhs] < keys[rhs]
	})

	const maxSize = 10
	if len(keys) > maxSize {
		keys = keys[:maxSize]
	}

	return keys
}
