package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
		expected []string
	} {
		{
			input: "hello world",
			expected: []string {
				"hello",
				"world",
			},
		},
		{
			input: "HELLO world",
			expected: []string {
				"hello",
				"world",
			},
		},
	}

	for _, cs := range cases{
		actual := clearInputs(cs.input)
		if len(actual) != len(cs.expected) {
			t.Errorf("The lengths are not equal: %v not equal to %v", 
					len(actual),
					len(cs.expected),
				)
				continue
		}
		for i := range actual {
			actualWord := actual[i]
			if actualWord != cs.expected[i] {
				t.Errorf("%v does not equal %v", 
				actualWord, 
				cs.expected[i],
			)
			}
		}
	}
}