package main

import (
	"reflect"
	"testing"
)

func TestSearchAnagram(t *testing.T) {
	tests := []struct {
		words        []string
		expected     map[string][]string
	}{
		{[]string{"пятак", "лИсток", "пЯтка", "acbd", "тяпка", "слиток", "столик", "abcd"},
			map[string][]string{"пятак": {"пятка", "тяпка"}, "слиток": {"лИсток", "столик"}, "abcd": {"acbd"}},
		},
		{[]string{"hello", "world", "olleh", "dlrow"},
			map[string][]string{"hello": {"olleh"}, "world": {"dlrow"}},
		},
	}

	for _, test := range tests {
		result := searchAnagram(test.words)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For input %v, expected %v, but got %v", test.words, test.expected, result)
		}
	}
}

func TestCountLettersUnicode(t *testing.T) {
	tests := []struct {
		input    string
		expected map[string]int
	}{
		{"пятак", map[string]int{"п": 1, "я": 1, "т": 1, "а": 1, "к": 1}},
		{"пЯтка", map[string]int{"п": 1, "я": 1, "т": 1, "к": 1, "а": 1}},
		{"abcd", map[string]int{"a": 1, "b": 1, "c": 1, "d": 1}},
	}

	for _, test := range tests {
		result, length := countLettersUnicode(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For input %s, expected %v, but got %v", test.input, test.expected, result)
		}
		if length != len(test.input) {
			t.Errorf("For input %s, expected length %d, but got %d", test.input, len(test.input), length)
		}
	}
}

func TestEqualMap(t *testing.T) {
	map1 := map[string]int{"a": 1, "b": 2, "c": 3}
	map2 := map[string]int{"a": 1, "b": 2, "c": 3}
	map3 := map[string]int{"a": 1, "b": 2}
	map4 := map[string]int{"a": 1, "b": 3, "c": 3}

	if !equalMap(map1, map2) {
		t.Error("Expected map1 and map2 to be equal")
	}

	if equalMap(map1, map3) {
		t.Error("Expected map1 and map3 to be different")
	}

	if equalMap(map1, map4) {
		t.Error("Expected map1 and map4 to be different")
	}
}
