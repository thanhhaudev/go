package main

import (
	"strings"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		input    []string
		expected string
	}{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
		{[]string{"interspecies", "interstellar", "interstate"}, "inters"},
		{[]string{"throne", "throne"}, "throne"},
		{[]string{"throne", "dungeon"}, ""},
		{[]string{"", ""}, ""},
		{[]string{"a"}, "a"},
		{[]string{"ab", "a"}, "a"},
		{[]string{"", "b"}, ""},
		{[]string{"b", ""}, ""},
	}

	for _, test := range tests {
		result := longestCommonPrefix(test.input)
		if result != test.expected {
			t.Errorf("For input %v, expected %q but got %q", test.input, test.expected, result)
		}
	}
}

func BenchmarkLongestCommonPrefix(b *testing.B) {
	// Create a large input for performance testing
	input := make([]string, 1000)
	for i := range input {
		input[i] = strings.Repeat("a", 1000)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		longestCommonPrefix(input)
	}
}
