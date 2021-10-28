package homework_2

import "testing"

func TestUnpackTask1(t *testing.T) {
	text := "a4bc2d5e"
	expected := "aaaabccddddde"

	if result := Unpack(text); result != expected {
		t.Fatalf("bad unpack for [%s]: got %s, expected %s", text, result, expected)
	}
}

func TestUnpackTask2(t *testing.T) {
	text := "abcd"
	expected := "abcd"

	if result := Unpack(text); result != expected {
		t.Fatalf("bad unpack for [%s]: got %s, expected %s", text, result, expected)
	}
}

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

func TestUnpackMatchCount(t *testing.T) {
	text := "a10"
	expected := "aaaaaaaaaa"

	if result := Unpack(text); result != expected {
		t.Fatalf("bad unpack for [%s]: got %s, expected %s", text, result, expected)
	}
}

func TestUnpackMatchCountAndAlone(t *testing.T) {
	text := "a10b"
	expected := "aaaaaaaaaab"

	if result := Unpack(text); result != expected {
		t.Fatalf("bad unpack for [%s]: got %s, expected %s", text, result, expected)
	}
}

func TestUnpackMatchCountAndAlone2(t *testing.T) {
	text := "a10b1"
	expected := "aaaaaaaaaab"

	if result := Unpack(text); result != expected {
		t.Fatalf("bad unpack for [%s]: got %s, expected %s", text, result, expected)
	}
}

func TestUnpackMatchCountAndAlone3(t *testing.T) {
	text := "a10b2"
	expected := "aaaaaaaaaabb"

	if result := Unpack(text); result != expected {
		t.Fatalf("bad unpack for [%s]: got %s, expected %s", text, result, expected)
	}
}
