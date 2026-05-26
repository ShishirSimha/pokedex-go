package main

import (
	"reflect"
	"testing"
)

func TestCleanInput(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
		wantErr  bool
	}{
		{
			name:     "Simple lowercase words",
			input:    "hello world",
			expected: []string{"hello", "world"},
			wantErr:  false,
		},
		{
			name:     "Extraneous surrounding and middle spacing",
			input:    "  hello   world  ",
			expected: []string{"hello", "world"},
			wantErr:  false,
		},
		{
			name:     "Mixed casing",
			input:    "HELLO WORLD",
			expected: []string{"hello", "world"},
			wantErr:  false,
		},
		{
			name:     "Empty input error case",
			input:    "",
			expected: nil,
			wantErr:  true,
		},
		{
			name:     "Spacing only input error case",
			input:    "     ",
			expected: nil,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := cleanInput(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("cleanInput(%q) error = %v; wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("cleanInput(%q) = %q; expected %q", tt.input, actual, tt.expected)
			}
		})
	}
}
