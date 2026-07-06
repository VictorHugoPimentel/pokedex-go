package main

import "testing"

func TestCleanInput(t *testing.T) {
	input := "Hello, World!"
	expected := []string{"hello,", "world!"}
	result := cleanInput(input)
	if len(result) != len(expected) {
		t.Errorf("Expected %d words, got %d", len(expected), len(result))
	}
	for i, word := range expected {
		if result[i] != word {
			t.Errorf("Expected word %d to be %s, got %s", i, word, result[i])
		}
	}
}