package main

import "testing"

func TestRomanToIn(t *testing.T) {
	cases := []struct {
		input    string
		expected int
	}{
		{input: "III", expected: 3},
		{input: "LVIII", expected: 58},
		{input: "MCMXCIV", expected: 1994},
	}

	for _, v := range cases {
		res := romanToInt(v.input)
		if res != v.expected {
			t.Errorf("%s expected %d but got %d", v.input, v.expected, res)
		}
	}
}
