package homework_2

import "testing"

func TestUnpackSimpleEmpty(t *testing.T) {
	text := ""
	expected := ""

	if result := Unpack(text); result != expected {
		t.Fatalf("bad unpack for [%s]: got %s, expected %s", text, result, expected)
	}
}

func TestUnpackSimpleAlone(t *testing.T) {
	text := "a"
	expected := "a"

	if result := Unpack(text); result != expected {
		t.Fatalf("bad unpack for [%s]: got %s, expected %s", text, result, expected)
	}
}

func TestUnpackOne(t *testing.T) {
	text := "a1"
	expected := "a"

	if result := Unpack(text); result != expected {
		t.Fatalf("bad unpack for [%s]: got %s, expected %s", text, result, expected)
	}
}

func TestUnpackTwoSym(t *testing.T) {
	text := "ab"
	expected := "ab"

	if result := Unpack(text); result != expected {
		t.Fatalf("bad unpack for [%s]: got %s, expected %s", text, result, expected)
	}
}

func TestUnpackTwoSymTwo(t *testing.T) {
	text := "a2b"
	expected := "aab"

	if result := Unpack(text); result != expected {
		t.Fatalf("bad unpack for [%s]: got %s, expected %s", text, result, expected)
	}
}

func TestUnpackTwoSymForTwo(t *testing.T) {
	text := "a2b3"
	expected := "aabbb"

	if result := Unpack(text); result != expected {
		t.Fatalf("bad unpack for [%s]: got %s, expected %s", text, result, expected)
	}
}

func TestUnpackFourSymForTwo(t *testing.T) {
	text := "a2b3c4d"
	expected := "aabbbccccd"

	if result := Unpack(text); result != expected {
		t.Fatalf("bad unpack for [%s]: got %s, expected %s", text, result, expected)
	}
}
