package main

import (
	"testing"
)

func TestStringUnpack(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		wantErr  bool
	}{
		{"a4bc2d5e", "aaaabccddddde", false},
		{"abcd", "abcd", false},
		{"45", "", true},
		{"", "", false},
		{"qwe\\4\\5", "qwe45", false},
		{"m2a3m6y8", "mmaaaammmmmmyyyyyyyy", false},
		{"6e", "", true},
	}

	for _, test := range tests {
		unpacked, err := stringUnpack(test.input)
		if (err != nil) != test.wantErr {
			t.Errorf("stringUnpack(%q) error = %v, wantErr %v", test.input, err, test.wantErr)
			continue
		}
		if unpacked != test.expected {
			t.Errorf("stringUnpack(%q) = %q, want %q", test.input, unpacked, test.expected)
		}
	}
}
