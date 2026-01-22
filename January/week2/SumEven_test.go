package week2

import (
	"testing"
)

func TestSumEven(t *testing.T) {
	tests := []struct {
		name string

		input    []int
		expected int
	}{
		{
			name:     "normal numbers",
			input:    []int{1, 2, 3, 4, 5, 6},
			expected: 12,
		},
		{
			name:     "Odd numbers only!",
			input:    []int{1, 3, 5, 7, 11},
			expected: 0,
		},
		{
			name:     "Even numbers only!",
			input:    []int{2, 4, 20, 100, 6, 8},
			expected: 140,
		},

		{
			name:     "Empty input",
			input:    []int{},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SumEven(tt.input)
			if result != tt.expected {
				t.Fatalf("expected %d, got %d", tt.expected, result)
			}
		})

	}
}
