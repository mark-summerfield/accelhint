# accelhint

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
		"Find && Replace", // literal '&'
	}
	hinted, err := accelhint.Hints(editMenuStrings) 
	// hinted is:
	[]string{
		"&Undo",
		"&Redo",
		"&Copy",
		"Cu&t",
		"&Paste",
		"&Find",
		"Find &Again",
		"F&ind && Replace"
	}

Use `HintsFull` to control the marker and alphabet.
Use `Indexes` or `IndexesFull` to get the index positions where the
accelerators should go.

## License

Apache-2.0

---
