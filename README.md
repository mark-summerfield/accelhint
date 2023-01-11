# accelhints

A Go library for setting keyboard accelerators in a list of strings.

For example:

	editMenuStrings := []string{
		"Undo",
		"Redo",
		"Copy",
		"Cu&t", // preset
		"Paste",
		"Find",
		"Find Again",
	}
	err := AddHints(editMenuStrings) // changes in-place

	// editMenuStrings is now:

	[]string{
	"&Undo",
	"&Redo",
	"&Copy",
	"Cu&t",
	"&Paste",
	"&Find",
	"Find &Again"}

If you don't want in-place use `slices.Clone()`.

Use `AddHintsFull` to control the marker and alphabet.

## License

Apache-2.0

---
