package accelhints

import (
	"testing"
)

func Test001(t *testing.T) {
	original := []string{"O&ne", "Two"}
	expected := []string{"O&ne", "&Two"}
	actual, err := Hints(original)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	_, n := Indexes(actual)
	if n != 2 {
		t.Errorf("expected 2 accelrated got %d", n)
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
	actual, err := Hints(original)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	_, n := Indexes(actual)
	if n != 7 {
		t.Errorf("expected 7 accelrated got %d", n)
	}
	for i := 0; i < len(original); i++ {
		if actual[i] != expected[i] {
			t.Errorf("expected %q, got %q", expected[i], actual[i])
		}
	}
}

func Test003(t *testing.T) {
	original := []string{
		"Undo",
		"Redo",
		"Copy",
		"Cu&t",
		"Paste",
		"Find",
		"Find Again",
		"Find && Replace"}
	expected := []string{
		"&Undo",
		"&Redo",
		"&Copy",
		"Cu&t",
		"&Paste",
		"&Find",
		"Find &Again",
		"F&ind && Replace"}
	expectedIndexes := []int{0, 0, 0, 2, 0, 0, 5, 1}
	actual, err := Hints(original)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	indexes, n := Indexes(actual)
	if n != 8 {
		t.Errorf("expected 8 accelrated got %d", n)
	}
	for i := 0; i < len(original); i++ {
		if actual[i] != expected[i] {
			t.Errorf("expected %q, got %q", expected[i], actual[i])
		}
	}
	for i := 0; i < len(original); i++ {
		if indexes[i] != expectedIndexes[i] {
			t.Errorf("expected %d, got %d", expectedIndexes[i], indexes[i])
		}
	}
}
