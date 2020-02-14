package hangul

import (
  "unicode"
)

// L
var Leading_Jamo = &unicode.RangeTable{
  R16: []unicode.Range16{
    {0x1100, 0x115F, 1},
    {0xA960, 0xA97C, 1},
  },
}

// V
var Vowel_Jamo = &unicode.RangeTable{
  R16: []unicode.Range16{
    {0x1160, 0x11A7, 1},
    {0xD7B0, 0xD7C6, 1},
  },
}

// T
var Trailing_Jamo = &unicode.RangeTable{
  R16: []unicode.Range16{
    {0x11A8, 0x11FF, 1},
    {0xD7CB, 0xD7FB, 1},
  },
}

// LV
var LV_Syllable = &unicode.RangeTable{
  R16: []unicode.Range16{
    {0xAC00, 0xD788, 28},
  },
}

// LVT
var LVT_Syllable = &unicode.RangeTable{
  R16: []unicode.Range16{
    {0xAC01, 0xAC1B, 1},
    {0xAC1D, 0xAC37, 1},
    {0xAC39, 0xAC53, 1},
    {0xAC55, 0xAC6F, 1},
    {0xAC71, 0xAC8B, 1},
    {0xAC8D, 0xACA7, 1},
    {0xACA9, 0xACC3, 1},
    {0xACC5, 0xACDF, 1},
    {0xACE1, 0xACFB, 1},
    {0xACFD, 0xAD17, 1},
    {0xAD19, 0xAD33, 1},
    {0xAD35, 0xAD4F, 1},
    {0xAD51, 0xAD6B, 1},
    {0xAD6D, 0xAD87, 1},
    {0xAD89, 0xADA3, 1},
    {0xADA5, 0xADBF, 1},
    {0xADC1, 0xADDB, 1},
    {0xADDD, 0xADF7, 1},
    {0xADF9, 0xAE13, 1},
    {0xAE15, 0xAE2F, 1},
    {0xAE31, 0xAE4B, 1},
    {0xAE4D, 0xAE67, 1},
    {0xAE69, 0xAE83, 1},
    {0xAE85, 0xAE9F, 1},
    {0xAEA1, 0xAEBB, 1},
    {0xAEBD, 0xAED7, 1},
    {0xAED9, 0xAEF3, 1},
    {0xAEF5, 0xAF0F, 1},
    {0xAF11, 0xAF2B, 1},
    {0xAF2D, 0xAF47, 1},
    {0xAF49, 0xAF63, 1},
    {0xAF65, 0xAF7F, 1},
    {0xAF81, 0xAF9B, 1},
    {0xAF9D, 0xAFB7, 1},
    {0xAFB9, 0xAFD3, 1},
    {0xAFD5, 0xAFEF, 1},
    {0xAFF1, 0xB00B, 1},
    {0xB00D, 0xB027, 1},
    {0xB029, 0xB043, 1},
    {0xB045, 0xB05F, 1},
    {0xB061, 0xB07B, 1},
    {0xB07D, 0xB097, 1},
    {0xB099, 0xB0B3, 1},
    {0xB0B5, 0xB0CF, 1},
    {0xB0D1, 0xB0EB, 1},
    {0xB0ED, 0xB107, 1},
    {0xB109, 0xB123, 1},
    {0xB125, 0xB13F, 1},
    {0xB141, 0xB15B, 1},
    {0xB15D, 0xB177, 1},
    {0xB179, 0xB193, 1},
    {0xB195, 0xB1AF, 1},
    {0xB1B1, 0xB1CB, 1},
    {0xB1CD, 0xB1E7, 1},
    {0xB1E9, 0xB203, 1},
    {0xB205, 0xB21F, 1},
    {0xB221, 0xB23B, 1},
    {0xB23D, 0xB257, 1},
    {0xB259, 0xB273, 1},
    {0xB275, 0xB28F, 1},
    {0xB291, 0xB2AB, 1},
    {0xB2AD, 0xB2C7, 1},
    {0xB2C9, 0xB2E3, 1},
    {0xB2E5, 0xB2FF, 1},
    {0xB301, 0xB31B, 1},
    {0xB31D, 0xB337, 1},
    {0xB339, 0xB353, 1},
    {0xB355, 0xB36F, 1},
    {0xB371, 0xB38B, 1},
    {0xB38D, 0xB3A7, 1},
    {0xB3A9, 0xB3C3, 1},
    {0xB3C5, 0xB3DF, 1},
    {0xB3E1, 0xB3FB, 1},
    {0xB3FD, 0xB417, 1},
    {0xB419, 0xB433, 1},
    {0xB435, 0xB44F, 1},
    {0xB451, 0xB46B, 1},
    {0xB46D, 0xB487, 1},
    {0xB489, 0xB4A3, 1},
    {0xB4A5, 0xB4BF, 1},
    {0xB4C1, 0xB4DB, 1},
    {0xB4DD, 0xB4F7, 1},
    {0xB4F9, 0xB513, 1},
    {0xB515, 0xB52F, 1},
    {0xB531, 0xB54B, 1},
    {0xB54D, 0xB567, 1},
    {0xB569, 0xB583, 1},
    {0xB585, 0xB59F, 1},
    {0xB5A1, 0xB5BB, 1},
    {0xB5BD, 0xB5D7, 1},
    {0xB5D9, 0xB5F3, 1},
    {0xB5F5, 0xB60F, 1},
    {0xB611, 0xB62B, 1},
    {0xB62D, 0xB647, 1},
    {0xB649, 0xB663, 1},
    {0xB665, 0xB67F, 1},
    {0xB681, 0xB69B, 1},
    {0xB69D, 0xB6B7, 1},
    {0xB6B9, 0xB6D3, 1},
    {0xB6D5, 0xB6EF, 1},
    {0xB6F1, 0xB70B, 1},
    {0xB70D, 0xB727, 1},
    {0xB729, 0xB743, 1},
    {0xB745, 0xB75F, 1},
    {0xB761, 0xB77B, 1},
    {0xB77D, 0xB797, 1},
    {0xB799, 0xB7B3, 1},
    {0xB7B5, 0xB7CF, 1},
    {0xB7D1, 0xB7EB, 1},
    {0xB7ED, 0xB807, 1},
    {0xB809, 0xB823, 1},
    {0xB825, 0xB83F, 1},
    {0xB841, 0xB85B, 1},
    {0xB85D, 0xB877, 1},
    {0xB879, 0xB893, 1},
    {0xB895, 0xB8AF, 1},
    {0xB8B1, 0xB8CB, 1},
    {0xB8CD, 0xB8E7, 1},
    {0xB8E9, 0xB903, 1},
    {0xB905, 0xB91F, 1},
    {0xB921, 0xB93B, 1},
    {0xB93D, 0xB957, 1},
    {0xB959, 0xB973, 1},
    {0xB975, 0xB98F, 1},
    {0xB991, 0xB9AB, 1},
    {0xB9AD, 0xB9C7, 1},
    {0xB9C9, 0xB9E3, 1},
    {0xB9E5, 0xB9FF, 1},
    {0xBA01, 0xBA1B, 1},
    {0xBA1D, 0xBA37, 1},
    {0xBA39, 0xBA53, 1},
    {0xBA55, 0xBA6F, 1},
    {0xBA71, 0xBA8B, 1},
    {0xBA8D, 0xBAA7, 1},
    {0xBAA9, 0xBAC3, 1},
    {0xBAC5, 0xBADF, 1},
    {0xBAE1, 0xBAFB, 1},
    {0xBAFD, 0xBB17, 1},
    {0xBB19, 0xBB33, 1},
    {0xBB35, 0xBB4F, 1},
    {0xBB51, 0xBB6B, 1},
    {0xBB6D, 0xBB87, 1},
    {0xBB89, 0xBBA3, 1},
    {0xBBA5, 0xBBBF, 1},
    {0xBBC1, 0xBBDB, 1},
    {0xBBDD, 0xBBF7, 1},
    {0xBBF9, 0xBC13, 1},
    {0xBC15, 0xBC2F, 1},
    {0xBC31, 0xBC4B, 1},
    {0xBC4D, 0xBC67, 1},
    {0xBC69, 0xBC83, 1},
    {0xBC85, 0xBC9F, 1},
    {0xBCA1, 0xBCBB, 1},
    {0xBCBD, 0xBCD7, 1},
    {0xBCD9, 0xBCF3, 1},
    {0xBCF5, 0xBD0F, 1},
    {0xBD11, 0xBD2B, 1},
    {0xBD2D, 0xBD47, 1},
    {0xBD49, 0xBD63, 1},
    {0xBD65, 0xBD7F, 1},
    {0xBD81, 0xBD9B, 1},
    {0xBD9D, 0xBDB7, 1},
    {0xBDB9, 0xBDD3, 1},
    {0xBDD5, 0xBDEF, 1},
    {0xBDF1, 0xBE0B, 1},
    {0xBE0D, 0xBE27, 1},
    {0xBE29, 0xBE43, 1},
    {0xBE45, 0xBE5F, 1},
    {0xBE61, 0xBE7B, 1},
    {0xBE7D, 0xBE97, 1},
    {0xBE99, 0xBEB3, 1},
    {0xBEB5, 0xBECF, 1},
    {0xBED1, 0xBEEB, 1},
    {0xBEED, 0xBF07, 1},
    {0xBF09, 0xBF23, 1},
    {0xBF25, 0xBF3F, 1},
    {0xBF41, 0xBF5B, 1},
    {0xBF5D, 0xBF77, 1},
    {0xBF79, 0xBF93, 1},
    {0xBF95, 0xBFAF, 1},
    {0xBFB1, 0xBFCB, 1},
    {0xBFCD, 0xBFE7, 1},
    {0xBFE9, 0xC003, 1},
    {0xC005, 0xC01F, 1},
    {0xC021, 0xC03B, 1},
    {0xC03D, 0xC057, 1},
    {0xC059, 0xC073, 1},
    {0xC075, 0xC08F, 1},
    {0xC091, 0xC0AB, 1},
    {0xC0AD, 0xC0C7, 1},
    {0xC0C9, 0xC0E3, 1},
    {0xC0E5, 0xC0FF, 1},
    {0xC101, 0xC11B, 1},
    {0xC11D, 0xC137, 1},
    {0xC139, 0xC153, 1},
    {0xC155, 0xC16F, 1},
    {0xC171, 0xC18B, 1},
    {0xC18D, 0xC1A7, 1},
    {0xC1A9, 0xC1C3, 1},
    {0xC1C5, 0xC1DF, 1},
    {0xC1E1, 0xC1FB, 1},
    {0xC1FD, 0xC217, 1},
    {0xC219, 0xC233, 1},
    {0xC235, 0xC24F, 1},
    {0xC251, 0xC26B, 1},
    {0xC26D, 0xC287, 1},
    {0xC289, 0xC2A3, 1},
    {0xC2A5, 0xC2BF, 1},
    {0xC2C1, 0xC2DB, 1},
    {0xC2DD, 0xC2F7, 1},
    {0xC2F9, 0xC313, 1},
    {0xC315, 0xC32F, 1},
    {0xC331, 0xC34B, 1},
    {0xC34D, 0xC367, 1},
    {0xC369, 0xC383, 1},
    {0xC385, 0xC39F, 1},
    {0xC3A1, 0xC3BB, 1},
    {0xC3BD, 0xC3D7, 1},
    {0xC3D9, 0xC3F3, 1},
    {0xC3F5, 0xC40F, 1},
    {0xC411, 0xC42B, 1},
    {0xC42D, 0xC447, 1},
    {0xC449, 0xC463, 1},
    {0xC465, 0xC47F, 1},
    {0xC481, 0xC49B, 1},
    {0xC49D, 0xC4B7, 1},
    {0xC4B9, 0xC4D3, 1},
    {0xC4D5, 0xC4EF, 1},
    {0xC4F1, 0xC50B, 1},
    {0xC50D, 0xC527, 1},
    {0xC529, 0xC543, 1},
    {0xC545, 0xC55F, 1},
    {0xC561, 0xC57B, 1},
    {0xC57D, 0xC597, 1},
    {0xC599, 0xC5B3, 1},
    {0xC5B5, 0xC5CF, 1},
    {0xC5D1, 0xC5EB, 1},
    {0xC5ED, 0xC607, 1},
    {0xC609, 0xC623, 1},
    {0xC625, 0xC63F, 1},
    {0xC641, 0xC65B, 1},
    {0xC65D, 0xC677, 1},
    {0xC679, 0xC693, 1},
    {0xC695, 0xC6AF, 1},
    {0xC6B1, 0xC6CB, 1},
    {0xC6CD, 0xC6E7, 1},
    {0xC6E9, 0xC703, 1},
    {0xC705, 0xC71F, 1},
    {0xC721, 0xC73B, 1},
    {0xC73D, 0xC757, 1},
    {0xC759, 0xC773, 1},
    {0xC775, 0xC78F, 1},
    {0xC791, 0xC7AB, 1},
    {0xC7AD, 0xC7C7, 1},
    {0xC7C9, 0xC7E3, 1},
    {0xC7E5, 0xC7FF, 1},
    {0xC801, 0xC81B, 1},
    {0xC81D, 0xC837, 1},
    {0xC839, 0xC853, 1},
    {0xC855, 0xC86F, 1},
    {0xC871, 0xC88B, 1},
    {0xC88D, 0xC8A7, 1},
    {0xC8A9, 0xC8C3, 1},
    {0xC8C5, 0xC8DF, 1},
    {0xC8E1, 0xC8FB, 1},
    {0xC8FD, 0xC917, 1},
    {0xC919, 0xC933, 1},
    {0xC935, 0xC94F, 1},
    {0xC951, 0xC96B, 1},
    {0xC96D, 0xC987, 1},
    {0xC989, 0xC9A3, 1},
    {0xC9A5, 0xC9BF, 1},
    {0xC9C1, 0xC9DB, 1},
    {0xC9DD, 0xC9F7, 1},
    {0xC9F9, 0xCA13, 1},
    {0xCA15, 0xCA2F, 1},
    {0xCA31, 0xCA4B, 1},
    {0xCA4D, 0xCA67, 1},
    {0xCA69, 0xCA83, 1},
    {0xCA85, 0xCA9F, 1},
    {0xCAA1, 0xCABB, 1},
    {0xCABD, 0xCAD7, 1},
    {0xCAD9, 0xCAF3, 1},
    {0xCAF5, 0xCB0F, 1},
    {0xCB11, 0xCB2B, 1},
    {0xCB2D, 0xCB47, 1},
    {0xCB49, 0xCB63, 1},
    {0xCB65, 0xCB7F, 1},
    {0xCB81, 0xCB9B, 1},
    {0xCB9D, 0xCBB7, 1},
    {0xCBB9, 0xCBD3, 1},
    {0xCBD5, 0xCBEF, 1},
    {0xCBF1, 0xCC0B, 1},
    {0xCC0D, 0xCC27, 1},
    {0xCC29, 0xCC43, 1},
    {0xCC45, 0xCC5F, 1},
    {0xCC61, 0xCC7B, 1},
    {0xCC7D, 0xCC97, 1},
    {0xCC99, 0xCCB3, 1},
    {0xCCB5, 0xCCCF, 1},
    {0xCCD1, 0xCCEB, 1},
    {0xCCED, 0xCD07, 1},
    {0xCD09, 0xCD23, 1},
    {0xCD25, 0xCD3F, 1},
    {0xCD41, 0xCD5B, 1},
    {0xCD5D, 0xCD77, 1},
    {0xCD79, 0xCD93, 1},
    {0xCD95, 0xCDAF, 1},
    {0xCDB1, 0xCDCB, 1},
    {0xCDCD, 0xCDE7, 1},
    {0xCDE9, 0xCE03, 1},
    {0xCE05, 0xCE1F, 1},
    {0xCE21, 0xCE3B, 1},
    {0xCE3D, 0xCE57, 1},
    {0xCE59, 0xCE73, 1},
    {0xCE75, 0xCE8F, 1},
    {0xCE91, 0xCEAB, 1},
    {0xCEAD, 0xCEC7, 1},
    {0xCEC9, 0xCEE3, 1},
    {0xCEE5, 0xCEFF, 1},
    {0xCF01, 0xCF1B, 1},
    {0xCF1D, 0xCF37, 1},
    {0xCF39, 0xCF53, 1},
    {0xCF55, 0xCF6F, 1},
    {0xCF71, 0xCF8B, 1},
    {0xCF8D, 0xCFA7, 1},
    {0xCFA9, 0xCFC3, 1},
    {0xCFC5, 0xCFDF, 1},
    {0xCFE1, 0xCFFB, 1},
    {0xCFFD, 0xD017, 1},
    {0xD019, 0xD033, 1},
    {0xD035, 0xD04F, 1},
    {0xD051, 0xD06B, 1},
    {0xD06D, 0xD087, 1},
    {0xD089, 0xD0A3, 1},
    {0xD0A5, 0xD0BF, 1},
    {0xD0C1, 0xD0DB, 1},
    {0xD0DD, 0xD0F7, 1},
    {0xD0F9, 0xD113, 1},
    {0xD115, 0xD12F, 1},
    {0xD131, 0xD14B, 1},
    {0xD14D, 0xD167, 1},
    {0xD169, 0xD183, 1},
    {0xD185, 0xD19F, 1},
    {0xD1A1, 0xD1BB, 1},
    {0xD1BD, 0xD1D7, 1},
    {0xD1D9, 0xD1F3, 1},
    {0xD1F5, 0xD20F, 1},
    {0xD211, 0xD22B, 1},
    {0xD22D, 0xD247, 1},
    {0xD249, 0xD263, 1},
    {0xD265, 0xD27F, 1},
    {0xD281, 0xD29B, 1},
    {0xD29D, 0xD2B7, 1},
    {0xD2B9, 0xD2D3, 1},
    {0xD2D5, 0xD2EF, 1},
    {0xD2F1, 0xD30B, 1},
    {0xD30D, 0xD327, 1},
    {0xD329, 0xD343, 1},
    {0xD345, 0xD35F, 1},
    {0xD361, 0xD37B, 1},
    {0xD37D, 0xD397, 1},
    {0xD399, 0xD3B3, 1},
    {0xD3B5, 0xD3CF, 1},
    {0xD3D1, 0xD3EB, 1},
    {0xD3ED, 0xD407, 1},
    {0xD409, 0xD423, 1},
    {0xD425, 0xD43F, 1},
    {0xD441, 0xD45B, 1},
    {0xD45D, 0xD477, 1},
    {0xD479, 0xD493, 1},
    {0xD495, 0xD4AF, 1},
    {0xD4B1, 0xD4CB, 1},
    {0xD4CD, 0xD4E7, 1},
    {0xD4E9, 0xD503, 1},
    {0xD505, 0xD51F, 1},
    {0xD521, 0xD53B, 1},
    {0xD53D, 0xD557, 1},
    {0xD559, 0xD573, 1},
    {0xD575, 0xD58F, 1},
    {0xD591, 0xD5AB, 1},
    {0xD5AD, 0xD5C7, 1},
    {0xD5C9, 0xD5E3, 1},
    {0xD5E5, 0xD5FF, 1},
    {0xD601, 0xD61B, 1},
    {0xD61D, 0xD637, 1},
    {0xD639, 0xD653, 1},
    {0xD655, 0xD66F, 1},
    {0xD671, 0xD68B, 1},
    {0xD68D, 0xD6A7, 1},
    {0xD6A9, 0xD6C3, 1},
    {0xD6C5, 0xD6DF, 1},
    {0xD6E1, 0xD6FB, 1},
    {0xD6FD, 0xD717, 1},
    {0xD719, 0xD733, 1},
    {0xD735, 0xD74F, 1},
    {0xD751, 0xD76B, 1},
    {0xD76D, 0xD787, 1},
    {0xD789, 0xD7A3, 1},
  },
}