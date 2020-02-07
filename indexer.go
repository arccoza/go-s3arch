package main

import (
	"io"
	"bufio"
	"bytes"

	// "github.com/davecgh/go-spew/spew"
	"github.com/blevesearch/segment"
	"github.com/arccoza/go-s3arch/pkg/stopwords"
)

func TokenizeText(text io.Reader, types [5]bool) *bufio.Scanner {
	scanner := bufio.NewScanner(text)
	scanner.Split(SplitterGen(types))
	return scanner
}

func SplitterGen(skip [5]bool) func([]byte, bool) (int, []byte, error) {
	return func(data []byte, atEOF bool) (int, []byte, error) {
		advAll := 0
		Again:
			adv, tok, typ, err := segment.SegmentWords(data, atEOF)
			advAll += adv
		
		copy(tok[:adv], bytes.ToLower(tok[:adv]))
		if err == nil && tok != nil && adv > 0 && (skip[typ] || stopwords.Match(string(tok[:adv]))) {
			if atEOF {
				data = data[adv:]
				goto Again
			}
			tok = nil
		}

		return advAll, tok, err
	}
}