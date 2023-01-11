package accelhints

import (
	"testing"

	"golang.org/x/exp/slices"
	// "github.com/mark-summerfield/gong"
	// "golang.org/x/exp/maps"
	// "golang.org/x/exp/slices"
)

// maps.Equal() & maps.EqualFunc() & slices.Equal() & slices.EqualFunc()
// https://pkg.go.dev/golang.org/x/exp/maps
// https://pkg.go.dev/golang.org/x/exp/slices
// gong.IsRealClose() & gong.IsRealZero()

func Test001(t *testing.T) {
	original := []string{"O_ne", "_Two"}
	expected := []string{"O&ne", "&Two"}
	actual := slices.Clone(original)
	err := AddHints(actual)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
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
	err := AddHints(actual)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	for i := 0; i < len(original); i++ {
		if actual[i] != expected[i] {
			t.Errorf("expected %q, got %q", expected[i], actual[i])
		}
	}
}
