// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: Apache-2.0

package accelhint

import (
	_ "embed"
	"fmt"
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
	Alphabet    = "ABCDEFGHIJKLMNOPQRSTUVWXYZ123456789" // MUST be UPPERCASE
	Marker      = '&'
	GtkMarker   = '_'
	maxWeight   = 90000.0
	placeholder = "\a\a"
)

// Returns items with '&'s to indicate accelerators. Only characters in
// the Alphabet are candidates. Use '&&' for literal '&'s.
func Hints(items []string) ([]string, error) {
	return HintsFull(items, Marker, Alphabet)
}

// Returns items with marker's (only ASCII allowed) to indicate
// accelerators with characters from the given alphabet (of unique uppercase
// characters) as candidates. Use marker + marker for literal markers.
func HintsFull(items []string, marker byte, alphabet string) ([]string,
	error) {
	lines := normalized(items, marker)
	chars := []rune(alphabet)
	weights := getWeights(lines, marker, chars)
	m, err := munkres.NewHungarianAlgorithm(weights)
	if err != nil {
		return nil, err
	}
	indexes := m.Execute()
	lines = applyIndexes(items, marker, chars, indexes)
	return lines, nil
}

// Returns a list of accelerator indexes in the given hinted strings, and
// how many have indexes (i.e., a count of those with an index > -1).
// For those with no accelerator their index value is -1.
// Assumes the marker is '&'.
func Indexes(hinted []string) ([]int, int) {
	return IndexesFull(hinted, '&')
}

// Returns a list of accelerator indexes in the given hinted strings, and
// how many have indexes (i.e., a count of those with an index > -1).
// For those with no accelerator their index value is -1.
func IndexesFull(hinted []string, marker byte) ([]int, int) {
	indexes := make([]int, 0, len(hinted))
	m := string(marker)
	mm := m + m
	count := 0
	for _, hint := range hinted {
		hint = strings.ReplaceAll(hint, mm, placeholder)
		i := strings.IndexByte(hint, marker)
		indexes = append(indexes, i)
		if i > -1 {
			count++
		}
	}
	return indexes, count
}

func normalized(items []string, marker byte) []string {
	lines := make([]string, 0, len(items))
	rx := regexp.MustCompile(`[_&]`)
	m := string(marker)
	mm := m + m
	for _, line := range items {
		lines = append(lines, rx.ReplaceAllLiteralString(
			strings.ReplaceAll(line, mm, placeholder), m))
	}
	return lines
}

func getWeights(items []string, marker byte, chars []rune) weights {
	weights := makeMaxWeights(len(chars))
	updateWeights(items, weights, rune(marker), chars)
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
	chars []rune) {
	prev := rune(0)
	for row, item := range items {
		weight := 0.0
		item = strings.ToUpper(strings.ReplaceAll(item, "&&", placeholder))
		for column, c := range item {
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
				weight += (float64(column) / 100.0) // mildly prefer earlier
				if weights[row][i] > weight {
					weights[row][i] = weight
				}
			}
			prev = c
		}
	}
}

func applyIndexes(items []string, marker byte, chars []rune,
	indexes []int) []string {
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
		i := strings.IndexByte(uline, marker)
		if i > -1 {
			lines = append(lines, line)
			continue // user preset
		}
		c := chars[column]
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
		}
		lines = append(lines, line)
	}
	return lines
}

func sanityCheck(items, hinted []string) bool {
	indexes, _ := Indexes(hinted)
	used := make(map[rune]bool, len(hinted))
	for i := 0; i < len(hinted); i++ {
		index := indexes[i]
		if index > -1 {
			chars := []rune(items[i])
			c := chars[index]
			if c == '&' {
				continue
			}
			_, found := used[c]
			if found {
				fmt.Println("items:", strings.Join(items, "| "))
				fmt.Println("hints:", strings.Join(hinted, "| "))
				fmt.Println("indxs:", indexes, "dup:", string(c))
				return false // duplicate
			}
			used[c] = true
		}
	}
	return true
}
