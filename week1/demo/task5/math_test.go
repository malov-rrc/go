package main

import "testing"

func TestMin(t *testing.T) {
	tests := []struct {
		name      string // описание теста (опционально, но полезно)
		input     []int  // входные данные
		wantValue int    // ожидаемый результат
		wantErr   bool
	}{
		{"пустой срез", []int{}, 0, true},
		{"один элемент", []int{5}, 5, false},
		{"несколько положительных", []int{1, 3, 2}, 1, false},
		{"с отрицательными", []int{-1, -5, -2}, -5, false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := Min(tc.input)
			if tc.wantErr && err == nil {
				t.Error("ожидалась ошибка, но получили nil")
			}
			if !tc.wantErr && err != nil {
				t.Errorf("не ожидали ошибку, получили %v", err)
			}
			if got != tc.wantValue {
				t.Errorf("Min(%v) = %v, want %v", tc.input, got, tc.wantValue)
			}
		})
	}
}

func TestMax(t *testing.T) {
	tests := []struct {
		name      string
		input     []int
		wantValue int
		wantErr   bool
	}{
		{"пустой срез", []int{}, 0, true},
		{"один элемент", []int{5}, 5, false},
		{"несколько положительных", []int{1, 3, 2}, 3, false},
		{"с отрицательными", []int{-1, -5, -2}, -1, false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := Max(tc.input)
			if tc.wantErr && err == nil {
				t.Error("ожидалась ошибка, но получили nil")
			}
			if !tc.wantErr && err != nil {
				t.Errorf("не ожидали ошибку, получили %v", err)
			}
			if got != tc.wantValue {
				t.Errorf("Max(%v) = %v, want %v", tc.input, got, tc.wantValue)
			}
		})
	}
}

func TestSum(t *testing.T) {
	tests := []struct {
		name      string
		input     []int
		wantValue int
		wantErr   bool
	}{
		{"пустой срез", []int{}, 0, true},
		{"один элемент", []int{5}, 5, false},
		{"несколько положительных", []int{1, 3, 2}, 6, false},
		{"с отрицательными", []int{-1, -5, -2}, -8, false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := Sum(tc.input)
			if tc.wantErr && err == nil {
				t.Error("ожидалась ошибка, но получили nil")
			}
			if !tc.wantErr && err != nil {
				t.Errorf("не ожидали ошибку, получили %v", err)
			}
			if got != tc.wantValue {
				t.Errorf("Sum(%v) = %v, want %v", tc.input, got, tc.wantValue)
			}
		})
	}
}

func TestIsPrime(t *testing.T) {
	tests := []struct {
		name   string
		input  int
		result bool
	}{
		{"тест -1", -1, false},
		{"тест 0", 0, false},
		{"тест 1", 1, false},
		{"тест 2", 2, true},
		{"тест 4", 4, false},
		{"тест 5", 5, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := IsPrime(tc.input)
			if got != tc.result {
				t.Errorf("isPrime(%v) = %v, want %v", tc.input, got, tc.result)
			}
		})
	}
}
