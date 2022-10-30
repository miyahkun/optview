package main

import (
	"testing"
)

func AllRuneEqual(left []rune, right []rune) bool {
	for i, v := range left {
		if v != right[i] {
			return false
		}
	}
	return true
}

func TestNewParseState(t *testing.T) {
	parser_state := NewParseState("sampletext")
	actual := parser_state.buffer
	expected := []rune("sampletext")
	if !AllRuneEqual(actual, expected) {
		t.Errorf("got: %v\nwant: %v\n", actual, expected)
	}
}
