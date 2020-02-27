package hangul

import (
	"unicode"
)

const (
	L = 1 << 8 << iota
	V
	T
	LV
	LVT
	L_V_LV_LVT = L | V | LV | LVT
	LV_V       = LV | V
	V_T        = V | T
	LVT_T      = LVT | T
	None       = 0
)

func b2i(b bool) int {
	if b {
		return 1
	}
	return 0
}

func IsL(r rune) int {
	return b2i(0x1100 <= r && r <= 0xA97C && unicode.Is(Leading_Jamo, r)) * L
}

func IsV(r rune) int {
	return b2i(0x1160 <= r && r <= 0xD7C6 && unicode.Is(Vowel_Jamo, r)) * V
}

func IsT(r rune) int {
	return b2i(0x11A8 <= r && r <= 0xD7FB && unicode.Is(Trailing_Jamo, r)) * T
}

func IsLV(r rune) int {
	return b2i(unicode.Is(LV_Syllable, r)) * LV
}

func IsLVT(r rune) int {
	return b2i(0xAC01 <= r && r <= 0xD7A3 && unicode.Is(LVT_Syllable, r)) * LVT
}

// func SyllableType(r rune) int {
// 	switch {
// 	case isL(r):
// 		return L
// 	case isV(r):
// 		return V
// 	case isT(r):
// 		return T
// 	case isLV(r):
// 		return LV
// 	case isLVT(r):
// 		return LVT
// 	}

// 	return None
// }
