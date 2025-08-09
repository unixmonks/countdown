package main

import (
	"testing"
	"time"
)

func TestParseDuration(t *testing.T) {
	tests := []struct {
		input    string
		expected time.Duration
		hasError bool
	}{
		{"30s", 30 * time.Second, false},
		{"5m", 5 * time.Minute, false},
		{"1h", time.Hour, false},
		{"1h30m", time.Hour + 30*time.Minute, false},
		{"1m30s", time.Minute + 30*time.Second, false},
		{"2h45m30s", 2*time.Hour + 45*time.Minute + 30*time.Second, false},
		{"invalid", 0, true},
		{"", 0, true},
		{"10", 0, true},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result, err := parseDuration(test.input)
			
			if test.hasError {
				if err == nil {
					t.Errorf("Expected error for input %q, but got none", test.input)
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error for input %q: %v", test.input, err)
				}
				if result != test.expected {
					t.Errorf("Expected %v for input %q, got %v", test.expected, test.input, result)
				}
			}
		})
	}
}

func TestFormatDuration(t *testing.T) {
	tests := []struct {
		input    time.Duration
		expected string
	}{
		{30 * time.Second, "00:30"},
		{5 * time.Minute, "05:00"},
		{time.Minute + 30*time.Second, "01:30"},
		{time.Hour, "01:00:00"},
		{time.Hour + 30*time.Minute, "01:30:00"},
		{time.Hour + 30*time.Minute + 45*time.Second, "01:30:45"},
		{2*time.Hour + 45*time.Minute + 30*time.Second, "02:45:30"},
		{10*time.Hour + 5*time.Minute + 15*time.Second, "10:05:15"},
		{0, "00:00"},
	}

	for _, test := range tests {
		t.Run(test.expected, func(t *testing.T) {
			result := formatDuration(test.input)
			if result != test.expected {
				t.Errorf("Expected %q for duration %v, got %q", test.expected, test.input, result)
			}
		})
	}
}