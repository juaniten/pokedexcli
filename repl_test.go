package main

import (
    "fmt"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "   hello  world    ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "HeLlo WorlD",
			expected: []string{"hello", "world"},
		},

		{
			input:    "    hello",
			expected: []string{"hello"},
		},
		{
			input:    "hello hello hello world",
			expected: []string{"hello", "hello", "hello", "world"},
		},
		{
			input:    "    HeLLO   hello  HELLO    WORLD       ",
			expected: []string{"hello", "hello", "hello", "world"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("lenght of output does not match expected")
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf(fmt.Sprintf("%v != %v, expected %v", actual, expectedWord, expectedWord))
			}
		}
	}
}
