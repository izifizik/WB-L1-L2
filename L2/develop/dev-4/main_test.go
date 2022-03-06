package main

import (
	"testing"
)

func TestSetAnagrams(t *testing.T) {
	testTable := []struct {
		input    []string
		expected map[string]*[]string
	}{
		{
			input:    []string{"пятак", "Пятка", "тяпка", "листок", "слиток", "стоЛик", "слиток", "тяПка"},
			expected: map[string]*[]string{"листок": {"листок", "слиток", "столик"}, "пятак": {"пятак", "пятка", "тяпка"}},
		},
		{
			input:    []string{"ПОЛКОВНИК", "ТЕРПЕЛИВОСТЬ", "клоповник", "ПРОСВЕТИТЕЛЬ"},
			expected: map[string]*[]string{"просветитель": {"просветитель", "терпеливость"}, "клоповник": {"клоповник", "полковник"}},
		},
	}

	for _, tt := range testTable {
		got := SetAnagrams(tt.input)
		if Equal(got, tt.expected) {
			t.Errorf("error")
		}
	}
}

func Equal(a, b map[string]*[]string) bool {
	for k, v := range a {
		if b[k] != v {
			return false
		}
	}
	return true
}
