package main

import (
	"reflect"
	"testing"
)

func TestCleanInput(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
	}{
		{
			input:    "hello world",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  hello   world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "HELLO WORLD",
			expected: []string{"hello", "world"},
		},
		{
			input:    "",
			expected: []string{},
		},
	}

	for _, test := range tests {
		actual := cleanInput(test.input)
		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("cleanInput(%q) = %q; expected %q", test.input, actual, test.expected)
		}
	}
}
