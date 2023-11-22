package gohelpers

import "testing"

func TestFloatToString(t *testing.T) {
	want := "0.222"
	got := FloatToString(0.222)

	if got != want {
		t.Errorf("got %s, but want %s", got, want)
	}
}
