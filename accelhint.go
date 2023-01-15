// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: Apache-2.0

package accelhint

import (
	_ "embed"
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
	placeholder = "\a\a"
)

// Returns items with '&'s to indicate accelerators, and the number
// accelerated. Only characters in the Alphabet are candidates. Use '&&' for
// literal '&'s.
func Hints(items []string) ([]string, int, error) {
	return HintsX(items, Marker, Alphabet)
}

// Returns items with marker's (only ASCII allowed) to indicate
// accelerators with characters from the given alphabet (of unique uppercase
// characters) as candidates, and how many were accelerated. Use marker +
// marker for literal markers.
func HintsX(items []string, marker byte, alphabet string) ([]string,
	int, error) {
	lines := normalized(items, marker)
	alphabetChars := []rune(alphabet)
	weights := getWeights(lines, marker, alphabetChars)
	m, err := munkres.NewHungarianAlgorithm(weights)
	if err != nil {
		return nil, 0, err
	}
	indexes := m.Execute()
	lines, count := applyIndexes(items, marker, alphabetChars, indexes)
	return lines, count, nil
}

// Returns the accelerated chars from the hinted strings assuming '&' in the
// accelerator marker. rune(0) indicates no accelerator.
func Accelerators(hinted []string) []rune {
	return AcceleratorsX(hinted, '&')
}

// Returns the accelerated chars from the hinted strings using the given
// accelerator marker. rune(0) indicates no accelerator.
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

func getWeights(items []string, marker byte, alphabet []rune) weights {
	weights := makeMaxWeights(len(alphabet))
	updateWeights(items, weights, rune(marker), alphabet)
	return weights
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
	alphabet []rune) {
	m := string(marker)
	mm := m + m
	prev := rune(0)
	for row, item := range items {
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
}

func applyIndexes(items []string, marker byte, alphabet []rune,
	indexes []int) ([]string, int) {
	lines := make([]string, 0, len(items))
	count := 0
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
		i := strings.IndexByte(uline, marker)
		if i > -1 {
			lines = append(lines, line)
			count++
			continue // user preset
		}
		c := alphabet[column]
		sc := string(c)
		var index int                      // first is best
		if !strings.HasPrefix(uline, sc) { // !first
			index = strings.Index(uline, " "+sc)
			if index > -1 { // start of word is second best
				index++ // skip the space
			} else { // anywhere is third best
				index = strings.IndexRune(uline, c)
			}
		}
		if index > -1 {
			line = line[:index] + m + line[index:]
			count++
		}
		lines = append(lines, line)
	}
	return lines, count
}
