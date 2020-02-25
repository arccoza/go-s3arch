//go:generate go run tests_gen.go

package segment

import (
	"bufio"
	"fmt"
	"io"
	"unicode"
	"unicode/utf8"

	"github.com/arccoza/go-s3arch/pkg/hangul"
	// "github.com/davecgh/go-spew/spew"
)

func Graphemes(text io.Reader) *bufio.Scanner {
	scanner := bufio.NewScanner(text)
	scanner.Split(graphemeSplit)
	return scanner
}

const (
	Char = hangul.LVT << (iota + 1)
	Control_CR
	Control_LF
	Control
	Extend
	RI
	Prepend
	SpacingMark
	ExtPict
	Extend_ExtCccZwj
	ZWJ_ExtCccZwj
	Other = 0
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

	// prev := None
	adv, _, stp := 0, 0, 0
	// lth := adv - skp
	buf := data[:]
	fmt.Println("-------------------", atEOF)
	gr := grStart

	for i, take := 0, 1; i < take && len(buf) > 0; i++ {
		r, s := utf8.DecodeRune(buf)
		stp = s
		adv += s
		buf = buf[stp:]

		if gr = gr(r); gr == nil {
			fmt.Println("break")
			// stp = 0
			break
		} else {
			stp = 0
			// spew.Dump(gr)
			take++
		}
	}

	adv -= stp
	return adv, data[:adv], nil
}

type grFn func(r rune) grFn
type isFn = func(r rune) bool

func ifttt(r rune, fns ...interface{}) grFn {
	fmt.Println("ifttt")
	for i := 0; i < len(fns); i += 2 {
		fmt.Println(i)
		f := fns[i].(isFn)
		if f(r) {
			return fns[i + 1].(func(rune) grFn)
		}
	}
	return nil
}

func grStart(r rune) grFn {
	fmt.Println("grStart")
	return ifttt(r,
		isCR, grCR,
		isLF, grLF,
		isControl, grControl,
		isExtend, grExtend,
		isRI, grRI,
		isPrepend, grPrepend,
		isSpacingMark, grSpacingMark,
		isExtend_ExtCccZwj, grExtend_ExtCccZwj,
		isZWJ_ExtCccZwj, grZWJ_ExtCccZwj,
		isOther, grOther)
}

func grOther(r rune) grFn {
	fmt.Println("grOther")
	return ifttt(r,
		isExtend, grExtend,
		isSpacingMark, grSpacingMark,
		isExtend_ExtCccZwj, grExtend_ExtCccZwj,
		isZWJ_ExtCccZwj, grZWJ_ExtCccZwj)
}

func grCR(r rune) grFn {
	fmt.Println("grCR")
	return ifttt(r, isLF, grLF)
}

func grLF(r rune) grFn {
	fmt.Println("grLF")
	return nil
}

func grControl(r rune) grFn {
	return nil
}

func grExtend(r rune) grFn {
	return ifttt(r,
		isExtend, grExtend,
		isSpacingMark, grSpacingMark,
		isExtend_ExtCccZwj, grExtend_ExtCccZwj,
		isZWJ_ExtCccZwj, grZWJ_ExtCccZwj)
}

func grRI(r rune) grFn {
	fmt.Println("grRI")
	return ifttt(r,
		isExtend, grExtend,
		isRI, grRI,
		isSpacingMark, grSpacingMark,
		isExtend_ExtCccZwj, grExtend_ExtCccZwj,
		isZWJ_ExtCccZwj, grZWJ_ExtCccZwj)
}
func grPrepend(r rune) grFn {
	return ifttt(r,
		isExtend, grExtend,
		isRI, grRI,
		isSpacingMark, grSpacingMark,
		isExtend_ExtCccZwj, grExtend_ExtCccZwj,
		isZWJ_ExtCccZwj, grZWJ_ExtCccZwj,
		isOther, grOther)
}
func grSpacingMark(r rune) grFn {
	fmt.Println("grSpacingMark")
	return ifttt(r,
		isExtend, grExtend,
		isSpacingMark, grSpacingMark,
		isExtend_ExtCccZwj, grExtend_ExtCccZwj,
		isZWJ_ExtCccZwj, grZWJ_ExtCccZwj)
}
func grExtPict(r rune) grFn {
	return nil
}
func grExtend_ExtCccZwj(r rune) grFn {
	return nil
}
func grZWJ_ExtCccZwj(r rune) grFn {
	return nil
}

func isOther(r rune) bool {
	return true
}

func isCR(r rune) bool {
	fmt.Println("isCR")
	return r == cr
}

func isLF(r rune) bool {
	fmt.Println("isLF")
	return r == lf
}

func isControl(r rune) bool {
	if r == cr || r == zwj || r == zwnj {
		return false
	}
	return unicode.IsControl(r)
}

func isExtend(r rune) bool {
	return unicode.In(r, unicode.Diacritic, unicode.Extender)
}

func isRI(r rune) bool {
	return unicode.Is(unicode.Regional_Indicator, r)
}

func isPrepend(r rune) bool {
	if unicode.In(r,
	unicode.Prepended_Concatenation_Mark,
	Consonant_Preceding_Repha,
	Consonant_Prefixed) {
		return true
	}
	return false
}

func isSpacingMark(r rune) bool {
	if unicode.Is(Grapheme_SpacingMarkExtras, r) {
		return true
	} else if unicode.Is(unicode.Mc, r) && !unicode.Is(Grapheme_SpacingMarkExceptions, r) {
		return true
	}
	return false
}

func isExtPict(r rune) bool {
	return false
}
func isExtend_ExtCccZwj(r rune) bool {
	return false
}
func isZWJ_ExtCccZwj(r rune) bool {
	return false
}
