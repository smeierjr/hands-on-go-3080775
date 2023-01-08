// challenges/testing/begin/main_test.go
package main

import "testing"

// write a test for letterCounter.count
func TestLetterCounter(t *testing.T) {

	var tests = []struct {
		name  string
		input string
		want  int
	}{
		{"one", "", 0},
		{"two", "a2 32 ^ &/)", 1},
		{"three", "812 %6ab//", 2},
	}

	lc := letterCounter{}
	for _, tc := range tests {

		t.Run(tc.name, func(t *testing.T) {
			got := letterCounter.count(lc, tc.input)
			if got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}

// write a test for numberCounter.count
func TestNumberCounter(t *testing.T) {

	var tests = []struct {
		name  string
		input string
		want  int
	}{
		{"one", "", 0},
		{"two", "a2 32 ^ &/)", 3},
		{"three", "812 %6ab//", 4},
	}

	for _, tc := range tests {
		var n numberCounter
		t.Run(tc.name, func(t *testing.T) {
			got := numberCounter.count(n, tc.input)
			if got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}

// write a test for symbolCounter.count
func TestSymbolCounter(t *testing.T) {

	var tests = []struct {
		name  string
		input string
		want  int
	}{
		{"one", "", 0},
		{"two", "abc 1,?/", 4},
		{"three", "abc 12&8 ^_", 5},
	}

	for _, tc := range tests {
		var s symbolCounter
		t.Run(tc.name, func(t *testing.T) {
			got := symbolCounter.count(s, tc.input)
			if got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}
