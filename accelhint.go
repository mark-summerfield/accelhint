// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: Apache-2.0

package accelhint

import (
	_ "embed"
	"fmt"
	"strings"
	"unicode"

	"github.com/charles-haynes/munkres"
	"golang.org/x/exp/slices"
)

//go:embed Version.dat
var Version string

type weights [][]float64

const (
	Alphabet    = "ABCDEFGHIJKLMNOPQRSTUVWXYZ123456789" // MUST be UPPERCASE
	Marker      = '&'
	GtkMarker   = '_'
	maxWeight   = 900100.0
	placeholder = "||"
)

// Returns items with '&'s to indicate accelerators, and the number
// accelerated. Only characters in the Alphabet are candidates. Use '&&' for
// literal '&'s.
// See also HintedX.
func Hinted(items []string) ([]string, int, error) {
	return HintedX(items, Marker, Alphabet)
}

// Returns items with marker's (only ASCII allowed) to indicate
// accelerators with characters from the given alphabet (of unique uppercase
// characters) as candidates, and how many were accelerated. Use marker +
// marker for literal markers.
// See also Hinted.
func HintedX(items []string, marker byte, alphabet string) ([]string,
	int, error) {
	lines := normalized(items, marker)
	alphabetChars := []rune(alphabet)
	weights, err := getWeights(lines, marker, alphabetChars)
	if err != nil {
		return nil, 0, err
	}
	m, err := munkres.NewHungarianAlgorithm(weights)
	if err != nil {
		return nil, 0, err
	}
	indexes := m.Execute()
	lines, count, err := applyIndexes(items, marker, alphabetChars, indexes)
	return lines, count, err
}

// Returns the accelerated chars from the hinted strings assuming '&' is the
// accelerator marker. rune(0) indicates no accelerator.
// See also AcceleratorsX.
func Accelerators(hinted []string) []rune {
	return AcceleratorsX(hinted, '&')
}

// Returns the accelerated chars from the hinted strings using the given
// accelerator marker. rune(0) indicates no accelerator.
// See also Accelerators.
func AcceleratorsX(hinted []string, marker byte) []rune {
	m := string(marker)
	mm := m + m
	chars := make([]rune, 0, len(hinted))
	for _, hint := range hinted {
		hint = strings.ReplaceAll(hint, mm, placeholder)
		hintChars := []rune(hint)
		i := slices.Index(hintChars, rune(marker))
		if i == -1 || i+1 == len(hintChars) {
			chars = append(chars, 0)
		} else {
			chars = append(chars, hintChars[i+1])
		}
	}
	return chars
}

func normalized(items []string, marker byte) []string {
	lines := make([]string, 0, len(items))
	m := string(marker)
	mm := m + m
	for _, line := range items {
		lines = append(lines, strings.ReplaceAll(line, mm, placeholder))
	}
	return lines
}

func getWeights(items []string, marker byte, alphabet []rune) (weights,
	error) {
	weights := makeMaxWeights(len(alphabet))
	err := updateWeights(items, weights, rune(marker), alphabet)
	return weights, err
}

func makeMaxWeights(size int) weights {
	weights := make(weights, 0, size)
	for row := 0; row < size; row++ {
		weights = append(weights, make([]float64, 0, size))
		for column := 0; column < size; column++ {
			weights[row] = append(weights[row], maxWeight)
		}
	}
	return weights
}

func updateWeights(items []string, weights weights, marker rune,
	alphabet []rune) error {
	m := string(marker)
	mm := m + m
	prev := rune(0)
	for row, item := range items {
		if row == len(weights) {
			break
		}
		weight := 0.0
		item = strings.ToUpper(strings.ReplaceAll(item, mm, placeholder))
		for column, c := range item {
			i := slices.Index(alphabet, c)
			if i > -1 { // c in alphabet
				if prev == marker { // preset
					weight = maxWeight - 99.0
				} else if column == 0 { // first
					weight = maxWeight - 4.0
				} else if unicode.IsSpace(prev) { // word start
					weight = maxWeight - 2.0
				} else { // anywhere
					weight = maxWeight - 1.0
				}
				// slightly prefer earlier column & later row
				weight += (float64(column) / 1100.0) - (float64(row) /
					1000.0)
				if weights[row][i] > weight {
					weights[row][i] = weight
				}
			}
			prev = c
		}
	}
	return nil
}

func applyIndexes(items []string, marker byte, alphabet []rune,
	indexes []int) ([]string, int, error) {
	const errTemplate = "duplicate accelerator %q in rows %d and %d"
	seen := make(map[rune]int) // key=char value=row in items
	lines := make([]string, 0, len(items))
	m := string(marker)
	mm := m + m
	for row, column := range indexes {
		if row == len(items) {
			break
		}
		line := items[row]
		if column == -1 || len(line) == 0 {
			lines = append(lines, line)
			continue // unassigned or empty
		}
		uline := strings.ReplaceAll(strings.ToUpper(line), mm, placeholder)
		chars := []rune(uline)
		i := slices.Index(chars, rune(marker))
		if i > -1 && i+1 < len(chars) {
			c := chars[i+1]
			if firstRow, found := seen[c]; found {
				return nil, 0, fmt.Errorf(errTemplate, c, firstRow, row)
			}
			seen[c] = row
			lines = append(lines, line)
			continue // user preset
		}
		c := alphabet[column]
		sc := string(c)
		var index int      // first is best
		if chars[0] != c { // !first
			index = strings.Index(uline, " "+sc)
			if index > -1 { // start of word is second best
				index++ // skip the space
			} else { // anywhere is third best
				index = slices.Index(chars, c)
			}
		}
		if index > -1 {
			if firstRow, found := seen[c]; found {
				return nil, 0, fmt.Errorf(errTemplate, c, firstRow, row)
			}
			seen[c] = row
			line = line[:index] + m + line[index:]
		}
		lines = append(lines, line)
	}
	for row := len(items) - (len(items) - len(indexes)); row < len(items); row++ {
		lines = append(lines, items[row])
	}

	return lines, len(seen), nil
}
