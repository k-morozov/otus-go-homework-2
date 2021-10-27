package foobar

func Count(s string, r rune) int {
	var count int
	for _, value := range s {
		if value == r {
			count++
		}
	}

	return count
}
