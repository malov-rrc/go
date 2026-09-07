package utils

import "testing"

func TestParseLine(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		command string
		amount  float64
		wantErr bool
	}{
		{"valid string", "WITHDRAW 90", "WITHDRAW", 90, false},
		{"invalid amount", "WITHDRAW f", "", 0, true},
		{"more than 2 words", "WITHDRAW WITHDRAW WITHDRAW", "", 0, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			gotCommand, gotAmount, err := ParseLine(tc.input)
			if tc.wantErr && err == nil {
				t.Error("Error is expected, but got nil")
			}
			if !tc.wantErr && err != nil {
				t.Errorf("Unexpected error, got %v", err)
			}
			if gotCommand != tc.command {
				t.Errorf("ParseLine(\"%s\") command invalid. Expected: %s, got %s", tc.input, tc.command, gotCommand)
			}
			if gotAmount != tc.amount {
				t.Errorf("ParseLine(\"%s\") amount invalid. Expected: %.2f, got %.2f", tc.input, tc.amount, gotAmount)
			}
		})
	}
}
