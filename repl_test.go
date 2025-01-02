package main

import (
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
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "   	",
			expected: []string{},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("lenght of output (%v) does not match expected (%v)", len(actual), len(c.expected))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("%s != %s, expected %s", word, expectedWord, expectedWord)
			}
		}
	}
}
