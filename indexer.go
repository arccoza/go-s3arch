package main

import (
	// "github.com/bbalet/stopwords"
	"github.com/davecgh/go-spew/spew"
	"bufio"
	"strings"
	// "fmt"
	"github.com/blevesearch/segment"
)

func StemText(text string) string {
	// strings.NewReader(text)

	scanner := bufio.NewScanner(strings.NewReader(text))
	scanner.Split(SplitterGen([5]bool{true, false, false, false, false}))
	for scanner.Scan() {
		tokenBytes := scanner.Text()
		spew.Dump(tokenBytes, len(tokenBytes))
	}
	if err := scanner.Err(); err != nil {
		spew.Dump("err:", err)
	}

	return ""
}

func SplitterGen(skip [5]bool) func([]byte, bool) (int, []byte, error) {
	return func(data []byte, atEOF bool) (int, []byte, error) {
		advance := 0
		Again:
			adv, tok, typ, err := segment.SegmentWords(data, atEOF)
			advance += adv
		
		if skip[typ] && err == nil && tok != nil && adv > 0 {
			if atEOF {
				data = data[adv:]
				goto Again
			}
			tok = nil
		}

		return advance, tok, err
	}
}