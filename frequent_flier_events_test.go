package main

import (
	"testing"
)

func TestStatusString(t *testing.T) {
	tests := []struct {
		status   Status
		expected string
	}{
		{StatusRed, "StatusRed"},
		{StatusSilver, "StatusSilver"},
		{StatusGold, "StatusGold"},
	}

	for _, test := range tests {
		if test.status.String() != test.expected {
			t.Errorf("Expected %s, got %s", test.expected, test.status.String())
		}
	}
}
