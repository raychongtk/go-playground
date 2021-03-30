package unittest

import (
	"testing"
)

func TestFindMax(t *testing.T) {
	got := FindMax([]int{0, 12, 3, 4, 5, 62, 7, 8, 29})
	want := 62

	assertMax(t, got, want)
}

func assertMax(t *testing.T, got int, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
