package main

import "testing"

func TestStringBuilder(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
		err      string
	}{
		{"a4bc2d5e", "aaaabccddddde", "(некорректная строка)"},
		{"abcd", "abcd", "(некорректная строка)"},
		{"45", "", "(некорректная строка)"},
		{"", "", "(некорректная строка)"},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result, err := stringBuilder([]rune(tc.input))
			if err != nil && err.Error() != tc.err {
				t.Errorf("New Error: %v", err)
			}
			if result != tc.expected {
				t.Errorf("Expected: %s, Real: %s", tc.expected, result)
			}
		})
	}
}

func TestStringBuilderWithEscape(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
		err      string
	}{
		{"qwe\\4\\5", "qwe45", "(некорректная строка)"},
		{"qwe\\45", "qwe44444", "(некорректная строка)"},
		{"qwe\\\\5", "qwe\\\\\\\\\\", "(некорректная строка)"},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result, err := stringBuilderWithEscape([]rune(tc.input))
			if err != nil && err.Error() != tc.err {
				t.Errorf("New Error: %v", err)
			}
			if result != tc.expected {
				t.Errorf("Expected: %s, Real: %s", tc.expected, result)
			}
		})
	}
}
