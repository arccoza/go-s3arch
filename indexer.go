package main

import (
	// "github.com/go-ego/gse"
	// "github.com/bbalet/stopwords"
	"github.com/davecgh/go-spew/spew"
	"bufio"
	"strings"
	"github.com/blevesearch/segment"
)

func StemText(text string) string {
	// var seg gse.Segmenter

	// seg.LoadDict()
	// words := seg.Segment([]byte(text))
	// // words := gse.SplitTextToWords([]byte(text))
	// count:= len(words)
	// spew.Dump(text, words)

	strings.NewReader(text)

	scanner := bufio.NewScanner(strings.NewReader(text))
	scanner.Split(segment.SplitWords)
	count := 0
	for scanner.Scan() {
		count++
		tokenBytes := scanner.Bytes()
		spew.Dump(tokenBytes)
	}
	if err := scanner.Err(); err != nil {
		spew.Dump(err)
	}

	spew.Dump(count)

	return ""
}

func StemWord(word string) string {
	return ""
}