//go:generate go run tests_gen.go

package segment

import (
	"bufio"
	"fmt"
	"io"
	"unicode"
	"unicode/utf8"

	"github.com/arccoza/go-s3arch/pkg/hangul"
	"github.com/davecgh/go-spew/spew"
)

func Graphemes(text io.Reader) *bufio.Scanner {
	scanner := bufio.NewScanner(text)
	scanner.Split(graphemeSplit)
	return scanner
}

func graphemeSplit(data []byte, atEOF bool) (int, []byte, error) {
	if len(data) == 0 {
		return 0, nil, nil
	}

	const (
		Char = hangul.LVT << (iota + 1)
		Extend
		Join
		None = 0
	)

	prev := None
	ZWJ := '\u200D'
	adv, _, stp := 0, 0, 0
	// lth := adv - skp
	buf := data[:]
	fmt.Println("-------------------", atEOF)

	for i, take := 0, 2; i < take && len(buf) > 0; i++ {
		r, s := utf8.DecodeRune(buf)
		stp = s
		adv += s
		buf = buf[stp:]

		if r == ZWJ {
			spew.Dump("ZWJ")
			take += 2
			prev = Join
		} else if unicode.Is(unicode.Regional_Indicator, r) {
			spew.Dump("GB12, GB13", len(buf))
			take += 1
			stp = 0
			prev = Extend
		} else if unicode.In(r, unicode.Mc,
			unicode.Prepended_Concatenation_Mark,
			Consonant_Preceding_Repha,
			Consonant_Prefixed) {
			spew.Dump("GB9a, GB9b", len(buf))
			take += 1
			stp = 0
			prev = Extend
		} else if unicode.Is(unicode.Extender, r) {
			spew.Dump("Extender", len(buf))
			take += 1
			stp = 0
			prev = Extend
		} else if typ := hangul.SyllableType(r); typ > 0 {
			spew.Dump("GB6/7/8 Hangul")

			if prev > 0 &&
				((prev == hangul.L && (typ&hangul.L_V_LV_LVT > 0)) ||
					(prev&hangul.LV_V > 0 && (typ&hangul.V_T > 0)) ||
					(prev&hangul.LVT_T > 0 && typ == hangul.T)) {
				take += 1
				stp = 0
			}

			prev = typ
		} else if prev == Char {
			break
		} else {
			spew.Dump("Char")
			stp = 0
			prev = Char
		}
	}

	adv -= stp
	return adv, data[:adv], nil
}
