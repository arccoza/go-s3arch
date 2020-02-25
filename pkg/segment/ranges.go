package segment

import (
	"unicode"
)

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

var Grapheme_SpacingMarkExtras = &unicode.RangeTable{
	R16: []unicode.Range16{
		{0x0E33, 0x0EB3, 128},
	},
}

var Grapheme_SpacingMarkExceptions = &unicode.RangeTable{
	R16: []unicode.Range16{
		{0x102B, 0x102C, 1},
		{0x1038, 0x1038, 1},
		{0x1062, 0x1064, 1},
		{0x1067, 0x106D, 1},
		{0x1083, 0x1083, 1},
		{0x1087, 0x108C, 1},
		{0x108F, 0x108F, 1},
		{0x109A, 0x109C, 1},
		{0x1A61, 0x1A61, 1},
		{0x1A63, 0x1A64, 1},
		{0xAA7B, 0xAA7D, 2},
	},
	R32: []unicode.Range32{
		{0x11720, 0x11721, 1},
	},
}
