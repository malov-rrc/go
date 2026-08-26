package main

import "testing"

func TestFactorial(t *testing.T) {
	tests := []struct {
		name      string
		input     int
		wantValue int
		wantErr   bool
	}{
		{"Number less than 0", -1, 0, true},
		{"Number 0", 0, 1, false},
		{"Number 1", 1, 1, false},
		{"Number 5", 5, 120, false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := Factorial(tc.input)
			if tc.wantErr && err == nil {
				t.Error("Wanted error, but got nil")
			}
			if !tc.wantErr && err != nil {
				t.Errorf("Not expected error, got: %v", err)
			}
			if got != tc.wantValue {
				t.Errorf("Factorial(%v) = %v, want %v", tc.input, got, tc.wantValue)
			}
		})
	}
}
