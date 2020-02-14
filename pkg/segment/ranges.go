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
