// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: Apache-2.0

package accelhints

import (
	_ "embed"
	"regexp"
	"strings"
	"unicode"

	"github.com/charles-haynes/munkres"
	"golang.org/x/exp/slices"
)

//go:embed Version.dat
var Version string

type weights [][]float64

const (
	Alphabet  = "ABCDEFGHIJKLMNOPQRSTUVWXYZ123456789" // MUST be UPPERCASE
	Marker    = '&'
	GtkMarker = '_'
	maxWeight = 90000.0
)

// Updates items inserting '&'s to indicate accelerators. Only characters in
// the Alphabet are candidates. Returns the number of accelerated items.
func AddHints(items []string) (int, error) {
	return AddHintsFull(items, Marker, Alphabet)
}

// Updates items inserting marker's (only ASCII allowed) to indicate
// accelerators with characters from the given alphabet (of unique uppercase
// characters) as candidates. Returns the number of accelerated items.
func AddHintsFull(items []string, marker byte, alphabet string) (int, error) {
	normalizeMarker(items, marker)
	chars := []rune(alphabet)
	weights := getWeights(items, marker, chars)
	m, err := munkres.NewHungarianAlgorithm(weights)
	if err != nil {
		return 0, err
	}
	indexes := m.Execute()
	done := updateItems(items, marker, chars, indexes)
	return done, nil
}

func normalizeMarker(items []string, marker byte) {
	rx := regexp.MustCompile(`[_&]`)
	m := string(marker)
	for i := 0; i < len(items); i++ {
		items[i] = rx.ReplaceAllLiteralString(items[i], m)
	}
}

func getWeights(items []string, marker byte, chars []rune) weights {
	weights := makeMaxWeights(len(chars))
	updateWeights(items, weights, rune(marker), chars)
	return weights
}

func makeMaxWeights(size int) weights {
	weights := make(weights, 0)
	for row := 0; row < size; row++ {
		weights = append(weights, []float64{})
		for column := 0; column < size; column++ {
			weights[row] = append(weights[row], maxWeight)
		}
	}
	return weights
}

func updateWeights(items []string, weights weights, marker rune,
	chars []rune) {
	prev := rune(0)
	for row, item := range items {
		weight := 0.0
		for column, c := range item {
			c = unicode.ToUpper(c)
			i := slices.Index(chars, c)
			if i > -1 { // c in alphabet
				if column == 0 { // first
					weight = maxWeight - 4.0
				} else if prev == marker { // preset
					weight = maxWeight - 99.0
				} else if unicode.IsSpace(prev) { // word start
					weight = maxWeight - 2.0
				} else { // anywhere
					weight = maxWeight - 1.0
				}
				if weights[row][i] > weight {
					weights[row][i] = weight
				}
			}
			prev = c
		}
	}
}

func updateItems(items []string, marker byte, chars []rune,
	indexes []int) int {
	done := 0
	m := string(marker)
	for row, column := range indexes {
		if column == -1 {
			continue // unassigned
		}
		c := chars[column]
		if row < len(items) {
			item := items[row]
			i := strings.IndexByte(item, marker)
			if i > -1 {
				done++
				continue // user preset
			}
			uitem := strings.ToUpper(item)
			if len(uitem) == 0 {
				continue // skip
			}
			var index int // first is best
			sc := string(c)
			if !strings.HasPrefix(uitem, sc) { // !first
				index = strings.Index(uitem, " "+sc)
				if index > -1 { // start of word is second best
					index++ // skip the space
				} else { // anywhere is third best
					index = strings.IndexRune(uitem, c)
				}
			}
			if index > -1 {
				items[row] = item[:index] + m + item[index:]
				done++
			}

		}
	}
	return done
}
