package frequencyAnalysis

import "testing"

func compare(lhs []string, rhs []string) bool {
	flag := true
	if len(lhs) == len(rhs) {
		for index, value := range rhs {
			if lhs[index] != value {
				flag = false
				break
			}
		}
	} else {
		flag = false
	}
	return flag
}

func TestFrequencyAnalysis1(t *testing.T) {
	text := "a b c"
	expected := []string{"a", "b", "c"}
	result := frequencyAnalysis(text)

	if !compare(result, expected) {
		t.Fatalf("bad requencyAnalysis for [%s]: got %s, expected %s", text, result, expected)
	}
}

func TestFrequencyAnalysis2(t *testing.T) {
	text := "a b c aa aa"
	expected := []string{"aa", "a", "b", "c"}
	result := frequencyAnalysis(text)

	if !compare(result, expected) {
		t.Fatalf("bad requencyAnalysis for [%s]: got %s, expected %s", text, result, expected)
	}
}

func TestFrequencyAnalysis3(t *testing.T) {
	text := "a b c aa aa aa c d e f g h i k"
	expected := []string{"aa", "c", "a", "b", "d", "e", "f", "g", "h", "i"}
	result := frequencyAnalysis(text)

	if !compare(result, expected) {
		t.Fatalf("bad requencyAnalysis for [%s]: got %s, expected %s", text, result, expected)
	}
}
