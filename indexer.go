package main

import (
	"io"
	"fmt"
	"bufio"
	"bytes"
	// "strings"
	"unicode"
	"unicode/utf8"

	"github.com/davecgh/go-spew/spew"
	"github.com/blevesearch/segment"
	"github.com/arccoza/go-s3arch/pkg/stopwords"
)

func init() {
	// spew.Dump(utf8.RuneError)
	// a, _ := utf8.DecodeRune([]byte{})
	// spew.Dump(a ^ utf8.RuneError)

	// scanner := GraphemeBreak(strings.NewReader("🇦🇽"))
	// scanner := GraphemeBreak(strings.NewReader("👨‍👩‍👦"))
	scanner := GraphemeBreak(bytes.NewReader([]byte{0x63, 0x61, 0x66, 0x65, 0xCC, 0x81}))
	for scanner.Scan() {
		spew.Dump(scanner.Text())
	}

	// scanner := TokenizeText(bytes.NewReader([]byte{0x63, 0x61, 0x66, 0x65, 0xCC, 0x81}), [5]bool{true, false, false, false, false})
	// // spew.Dump(scanner.Scan())
	// for scanner.Scan() {
	// 	spew.Dump(scanner.Text())
	// 	scanner = GraphemeBreak(bytes.NewReader(scanner.Bytes()))
	// 	for scanner.Scan() {
	// 		spew.Dump(scanner.Text())
	// 	}
	// }
}

func GraphemeBreak(text io.Reader) *bufio.Scanner {
	scanner := bufio.NewScanner(text)
	scanner.Split(GraphemeSplit)
	return scanner
}

func GraphemeSplit(data []byte, atEOF bool) (int, []byte, error) {
	if (len(data) == 0) {
		return 0, nil, nil
	}

	// Again:
	// if (len(data) == 0) {
	// 	return 0, nil, nil
	// }
	// r, s := utf8.DecodeRune(data)
	// data = data[s:]
	// spew.Dump(r, ZWJ)
	// goto Again

	// return 0, nil, nil

	ZWJ := '\u200D'
	adv, _, stp := 0, 0, 0
	// lth := adv - skp
	buf := data[:]
	// spew.Dump(len(buf), atEOF)
	fmt.Println("-------------------", atEOF)
	
	for i, take, chrSeq := 0, 2, false; i < take && len(buf) > 0; i++ {
		r, s := utf8.DecodeRune(buf)
		stp = s
		adv += s
		// spew.Dump(adv)

		buf = buf[stp:]
		// spew.Dump(r)
		if r == ZWJ || unicode.Is(unicode.Extender, r) {
			spew.Dump("ZWJ")
			take += 2
			chrSeq = false
		} else if unicode.Is(unicode.Regional_Indicator, r) {
			spew.Dump("GB12, GB13", len(buf))
			take += 1
			stp = 0
			chrSeq = false
		} else if unicode.In(r, unicode.Mc, unicode.Prepended_Concatenation_Mark) {
			spew.Dump("GB9a, GB9b", len(buf))
			take += 1
			stp = 0
			chrSeq = false
		} else if chrSeq {
			break
		} else {
			spew.Dump("Char")
			stp = 0
			chrSeq = true
		}
	}
	adv -= stp
	return adv, data[:adv], nil
}

func isGBreak(r rune) bool {
	ZWJ := '\u200D'
	// unicode.In(r, unicode.Other_Grapheme_Extend, unicode.Me, unicode.Mn)
	return !(r == ZWJ || unicode.Is(unicode.Extender, r))
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