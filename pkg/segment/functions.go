//go:generate go run tests_gen.go

package segment

import (
	"bufio"
	"fmt"
	"io"
	"unicode"
	"unicode/utf8"
	"math/bits"

	"github.com/arccoza/go-s3arch/pkg/hangul"
	"github.com/davecgh/go-spew/spew"
)

func Graphemes(text io.Reader) *bufio.Scanner {
	scanner := bufio.NewScanner(text)
	scanner.Split(graphemeSplit)
	return scanner
}

const (
	// Char = hangul.LVT << (iota + 1)
	CR = 1 << (iota + 2)
	// CR = (iota + 2)
	LF
	Control
	Extend
	RI
	Prepend
	SpacingMark
	L
	V
	T
	LV
	LVT
	ExtPict
	Extend_ExtCccZwj
	ZWJ_ExtCccZwj
	Other = 1
	None = 0
)

const (
	lf = '\x0A'
	cr = '\x0D'
	zwj = '\u200D'
	zwnj = '\u200C'
)

func graphemeSplit(data []byte, atEOF bool) (int, []byte, error) {
	if len(data) == 0 {
		return 0, nil, nil
	}

	prev := None
	adv, _, stp := 0, 0, 0
	// lth := adv - skp
	buf := data[:]
	idx := 0
	fmt.Println("-------------------", atEOF)
	// fmt.Println(SpacingMark, hangul.L, L, hangul.LVT, LVT, ExtPict)

	for i, take := 0, 1; i < take && len(buf) > 0; i++ {
		r, s := utf8.DecodeRune(buf)
		stp = s
		adv += s
		buf = buf[stp:]

		if prev > 1 {
			idx = bits.TrailingZeros(uint(prev))
		} else {
			idx = prev
		}
		fmt.Println(prev, idx)
		if prev = firstOf(r, breakTable[idx]...); prev == -1 {
			fmt.Println("break")
			// stp = 0
			break
		} else {
			stp = 0
			// spew.Dump(gr)
			take++
		}

		// fmt.Println(firstOf(r, breakTable[prev]...))
	}

	adv -= stp
	spew.Dump(string(data[:adv]))
	return adv, data[:adv], nil
}

var breakTable = [][]func(rune) int{
	// Unknown
	{isCR, isLF, isControl, isExtend, isRI, isPrepend, isSpacingMark, isExtend_ExtCccZwj, isZWJ_ExtCccZwj, isOther,},
	// Other
	{isExtend, isSpacingMark, isExtend_ExtCccZwj, isZWJ_ExtCccZwj,},
	// CR
	{isLF,},
	// LF
	{},
	// Control
	{},
	// Extend
	{isExtend, isSpacingMark, isExtend_ExtCccZwj, isZWJ_ExtCccZwj,},
	// RI
	{isExtend, isRI, isSpacingMark, isExtend_ExtCccZwj, isZWJ_ExtCccZwj,},
	// Prepend
	// {},
	{isExtend, isRI, isPrepend, isSpacingMark, hangul.IsL, hangul.IsV, hangul.IsT, hangul.IsLV, hangul.IsLVT, isExtPict, isExtend_ExtCccZwj, isZWJ_ExtCccZwj, isOther,},
	// SpacingMark
	{isExtend, isSpacingMark, isExtend_ExtCccZwj, isZWJ_ExtCccZwj,},
	// L
	{hangul.IsL, hangul.IsV, hangul.IsLV, hangul.IsLVT, isExtend, isSpacingMark,},
	// V
	{hangul.IsV, hangul.IsT, isExtend, isSpacingMark,},
	// T
	{hangul.IsT, isExtend, isSpacingMark,},
	// LV
	{hangul.IsV, hangul.IsT, isExtend, isSpacingMark,},
	// LVT
	{hangul.IsT, isExtend, isSpacingMark,},
}

func firstOf(r rune, fns ...func(rune) int) int {
	for _, fn := range fns {
		if v := fn(r); v > 0 {
			return v
		}
	}
	return -1
}

func b2i(b bool) int {
	if b {
		return 1
	}
	return 0
}

func isOther(r rune) int {
	fmt.Println("isOther")
	return Other
}

func isCR(r rune) int {
	fmt.Println("isCR")
	return b2i(r == cr) * CR
}

func isLF(r rune) int {
	fmt.Println("isLF")
	return b2i(r == lf) * LF
}

func isControl(r rune) int {
	if r == cr || r == zwj || r == zwnj {
		return 0
	}
	return b2i(unicode.IsControl(r)) * Control
}

func isExtend(r rune) int {
	return b2i(unicode.In(r, unicode.Diacritic, unicode.Extender)) * Extend
}

func isRI(r rune) int {
	fmt.Println("isRI")
	return b2i(unicode.Is(unicode.Regional_Indicator, r)) * RI
}

func isPrepend(r rune) int {
	if unicode.In(r,
	unicode.Prepended_Concatenation_Mark,
	Consonant_Preceding_Repha,
	Consonant_Prefixed) {
		return Prepend
	}
	return 0
}

func isSpacingMark(r rune) int {
	if unicode.Is(Grapheme_SpacingMarkExtras, r) {
		return SpacingMark
	} else if unicode.Is(unicode.Mc, r) && !unicode.Is(Grapheme_SpacingMarkExceptions, r) {
		return SpacingMark
	}
	return 0
}

func isExtPict(r rune) int {
	return 0
}
func isExtend_ExtCccZwj(r rune) int {
	return 0
}
func isZWJ_ExtCccZwj(r rune) int {
	return 0
}
