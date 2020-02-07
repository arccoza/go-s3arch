package main

import (
	// "github.com/bbalet/stopwords"
	"github.com/davecgh/go-spew/spew"
	"bufio"
	"strings"
	"bytes"
	// "fmt"
	"github.com/blevesearch/segment"
	"github.com/kljensen/snowball"
	"github.com/arccoza/go-s3arch/pkg/stopwords"
)

func StemText(text string) string {
	// strings.NewReader(text)

	scanner := bufio.NewScanner(strings.NewReader(text))
	scanner.Split(SplitterGen([5]bool{true, false, false, false, false}))
	for scanner.Scan() {
		token := scanner.Text()
		// spew.Dump(ReStopWords.MatchString(token), token)
		spew.Dump(snowball.Stem(token, "english", true))
	}
	// spew.Dump(ReStopWords.MatchString("fall"), "fall", len(`|`))
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
		
		copy(tok[:adv], bytes.ToLower(tok[:adv]))
		if err == nil && tok != nil && adv > 0 && (skip[typ] || stopwords.Match(string(tok[:adv]))) {
			if atEOF {
				data = data[adv:]
				goto Again
			}
			tok = nil
		}

		return advance, tok, err
	}
}