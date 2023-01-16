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
    hinted, err := accelhint.Hinted(editMenuStrings) 
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

Use `HintedX` to control the marker and alphabet.
Use `Accelerators` or `AcceleratorsX` to get a slice of the accelerator runes.

For example, to populate a dynamically created menu, use something like this:

    items := make([]string, len(menuItems)) // assumes menuItems
    for _, menuItem := range menuItems {
        items = append(items, menuItem.Text())
    }
    hinted, err := accelhint.Hinted(items)
    if err != nil {
        log.Fatal(err)
    }
    accels := accelhint.Accelerators(hinted)
    for i := 0; i < len(menuItems); i++ {
        accel := accels[i]
        if accel != 0 { // has an accelerator
            chars := []rune(items[i])
            j := slices.Index(chars, accel)
            if j > -1 { // should always be true
                chars = slices.Insert(chars, j, '&') // or '_' for Gtk
                menuItems[i].SetText(string(chars))
            }
        }
    }

## License

Apache-2.0

---
