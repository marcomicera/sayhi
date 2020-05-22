package stringutils

import (
	"testing"
)

func TestSplitCamelCase(t *testing.T) {

	tests := []struct {
		name     string
		input    string
		expected string
		error    bool
	}{
		{
			name:  "Empty string",
			input: "",
			error: true,
		},
		{
			name:     "Requirement example",
			input:    "AlfredENeumann",
			expected: "Alfred E Neumann",
		},
		{
			name:     "UpperCamel case",
			input:    "JackWhite",
			expected: "Jack White",
		},
		{
			name:     "lowerCamel case",
			input:    "jackWhite",
			expected: "jack White",
		},
		{
			name:     "Name with acronym in the middle",
			input:    "FredFDDurst",
			expected: "Fred FD Durst",
		},
		{
			name:  "Invalid characters",
			input: "\x990\x980",
			error: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := SplitCamelCasePersonName(tt.input)

			if err != nil && tt.error != true {
				t.Errorf("SplitCamelCasePersonName() = did not expect error %v", err)
			}

			if err == nil && got != tt.expected {
				t.Errorf("SplitCamelCasePersonName() = got %v, expected %v", got, tt.expected)
			}
		})
	}
}
