package main

import "testing"

func TestParseLine(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		wantValue string
	}{
		//пробелы по краям, смешанный регистр, несколько пробелов подряд внутри, уже чистая строка
		{"spaces around", "   hello ", "hello"},
		{"mixed register", "HeLlO WorLd", "hello world"},
		{"few spaces inside", "hello           world                      2", "hello world 2"},
		{"clean string", "clean string lol", "clean string lol"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := CleanInput(tc.input)
			if got != tc.wantValue {
				t.Errorf("CleanInput(\"%s\") = %v, want %v", tc.input, got, tc.wantValue)
			}
		})
	}
}
