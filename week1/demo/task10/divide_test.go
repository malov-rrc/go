package main

import "testing"

func TestDivide(t *testing.T) {
	tests := []struct {
		name        string
		firstInput  float64
		secondInput float64
		wantValue   float64
		wantErr     bool
	}{
		{"Деление на 0", 2, 0, 0, true},
		{"обычное деление", 5, 2, 2.5, false},
		{"Деление отрицательных чисел", -5, -2, 2.5, false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := Divide(tc.firstInput, tc.secondInput)
			if tc.wantErr && err == nil {
				t.Error("ожидалась ошибка, но получили nil")
			}
			if !tc.wantErr && err != nil {
				t.Errorf("не ожидали ошибку, получили %v", err)
			}
			if got != tc.wantValue {
				t.Errorf("Divide(%.2f, %.2f) = %.2f, want %.2f", tc.firstInput, tc.secondInput, got, tc.wantValue)
			}
		})
	}
}
