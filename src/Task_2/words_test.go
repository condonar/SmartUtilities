package main

import (
	"strings"
	"testing"
)

func TestBasic(t *testing.T) {
	resSlice := GetTopWords("aa bb cc aa cc cc cc aa ab ac bb", 3)
	actual := strings.Join(resSlice, " ")
	expected := "cc aa bb"

	if actual != expected {
		t.Errorf("\nExpected: %s\nActual: %s\n", expected, actual)
	}
}

func TestBigK(t *testing.T) {
	resSlice := GetTopWords("aa cc cc bb bb dd kk", 20)
	actual := strings.Join(resSlice, " ")
	expected := "bb cc aa dd kk"

	if actual != expected {
		t.Errorf("\nExpected: %s\nActual: %s\n", expected, actual)
	}
}

func TestGetTopWords(t *testing.T) {
	tests := []struct {
		name     string
		inputStr string
		K        int
		expected string
	}{
		{"Basic", "aa bb cc aa cc cc cc aa ab ac bb", 3, "cc aa bb"},
		{"Big num K", "bb bb aa aa cc dd dd", 10, "aa bb dd cc"},
		{"Empty input", "", 3, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := GetTopWords(tt.inputStr, tt.K)
			actual := strings.Join(res, " ")
			if actual != tt.expected {
				t.Errorf("\nExpected: %s\nActual: %s\n", tt.expected, actual)
			}
		})
	}
}
