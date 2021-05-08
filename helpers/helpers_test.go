package helpers

import "testing"

var tests = []struct {
	name        string
	dividend    float32
	divisor     float32
	expectedVal float32
	isError     bool
}{
	{"valid data", 10.0, 2.0, 5.0, false},
	{"valid data simple", 15.0, 3.0, 5.0, false},
	{"valid data 2", 30.0, 2.0, 15.0, false},
	{"not valid data", 10.0, 0.0, 0.0, true},
}

func TestDivision(t *testing.T) {
	for _, tt := range tests {
		got, err := Divive(tt.dividend, tt.divisor)
		if tt.isError {
			if err == nil {
				t.Error("expected error and didnt get it")
			}
		} else {
			if err != nil {
				t.Error("did not expecte an error")
			}
		}

		if got != tt.expectedVal {
			t.Errorf("expected %f but got %f instead", tt.expectedVal, got)
		}
	}
}

