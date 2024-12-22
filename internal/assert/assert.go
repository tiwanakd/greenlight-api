package assert

import (
	"strings"
	"testing"
)

func Equal[T comparable](t *testing.T, actual, expected T) {
	t.Helper()

	if actual != expected {
		t.Errorf("got: %v; want %v", actual, expected)
	}
}

func StringContains(t *testing.T, acutal, expextedSubstring string) {
	t.Helper()

	if !strings.Contains(acutal, expextedSubstring) {
		t.Errorf("got: %q; expected to contain: %q", acutal, expextedSubstring)
	}
}
