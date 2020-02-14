package main

import (
	"io"
	// "fmt"
	"bufio"
	"bytes"
	// "strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/blevesearch/segment"
	"github.com/arccoza/go-s3arch/pkg/stopwords"
	seg "github.com/arccoza/go-s3arch/pkg/segment"
)

func init() {
	// spew.Dump(utf8.RuneError)
	// a, _ := utf8.DecodeRune([]byte{})
	// spew.Dump(a ^ utf8.RuneError)

	// scanner := seg.Graphemes(strings.NewReader("🇦🇽"))
	// scanner := seg.Graphemes(strings.NewReader("👨‍👩‍👦"))
	// scanner := seg.Graphemes(bytes.NewReader([]byte{0x63, 0x61, 0x66, 0x65, 0xCC, 0x81}))
	// 0020 × 0308 ÷ 1100
	// scanner := seg.Graphemes(bytes.NewReader([]byte(string([]rune{0x0020, 0x0308, 0x1100}))))
	// 1100 × AC00
	scanner := seg.Graphemes(bytes.NewReader([]byte(string([]rune{0x1100, 0xAC00}))))
	for scanner.Scan() {
		spew.Dump(scanner.Text())
	}

	// scanner := TokenizeText(bytes.NewReader([]byte{0x63, 0x61, 0x66, 0x65, 0xCC, 0x81}), [5]bool{true, false, false, false, false})
	// // spew.Dump(scanner.Scan())
	// for scanner.Scan() {
	// 	spew.Dump(scanner.Text())
	// 	scanner = seg.Graphemes(bytes.NewReader(scanner.Bytes()))
	// 	for scanner.Scan() {
	// 		spew.Dump(scanner.Text())
	// 	}
	// }
}


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