package accelhints

import (
	"testing"

	"golang.org/x/exp/slices"
)

func Test001(t *testing.T) {
	original := []string{"O_ne", "_Two"}
	expected := []string{"O&ne", "&Two"}
	actual := slices.Clone(original)
	n, err := AddHints(actual)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	if n != len(original) {
		t.Errorf("expected %d accelrated got %d", n, len(original))
	}
	for i := 0; i < len(original); i++ {
		if actual[i] != expected[i] {
			t.Errorf("expected %q, got %q", expected[i], actual[i])
		}
	}
}

func Test002(t *testing.T) {
	original := []string{
		"Undo",
		"Redo",
		"Copy",
		"Cu&t",
		"Paste",
		"Find",
		"Find Again",
	}
	expected := []string{
		"&Undo",
		"&Redo",
		"&Copy",
		"Cu&t",
		"&Paste",
		"&Find",
		"Find &Again"}
	actual := slices.Clone(original)
	n, err := AddHints(actual)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	if n != len(original) {
		t.Errorf("expected %d accelrated got %d", n, len(original))
	}
	for i := 0; i < len(original); i++ {
		if actual[i] != expected[i] {
			t.Errorf("expected %q, got %q", expected[i], actual[i])
		}
	}
}
