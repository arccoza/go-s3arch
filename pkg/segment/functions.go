package segment

import (
	"io"
	"fmt"
	"bufio"
	"unicode"
	"unicode/utf8"

	"github.com/davecgh/go-spew/spew"
	"github.com/arccoza/go-s3arch/pkg/hangul"
)

func Graphemes(text io.Reader) *bufio.Scanner {
	scanner := bufio.NewScanner(text)
	scanner.Split(graphemeSplit)
	return scanner
}

func graphemeSplit(data []byte, atEOF bool) (int, []byte, error) {
	if (len(data) == 0) {
		return 0, nil, nil
	}

	var Consonant_Prefixed = &unicode.RangeTable{
	  R32: []unicode.Range32{
	    {0x111C2, 0x111C3, 1},
	    {0x11A3A, 0x11A3A, 1},
	    {0x11A84, 0x11A89, 1},
	  },
	}

	var Consonant_Preceding_Repha = &unicode.RangeTable{
	  R16: []unicode.Range16{
	    {0x0D4E, 0x0D4E, 1},
	  },
	  R32: []unicode.Range32{
	    {0x11D46, 0x11D46, 1},
	  },
	}

	const (
		Char = hangul.L << iota
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
		} else if unicode.In(r, unicode.Mc, unicode.Prepended_Concatenation_Mark, Consonant_Preceding_Repha, Consonant_Prefixed) {
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
			((prev == hangul.L && (typ & hangul.L_V_LV_LVT > 0)) ||
			(prev & hangul.LV_V > 0 && (typ & hangul.V_T > 0)) ||
			(prev & hangul.LVT_T > 0 && typ == hangul.T)) {
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
