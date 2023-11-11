package main

import (
	"fmt"
	"io"
	"os"
	"sort"
)

var (
	stdin  io.Reader = os.Stdin
	stdout io.Writer = os.Stdout
	stderr io.Writer = os.Stderr
	exit             = os.Exit
)

func main() {
	// stats is a map of strings in "FirstLetterLastLetter" format.
	// The inner map contains the sorted letters between the first and last letter.
	stats := make(map[string]map[string][]string)

	var word string
	for _, err := fmt.Fscanf(stdin, "%s", &word); err != io.EOF; _, err = fmt.Fscanf(stdin, "%s", &word) {
		if err != nil || len(word) < 3 { // word is too small to consider
			continue
		}

		firstKey := fmt.Sprintf("%s%s", word[:1], word[len(word)-1:])
		if _, ok := stats[firstKey]; !ok {
			stats[firstKey] = make(map[string][]string)
		}

		middle := word[1 : len(word)-1]
		runes := []rune(middle)
		sort.Slice(runes, func(i, j int) bool {
			return runes[i] < runes[j]
		})
		secondKey := string(runes)

		stats[firstKey][secondKey] = append(stats[firstKey][secondKey], word)
	}

	fmt.Fprintln(stdout, "Unique first/last pairs", len(stats))

	mostOverlappingWords := 0
	var words [][]string

	for _, inner := range stats {
		for _, wordz := range inner {
			if length := len(wordz); length > mostOverlappingWords {
				mostOverlappingWords = length
				words = append([][]string{}, wordz)
			} else if length == mostOverlappingWords {
				words = append(words, wordz)
			}
		}
	}

	fmt.Fprintln(stdout, mostOverlappingWords)
	fmt.Fprintln(stdout, len(words))
	for _, group := range words {
		fmt.Fprintln(stdout, group)
	}
}
