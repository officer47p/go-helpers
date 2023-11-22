package gohelpers

import "testing"

func TestFloatToString(t *testing.T) {
	tests := []struct {
		name  string
		input float64
		want  string
	}{
		{name: "only decimals", input: 0.222, want: "0.222"},
		{name: "decimals with integer", input: 12.23003, want: "12.23003"},
	}

	for _, test := range tests {
		got := FloatToString(test.input)
		if got != test.want {
			t.Errorf("%s: expected: %v, got: %v", test.name, test.want, got)
		}
	}
}

func TestStringToFloat(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  float64
	}{
		{name: "only decimals", input: "0.222", want: 0.222},
		{name: "decimals with integer", input: "12.23003", want: 12.23003},
	}

	for _, test := range tests {
		got, err := StringToFloat(test.input)
		if err != nil {
			t.Fatalf("error while converting string to float. Err: %s", err.Error())
		}
		if got != test.want {
			t.Errorf("%s: expected: %v, got: %v", test.name, test.want, got)
		}
	}
}
