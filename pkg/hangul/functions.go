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

func SyllableType(r rune) int {
	switch {
	case unicode.Is(Syllable_L, r):
		return TypeL
	case unicode.Is(Syllable_V, r):
		return TypeV
	case unicode.Is(Syllable_T, r):
		return TypeT
	case unicode.Is(Syllable_LV, r):
		return TypeLV
	case unicode.Is(Syllable_LVT, r):
		return TypeLVT
	}

	return TypeNone
}