// testing/table-driven/begin/main_test.go
package main

import "testing"

func TestSum(t *testing.T) {
	var tests = []struct {
		name  string
		input string
		want  int
	}{
		{"one", "hello", 5},
		{"two", "steven", 6},
		{"three", "priscilla", 9},
	}

	

	// range over the tests and run them as subtests
	//
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := sum(tc.input...)
			if got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}
