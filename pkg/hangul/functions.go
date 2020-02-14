package hangul

import (
  "unicode"
)

const (
	TypeNone = iota
    TypeL
    TypeV
    TypeT
    TypeLV
    TypeLVT
)

func isL(r rune) bool {
	return 0x1100 <= r && r <= 0xA97C && unicode.Is(Leading_Jamo, r)
}

func isV(r rune) bool {
	return 0x1160 <= r && r <= 0xD7C6 && unicode.Is(Vowel_Jamo, r)
}

func isT(r rune) bool {
	return 0x11A8 <= r && r <= 0xD7FB && unicode.Is(Trailing_Jamo, r)
}

func isLV(r rune) bool {
	return unicode.Is(LV_Syllable, r)
}

func isLVT(r rune) bool {
	return 0xAC01 <= r && r <= 0xD7A3 && unicode.Is(LVT_Syllable, r)
}

func SyllableType(r rune) int {
	switch {
	case isL(r):
		return TypeL
	case isV(r):
		return TypeV
	case isT(r):
		return TypeT
	case isLV(r):
		return TypeLV
	case isLVT(r):
		return TypeLVT
	}

	return TypeNone
}
