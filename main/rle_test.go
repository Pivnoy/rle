package main

import (
	"testing"
)

func TestRle1(t *testing.T) {
	actual := []string{"A"}
	expected := []string{"A"}

	rle := NewRleDecoder()
	result := rle.RunLengthEncode(actual)

	if len(expected) != len(result) {
		t.Errorf("Error in result len. Expected  - %d, got - %d", len(actual), len(result))
	}

	for j := 0; j < len(expected); j++ {
		if expected[j] != result[j] {
			t.Errorf("Error in string equalence. Expected  - %s, got - %s", expected[j], result[j])
		}
	}
}

func TestRle2(t *testing.T) {
	actual := []string{"A",
		"AA",
		"AAAAAAA"}
	expected := []string{"A",
		"2A",
		"7A"}

	rle := NewRleDecoder()
	result := rle.RunLengthEncode(actual)

	if len(expected) != len(result) {
		t.Errorf("Error in result len. Expected  - %d, got - %d", len(actual), len(result))
	}

	for j := 0; j < len(expected); j++ {
		if expected[j] != result[j] {
			t.Errorf("Error in string equalence. Expected  - %s, got - %s", expected[j], result[j])
		}
	}
}

func TestRle3(t *testing.T) {
	actual := []string{
		"A",
		"AA",
		"AAAAAAA",
		"BAAAAAAA",
		"AAAAAA",
		"CAAABAAAAAA   A",
		"ACD   AAABAAAAAAAAABAAA",
		"DDDAAAAAAABBBBBBBBBAAAAA AAAA",
		"AAABAAABBBBBBBBBBBBBBBAAAAAA",
		"AABBAAABBBBBBBBBBBBBBBBBBBBAAB",
		"AAAABBBBBBBBBBBBBBBBBBBBBBBBAAA",
		"BACBAABBBBBBBBBBBBBBBBBBBBBBBBBBAADD",
		"C         AAAABBBBBBBBBBBCCCCCCCCCCBBBBBBBAA",
		"ABBACBCAD B     BAABBBBBBBBBBCCCCCCCCCCCCCCBBBBBAB",
		"CCDDDDCCCCB   AAABBBBBBBBBBCCCCCCCDDDDCCCCCBBBBABA",
		"DCDDDDDDDDDCCB BAAABBBBBBBBBCCCCCDDDDDDDDDDCCCBBBAB",
		"DDDDDDDDDDDDCCBAAABBBBBBBBBCCCCCDDDDDDDDDDDDCCBBBAA",
		"ABBCDDDDDDDDDDDDCCBAAABBBBBBBBBCCCCCDDDDDDDDDDDDCCBBAA",
		"AAAADAAAABAAAAADBAAABBBCDDDDDDDDDDDDCCBBAABBBBBBBBBCCCCDDDDDDDDDDDDDCCB"}
	expected := []string{
		"A",
		"2A",
		"7A",
		"B7A",
		"6A",
		"C3AB6A3 A",
		"ACD3 3AB9AB3A",
		"3D7A9B5A 4A",
		"3AB3A15B6A",
		"2A2B3A20B2AB",
		"4A24B3A",
		"BACB2A26B2A2D",
		"C9 4A11B10C7B2A",
		"A2BACBCAD B5 B2A10B14C5BAB",
		"2C4D4CB3 3A10B7C4D5C4BABA",
		"DC9D2CB B3A9B5C10D3C3BAB",
		"12D2CB3A9B5C12D2C3B2A",
		"A2BC12D2CB3A9B5C12D2C2B2A",
		"4AD4AB5ADB3A3BC12D2C2B2A9B4C13D2CB"}

	rle := NewRleDecoder()
	result := rle.RunLengthEncode(actual)

	if len(expected) != len(result) {
		t.Errorf("Error in result len. Expected  - %d, got - %d", len(actual), len(result))
	}

	for j := 0; j < len(expected); j++ {
		if expected[j] != result[j] {
			t.Errorf("Error in string equalence. Expected  - %s, got - %s", expected[j], result[j])
		}
	}
}
