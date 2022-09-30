// Code generated by script/generate.go. DO NOT EDIT.

package runewidth

var combining = table{
	{0x0300, 0x036F}, {0x0483, 0x0489}, {0x07EB, 0x07F3},
	{0x0C00, 0x0C00}, {0x0C04, 0x0C04}, {0x0D00, 0x0D01},
	{0x135D, 0x135F}, {0x1A7F, 0x1A7F}, {0x1AB0, 0x1AC0},
	{0x1B6B, 0x1B73}, {0x1DC0, 0x1DF9}, {0x1DFB, 0x1DFF},
	{0x20D0, 0x20F0}, {0x2CEF, 0x2CF1}, {0x2DE0, 0x2DFF},
	{0x3099, 0x309A}, {0xA66F, 0xA672}, {0xA674, 0xA67D},
	{0xA69E, 0xA69F}, {0xA6F0, 0xA6F1}, {0xA8E0, 0xA8F1},
	{0xFE20, 0xFE2F}, {0x101FD, 0x101FD}, {0x10376, 0x1037A},
	{0x10EAB, 0x10EAC}, {0x10F46, 0x10F50}, {0x11300, 0x11301},
	{0x1133B, 0x1133C}, {0x11366, 0x1136C}, {0x11370, 0x11374},
	{0x16AF0, 0x16AF4}, {0x1D165, 0x1D169}, {0x1D16D, 0x1D172},
	{0x1D17B, 0x1D182}, {0x1D185, 0x1D18B}, {0x1D1AA, 0x1D1AD},
	{0x1D242, 0x1D244}, {0x1E000, 0x1E006}, {0x1E008, 0x1E018},
	{0x1E01B, 0x1E021}, {0x1E023, 0x1E024}, {0x1E026, 0x1E02A},
	{0x1E8D0, 0x1E8D6},
}

var doublewidth = table{
	{0x1100, 0x115F}, {0x231A, 0x231B}, {0x2329, 0x232A},
	{0x23E9, 0x23EC}, {0x23F0, 0x23F0}, {0x23F3, 0x23F3},
	{0x25FD, 0x25FE}, {0x2614, 0x2615}, {0x2648, 0x2653},
	{0x267F, 0x267F}, {0x2693, 0x2693}, {0x26A1, 0x26A1},
	{0x26AA, 0x26AB}, {0x26BD, 0x26BE}, {0x26C4, 0x26C5},
	{0x26CE, 0x26CE}, {0x26D4, 0x26D4}, {0x26EA, 0x26EA},
	{0x26F2, 0x26F3}, {0x26F5, 0x26F5}, {0x26FA, 0x26FA},
	{0x26FD, 0x26FD}, {0x2705, 0x2705}, {0x270A, 0x270B},
	{0x2728, 0x2728}, {0x274C, 0x274C}, {0x274E, 0x274E},
	{0x2753, 0x2755}, {0x2757, 0x2757}, {0x2795, 0x2797},
	{0x27B0, 0x27B0}, {0x27BF, 0x27BF}, {0x2B1B, 0x2B1C},
	{0x2B50, 0x2B50}, {0x2B55, 0x2B55}, {0x2E80, 0x2E99},
	{0x2E9B, 0x2EF3}, {0x2F00, 0x2FD5}, {0x2FF0, 0x2FFB},
	{0x3000, 0x303E}, {0x3041, 0x3096}, {0x3099, 0x30FF},
	{0x3105, 0x312F}, {0x3131, 0x318E}, {0x3190, 0x31E3},
	{0x31F0, 0x321E}, {0x3220, 0x3247}, {0x3250, 0x4DBF},
	{0x4E00, 0xA48C}, {0xA490, 0xA4C6}, {0xA960, 0xA97C},
	{0xAC00, 0xD7A3}, {0xF900, 0xFAFF}, {0xFE10, 0xFE19},
	{0xFE30, 0xFE52}, {0xFE54, 0xFE66}, {0xFE68, 0xFE6B},
	{0xFF01, 0xFF60}, {0xFFE0, 0xFFE6}, {0x16FE0, 0x16FE4},
	{0x16FF0, 0x16FF1}, {0x17000, 0x187F7}, {0x18800, 0x18CD5},
	{0x18D00, 0x18D08}, {0x1B000, 0x1B11E}, {0x1B150, 0x1B152},
	{0x1B164, 0x1B167}, {0x1B170, 0x1B2FB}, {0x1F004, 0x1F004},
	{0x1F0CF, 0x1F0CF}, {0x1F18E, 0x1F18E}, {0x1F191, 0x1F19A},
	{0x1F200, 0x1F202}, {0x1F210, 0x1F23B}, {0x1F240, 0x1F248},
	{0x1F250, 0x1F251}, {0x1F260, 0x1F265}, {0x1F300, 0x1F320},
	{0x1F32D, 0x1F335}, {0x1F337, 0x1F37C}, {0x1F37E, 0x1F393},
	{0x1F3A0, 0x1F3CA}, {0x1F3CF, 0x1F3D3}, {0x1F3E0, 0x1F3F0},
	{0x1F3F4, 0x1F3F4}, {0x1F3F8, 0x1F43E}, {0x1F440, 0x1F440},
	{0x1F442, 0x1F4FC}, {0x1F4FF, 0x1F53D}, {0x1F54B, 0x1F54E},
	{0x1F550, 0x1F567}, {0x1F57A, 0x1F57A}, {0x1F595, 0x1F596},
	{0x1F5A4, 0x1F5A4}, {0x1F5FB, 0x1F64F}, {0x1F680, 0x1F6C5},
	{0x1F6CC, 0x1F6CC}, {0x1F6D0, 0x1F6D2}, {0x1F6D5, 0x1F6D7},
	{0x1F6EB, 0x1F6EC}, {0x1F6F4, 0x1F6FC}, {0x1F7E0, 0x1F7EB},
	{0x1F90C, 0x1F93A}, {0x1F93C, 0x1F945}, {0x1F947, 0x1F978},
	{0x1F97A, 0x1F9CB}, {0x1F9CD, 0x1F9FF}, {0x1FA70, 0x1FA74},
	{0x1FA78, 0x1FA7A}, {0x1FA80, 0x1FA86}, {0x1FA90, 0x1FAA8},
	{0x1FAB0, 0x1FAB6}, {0x1FAC0, 0x1FAC2}, {0x1FAD0, 0x1FAD6},
	{0x20000, 0x2FFFD}, {0x30000, 0x3FFFD},
}

var ambiguous = table{
	{0x00A1, 0x00A1}, {0x00A4, 0x00A4}, {0x00A7, 0x00A8},
	{0x00AA, 0x00AA}, {0x00AD, 0x00AE}, {0x00B0, 0x00B4},
	{0x00B6, 0x00BA}, {0x00BC, 0x00BF}, {0x00C6, 0x00C6},
	{0x00D0, 0x00D0}, {0x00D7, 0x00D8}, {0x00DE, 0x00E1},
	{0x00E6, 0x00E6}, {0x00E8, 0x00EA}, {0x00EC, 0x00ED},
	{0x00F0, 0x00F0}, {0x00F2, 0x00F3}, {0x00F7, 0x00FA},
	{0x00FC, 0x00FC}, {0x00FE, 0x00FE}, {0x0101, 0x0101},
	{0x0111, 0x0111}, {0x0113, 0x0113}, {0x011B, 0x011B},
	{0x0126, 0x0127}, {0x012B, 0x012B}, {0x0131, 0x0133},
	{0x0138, 0x0138}, {0x013F, 0x0142}, {0x0144, 0x0144},
	{0x0148, 0x014B}, {0x014D, 0x014D}, {0x0152, 0x0153},
	{0x0166, 0x0167}, {0x016B, 0x016B}, {0x01CE, 0x01CE},
	{0x01D0, 0x01D0}, {0x01D2, 0x01D2}, {0x01D4, 0x01D4},
	{0x01D6, 0x01D6}, {0x01D8, 0x01D8}, {0x01DA, 0x01DA},
	{0x01DC, 0x01DC}, {0x0251, 0x0251}, {0x0261, 0x0261},
	{0x02C4, 0x02C4}, {0x02C7, 0x02C7}, {0x02C9, 0x02CB},
	{0x02CD, 0x02CD}, {0x02D0, 0x02D0}, {0x02D8, 0x02DB},
	{0x02DD, 0x02DD}, {0x02DF, 0x02DF}, {0x0300, 0x036F},
	{0x0391, 0x03A1}, {0x03A3, 0x03A9}, {0x03B1, 0x03C1},
	{0x03C3, 0x03C9}, {0x0401, 0x0401}, {0x0410, 0x044F},
	{0x0451, 0x0451}, {0x2010, 0x2010}, {0x2013, 0x2016},
	{0x2018, 0x2019}, {0x201C, 0x201D}, {0x2020, 0x2022},
	{0x2024, 0x2027}, {0x2030, 0x2030}, {0x2032, 0x2033},
	{0x2035, 0x2035}, {0x203B, 0x203B}, {0x203E, 0x203E},
	{0x2074, 0x2074}, {0x207F, 0x207F}, {0x2081, 0x2084},
	{0x20AC, 0x20AC}, {0x2103, 0x2103}, {0x2105, 0x2105},
	{0x2109, 0x2109}, {0x2113, 0x2113}, {0x2116, 0x2116},
	{0x2121, 0x2122}, {0x2126, 0x2126}, {0x212B, 0x212B},
	{0x2153, 0x2154}, {0x215B, 0x215E}, {0x2160, 0x216B},
	{0x2170, 0x2179}, {0x2189, 0x2189}, {0x2190, 0x2199},
	{0x21B8, 0x21B9}, {0x21D2, 0x21D2}, {0x21D4, 0x21D4},
	{0x21E7, 0x21E7}, {0x2200, 0x2200}, {0x2202, 0x2203},
	{0x2207, 0x2208}, {0x220B, 0x220B}, {0x220F, 0x220F},
	{0x2211, 0x2211}, {0x2215, 0x2215}, {0x221A, 0x221A},
	{0x221D, 0x2220}, {0x2223, 0x2223}, {0x2225, 0x2225},
	{0x2227, 0x222C}, {0x222E, 0x222E}, {0x2234, 0x2237},
	{0x223C, 0x223D}, {0x2248, 0x2248}, {0x224C, 0x224C},
	{0x2252, 0x2252}, {0x2260, 0x2261}, {0x2264, 0x2267},
	{0x226A, 0x226B}, {0x226E, 0x226F}, {0x2282, 0x2283},
	{0x2286, 0x2287}, {0x2295, 0x2295}, {0x2299, 0x2299},
	{0x22A5, 0x22A5}, {0x22BF, 0x22BF}, {0x2312, 0x2312},
	{0x2460, 0x24E9}, {0x24EB, 0x254B}, {0x2550, 0x2573},
	{0x2580, 0x258F}, {0x2592, 0x2595}, {0x25A0, 0x25A1},
	{0x25A3, 0x25A9}, {0x25B2, 0x25B3}, {0x25B6, 0x25B7},
	{0x25BC, 0x25BD}, {0x25C0, 0x25C1}, {0x25C6, 0x25C8},
	{0x25CB, 0x25CB}, {0x25CE, 0x25D1}, {0x25E2, 0x25E5},
	{0x25EF, 0x25EF}, {0x2605, 0x2606}, {0x2609, 0x2609},
	{0x260E, 0x260F}, {0x261C, 0x261C}, {0x261E, 0x261E},
	{0x2640, 0x2640}, {0x2642, 0x2642}, {0x2660, 0x2661},
	{0x2663, 0x2665}, {0x2667, 0x266A}, {0x266C, 0x266D},
	{0x266F, 0x266F}, {0x269E, 0x269F}, {0x26BF, 0x26BF},
	{0x26C6, 0x26CD}, {0x26CF, 0x26D3}, {0x26D5, 0x26E1},
	{0x26E3, 0x26E3}, {0x26E8, 0x26E9}, {0x26EB, 0x26F1},
	{0x26F4, 0x26F4}, {0x26F6, 0x26F9}, {0x26FB, 0x26FC},
	{0x26FE, 0x26FF}, {0x273D, 0x273D}, {0x2776, 0x277F},
	{0x2B56, 0x2B59}, {0x3248, 0x324F}, {0xE000, 0xF8FF},
	{0xFE00, 0xFE0F}, {0xFFFD, 0xFFFD}, {0x1F100, 0x1F10A},
	{0x1F110, 0x1F12D}, {0x1F130, 0x1F169}, {0x1F170, 0x1F18D},
	{0x1F18F, 0x1F190}, {0x1F19B, 0x1F1AC}, {0xE0100, 0xE01EF},
	{0xF0000, 0xFFFFD}, {0x100000, 0x10FFFD},
}
var narrow = table{
	{0x0020, 0x007E}, {0x00A2, 0x00A3}, {0x00A5, 0x00A6},
	{0x00AC, 0x00AC}, {0x00AF, 0x00AF}, {0x27E6, 0x27ED},
	{0x2985, 0x2986},
}

var neutral = table{
	{0x0000, 0x001F}, {0x007F, 0x00A0}, {0x00A9, 0x00A9},
	{0x00AB, 0x00AB}, {0x00B5, 0x00B5}, {0x00BB, 0x00BB},
	{0x00C0, 0x00C5}, {0x00C7, 0x00CF}, {0x00D1, 0x00D6},
	{0x00D9, 0x00DD}, {0x00E2, 0x00E5}, {0x00E7, 0x00E7},
	{0x00EB, 0x00EB}, {0x00EE, 0x00EF}, {0x00F1, 0x00F1},
	{0x00F4, 0x00F6}, {0x00FB, 0x00FB}, {0x00FD, 0x00FD},
	{0x00FF, 0x0100}, {0x0102, 0x0110}, {0x0112, 0x0112},
	{0x0114, 0x011A}, {0x011C, 0x0125}, {0x0128, 0x012A},
	{0x012C, 0x0130}, {0x0134, 0x0137}, {0x0139, 0x013E},
	{0x0143, 0x0143}, {0x0145, 0x0147}, {0x014C, 0x014C},
	{0x014E, 0x0151}, {0x0154, 0x0165}, {0x0168, 0x016A},
	{0x016C, 0x01CD}, {0x01CF, 0x01CF}, {0x01D1, 0x01D1},
	{0x01D3, 0x01D3}, {0x01D5, 0x01D5}, {0x01D7, 0x01D7},
	{0x01D9, 0x01D9}, {0x01DB, 0x01DB}, {0x01DD, 0x0250},
	{0x0252, 0x0260}, {0x0262, 0x02C3}, {0x02C5, 0x02C6},
	{0x02C8, 0x02C8}, {0x02CC, 0x02CC}, {0x02CE, 0x02CF},
	{0x02D1, 0x02D7}, {0x02DC, 0x02DC}, {0x02DE, 0x02DE},
	{0x02E0, 0x02FF}, {0x0370, 0x0377}, {0x037A, 0x037F},
	{0x0384, 0x038A}, {0x038C, 0x038C}, {0x038E, 0x0390},
	{0x03AA, 0x03B0}, {0x03C2, 0x03C2}, {0x03CA, 0x0400},
	{0x0402, 0x040F}, {0x0450, 0x0450}, {0x0452, 0x052F},
	{0x0531, 0x0556}, {0x0559, 0x058A}, {0x058D, 0x058F},
	{0x0591, 0x05C7}, {0x05D0, 0x05EA}, {0x05EF, 0x05F4},
	{0x0600, 0x061C}, {0x061E, 0x070D}, {0x070F, 0x074A},
	{0x074D, 0x07B1}, {0x07C0, 0x07FA}, {0x07FD, 0x082D},
	{0x0830, 0x083E}, {0x0840, 0x085B}, {0x085E, 0x085E},
	{0x0860, 0x086A}, {0x08A0, 0x08B4}, {0x08B6, 0x08C7},
	{0x08D3, 0x0983}, {0x0985, 0x098C}, {0x098F, 0x0990},
	{0x0993, 0x09A8}, {0x09AA, 0x09B0}, {0x09B2, 0x09B2},
	{0x09B6, 0x09B9}, {0x09BC, 0x09C4}, {0x09C7, 0x09C8},
	{0x09CB, 0x09CE}, {0x09D7, 0x09D7}, {0x09DC, 0x09DD},
	{0x09DF, 0x09E3}, {0x09E6, 0x09FE}, {0x0A01, 0x0A03},
	{0x0A05, 0x0A0A}, {0x0A0F, 0x0A10}, {0x0A13, 0x0A28},
	{0x0A2A, 0x0A30}, {0x0A32, 0x0A33}, {0x0A35, 0x0A36},
	{0x0A38, 0x0A39}, {0x0A3C, 0x0A3C}, {0x0A3E, 0x0A42},
	{0x0A47, 0x0A48}, {0x0A4B, 0x0A4D}, {0x0A51, 0x0A51},
	{0x0A59, 0x0A5C}, {0x0A5E, 0x0A5E}, {0x0A66, 0x0A76},
	{0x0A81, 0x0A83}, {0x0A85, 0x0A8D}, {0x0A8F, 0x0A91},
	{0x0A93, 0x0AA8}, {0x0AAA, 0x0AB0}, {0x0AB2, 0x0AB3},
	{0x0AB5, 0x0AB9}, {0x0ABC, 0x0AC5}, {0x0AC7, 0x0AC9},
	{0x0ACB, 0x0ACD}, {0x0AD0, 0x0AD0}, {0x0AE0, 0x0AE3},
	{0x0AE6, 0x0AF1}, {0x0AF9, 0x0AFF}, {0x0B01, 0x0B03},
	{0x0B05, 0x0B0C}, {0x0B0F, 0x0B10}, {0x0B13, 0x0B28},
	{0x0B2A, 0x0B30}, {0x0B32, 0x0B33}, {0x0B35, 0x0B39},
	{0x0B3C, 0x0B44}, {0x0B47, 0x0B48}, {0x0B4B, 0x0B4D},
	{0x0B55, 0x0B57}, {0x0B5C, 0x0B5D}, {0x0B5F, 0x0B63},
	{0x0B66, 0x0B77}, {0x0B82, 0x0B83}, {0x0B85, 0x0B8A},
	{0x0B8E, 0x0B90}, {0x0B92, 0x0B95}, {0x0B99, 0x0B9A},
	{0x0B9C, 0x0B9C}, {0x0B9E, 0x0B9F}, {0x0BA3, 0x0BA4},
	{0x0BA8, 0x0BAA}, {0x0BAE, 0x0BB9}, {0x0BBE, 0x0BC2},
	{0x0BC6, 0x0BC8}, {0x0BCA, 0x0BCD}, {0x0BD0, 0x0BD0},
	{0x0BD7, 0x0BD7}, {0x0BE6, 0x0BFA}, {0x0C00, 0x0C0C},
	{0x0C0E, 0x0C10}, {0x0C12, 0x0C28}, {0x0C2A, 0x0C39},
	{0x0C3D, 0x0C44}, {0x0C46, 0x0C48}, {0x0C4A, 0x0C4D},
	{0x0C55, 0x0C56}, {0x0C58, 0x0C5A}, {0x0C60, 0x0C63},
	{0x0C66, 0x0C6F}, {0x0C77, 0x0C8C}, {0x0C8E, 0x0C90},
	{0x0C92, 0x0CA8}, {0x0CAA, 0x0CB3}, {0x0CB5, 0x0CB9},
	{0x0CBC, 0x0CC4}, {0x0CC6, 0x0CC8}, {0x0CCA, 0x0CCD},
	{0x0CD5, 0x0CD6}, {0x0CDE, 0x0CDE}, {0x0CE0, 0x0CE3},
	{0x0CE6, 0x0CEF}, {0x0CF1, 0x0CF2}, {0x0D00, 0x0D0C},
	{0x0D0E, 0x0D10}, {0x0D12, 0x0D44}, {0x0D46, 0x0D48},
	{0x0D4A, 0x0D4F}, {0x0D54, 0x0D63}, {0x0D66, 0x0D7F},
	{0x0D81, 0x0D83}, {0x0D85, 0x0D96}, {0x0D9A, 0x0DB1},
	{0x0DB3, 0x0DBB}, {0x0DBD, 0x0DBD}, {0x0DC0, 0x0DC6},
	{0x0DCA, 0x0DCA}, {0x0DCF, 0x0DD4}, {0x0DD6, 0x0DD6},
	{0x0DD8, 0x0DDF}, {0x0DE6, 0x0DEF}, {0x0DF2, 0x0DF4},
	{0x0E01, 0x0E3A}, {0x0E3F, 0x0E5B}, {0x0E81, 0x0E82},
	{0x0E84, 0x0E84}, {0x0E86, 0x0E8A}, {0x0E8C, 0x0EA3},
	{0x0EA5, 0x0EA5}, {0x0EA7, 0x0EBD}, {0x0EC0, 0x0EC4},
	{0x0EC6, 0x0EC6}, {0x0EC8, 0x0ECD}, {0x0ED0, 0x0ED9},
	{0x0EDC, 0x0EDF}, {0x0F00, 0x0F47}, {0x0F49, 0x0F6C},
	{0x0F71, 0x0F97}, {0x0F99, 0x0FBC}, {0x0FBE, 0x0FCC},
	{0x0FCE, 0x0FDA}, {0x1000, 0x10C5}, {0x10C7, 0x10C7},
	{0x10CD, 0x10CD}, {0x10D0, 0x10FF}, {0x1160, 0x1248},
	{0x124A, 0x124D}, {0x1250, 0x1256}, {0x1258, 0x1258},
	{0x125A, 0x125D}, {0x1260, 0x1288}, {0x128A, 0x128D},
	{0x1290, 0x12B0}, {0x12B2, 0x12B5}, {0x12B8, 0x12BE},
	{0x12C0, 0x12C0}, {0x12C2, 0x12C5}, {0x12C8, 0x12D6},
	{0x12D8, 0x1310}, {0x1312, 0x1315}, {0x1318, 0x135A},
	{0x135D, 0x137C}, {0x1380, 0x1399}, {0x13A0, 0x13F5},
	{0x13F8, 0x13FD}, {0x1400, 0x169C}, {0x16A0, 0x16F8},
	{0x1700, 0x170C}, {0x170E, 0x1714}, {0x1720, 0x1736},
	{0x1740, 0x1753}, {0x1760, 0x176C}, {0x176E, 0x1770},
	{0x1772, 0x1773}, {0x1780, 0x17DD}, {0x17E0, 0x17E9},
	{0x17F0, 0x17F9}, {0x1800, 0x180E}, {0x1810, 0x1819},
	{0x1820, 0x1878}, {0x1880, 0x18AA}, {0x18B0, 0x18F5},
	{0x1900, 0x191E}, {0x1920, 0x192B}, {0x1930, 0x193B},
	{0x1940, 0x1940}, {0x1944, 0x196D}, {0x1970, 0x1974},
	{0x1980, 0x19AB}, {0x19B0, 0x19C9}, {0x19D0, 0x19DA},
	{0x19DE, 0x1A1B}, {0x1A1E, 0x1A5E}, {0x1A60, 0x1A7C},
	{0x1A7F, 0x1A89}, {0x1A90, 0x1A99}, {0x1AA0, 0x1AAD},
	{0x1AB0, 0x1AC0}, {0x1B00, 0x1B4B}, {0x1B50, 0x1B7C},
	{0x1B80, 0x1BF3}, {0x1BFC, 0x1C37}, {0x1C3B, 0x1C49},
	{0x1C4D, 0x1C88}, {0x1C90, 0x1CBA}, {0x1CBD, 0x1CC7},
	{0x1CD0, 0x1CFA}, {0x1D00, 0x1DF9}, {0x1DFB, 0x1F15},
	{0x1F18, 0x1F1D}, {0x1F20, 0x1F45}, {0x1F48, 0x1F4D},
	{0x1F50, 0x1F57}, {0x1F59, 0x1F59}, {0x1F5B, 0x1F5B},
	{0x1F5D, 0x1F5D}, {0x1F5F, 0x1F7D}, {0x1F80, 0x1FB4},
	{0x1FB6, 0x1FC4}, {0x1FC6, 0x1FD3}, {0x1FD6, 0x1FDB},
	{0x1FDD, 0x1FEF}, {0x1FF2, 0x1FF4}, {0x1FF6, 0x1FFE},
	{0x2000, 0x200F}, {0x2011, 0x2012}, {0x2017, 0x2017},
	{0x201A, 0x201B}, {0x201E, 0x201F}, {0x2023, 0x2023},
	{0x2028, 0x202F}, {0x2031, 0x2031}, {0x2034, 0x2034},
	{0x2036, 0x203A}, {0x203C, 0x203D}, {0x203F, 0x2064},
	{0x2066, 0x2071}, {0x2075, 0x207E}, {0x2080, 0x2080},
	{0x2085, 0x208E}, {0x2090, 0x209C}, {0x20A0, 0x20A8},
	{0x20AA, 0x20AB}, {0x20AD, 0x20BF}, {0x20D0, 0x20F0},
	{0x2100, 0x2102}, {0x2104, 0x2104}, {0x2106, 0x2108},
	{0x210A, 0x2112}, {0x2114, 0x2115}, {0x2117, 0x2120},
	{0x2123, 0x2125}, {0x2127, 0x212A}, {0x212C, 0x2152},
	{0x2155, 0x215A}, {0x215F, 0x215F}, {0x216C, 0x216F},
	{0x217A, 0x2188}, {0x218A, 0x218B}, {0x219A, 0x21B7},
	{0x21BA, 0x21D1}, {0x21D3, 0x21D3}, {0x21D5, 0x21E6},
	{0x21E8, 0x21FF}, {0x2201, 0x2201}, {0x2204, 0x2206},
	{0x2209, 0x220A}, {0x220C, 0x220E}, {0x2210, 0x2210},
	{0x2212, 0x2214}, {0x2216, 0x2219}, {0x221B, 0x221C},
	{0x2221, 0x2222}, {0x2224, 0x2224}, {0x2226, 0x2226},
	{0x222D, 0x222D}, {0x222F, 0x2233}, {0x2238, 0x223B},
	{0x223E, 0x2247}, {0x2249, 0x224B}, {0x224D, 0x2251},
	{0x2253, 0x225F}, {0x2262, 0x2263}, {0x2268, 0x2269},
	{0x226C, 0x226D}, {0x2270, 0x2281}, {0x2284, 0x2285},
	{0x2288, 0x2294}, {0x2296, 0x2298}, {0x229A, 0x22A4},
	{0x22A6, 0x22BE}, {0x22C0, 0x2311}, {0x2313, 0x2319},
	{0x231C, 0x2328}, {0x232B, 0x23E8}, {0x23ED, 0x23EF},
	{0x23F1, 0x23F2}, {0x23F4, 0x2426}, {0x2440, 0x244A},
	{0x24EA, 0x24EA}, {0x254C, 0x254F}, {0x2574, 0x257F},
	{0x2590, 0x2591}, {0x2596, 0x259F}, {0x25A2, 0x25A2},
	{0x25AA, 0x25B1}, {0x25B4, 0x25B5}, {0x25B8, 0x25BB},
	{0x25BE, 0x25BF}, {0x25C2, 0x25C5}, {0x25C9, 0x25CA},
	{0x25CC, 0x25CD}, {0x25D2, 0x25E1}, {0x25E6, 0x25EE},
	{0x25F0, 0x25FC}, {0x25FF, 0x2604}, {0x2607, 0x2608},
	{0x260A, 0x260D}, {0x2610, 0x2613}, {0x2616, 0x261B},
	{0x261D, 0x261D}, {0x261F, 0x263F}, {0x2641, 0x2641},
	{0x2643, 0x2647}, {0x2654, 0x265F}, {0x2662, 0x2662},
	{0x2666, 0x2666}, {0x266B, 0x266B}, {0x266E, 0x266E},
	{0x2670, 0x267E}, {0x2680, 0x2692}, {0x2694, 0x269D},
	{0x26A0, 0x26A0}, {0x26A2, 0x26A9}, {0x26AC, 0x26BC},
	{0x26C0, 0x26C3}, {0x26E2, 0x26E2}, {0x26E4, 0x26E7},
	{0x2700, 0x2704}, {0x2706, 0x2709}, {0x270C, 0x2727},
	{0x2729, 0x273C}, {0x273E, 0x274B}, {0x274D, 0x274D},
	{0x274F, 0x2752}, {0x2756, 0x2756}, {0x2758, 0x2775},
	{0x2780, 0x2794}, {0x2798, 0x27AF}, {0x27B1, 0x27BE},
	{0x27C0, 0x27E5}, {0x27EE, 0x2984}, {0x2987, 0x2B1A},
	{0x2B1D, 0x2B4F}, {0x2B51, 0x2B54}, {0x2B5A, 0x2B73},
	{0x2B76, 0x2B95}, {0x2B97, 0x2C2E}, {0x2C30, 0x2C5E},
	{0x2C60, 0x2CF3}, {0x2CF9, 0x2D25}, {0x2D27, 0x2D27},
	{0x2D2D, 0x2D2D}, {0x2D30, 0x2D67}, {0x2D6F, 0x2D70},
	{0x2D7F, 0x2D96}, {0x2DA0, 0x2DA6}, {0x2DA8, 0x2DAE},
	{0x2DB0, 0x2DB6}, {0x2DB8, 0x2DBE}, {0x2DC0, 0x2DC6},
	{0x2DC8, 0x2DCE}, {0x2DD0, 0x2DD6}, {0x2DD8, 0x2DDE},
	{0x2DE0, 0x2E52}, {0x303F, 0x303F}, {0x4DC0, 0x4DFF},
	{0xA4D0, 0xA62B}, {0xA640, 0xA6F7}, {0xA700, 0xA7BF},
	{0xA7C2, 0xA7CA}, {0xA7F5, 0xA82C}, {0xA830, 0xA839},
	{0xA840, 0xA877}, {0xA880, 0xA8C5}, {0xA8CE, 0xA8D9},
	{0xA8E0, 0xA953}, {0xA95F, 0xA95F}, {0xA980, 0xA9CD},
	{0xA9CF, 0xA9D9}, {0xA9DE, 0xA9FE}, {0xAA00, 0xAA36},
	{0xAA40, 0xAA4D}, {0xAA50, 0xAA59}, {0xAA5C, 0xAAC2},
	{0xAADB, 0xAAF6}, {0xAB01, 0xAB06}, {0xAB09, 0xAB0E},
	{0xAB11, 0xAB16}, {0xAB20, 0xAB26}, {0xAB28, 0xAB2E},
	{0xAB30, 0xAB6B}, {0xAB70, 0xABED}, {0xABF0, 0xABF9},
	{0xD7B0, 0xD7C6}, {0xD7CB, 0xD7FB}, {0xD800, 0xDFFF},
	{0xFB00, 0xFB06}, {0xFB13, 0xFB17}, {0xFB1D, 0xFB36},
	{0xFB38, 0xFB3C}, {0xFB3E, 0xFB3E}, {0xFB40, 0xFB41},
	{0xFB43, 0xFB44}, {0xFB46, 0xFBC1}, {0xFBD3, 0xFD3F},
	{0xFD50, 0xFD8F}, {0xFD92, 0xFDC7}, {0xFDF0, 0xFDFD},
	{0xFE20, 0xFE2F}, {0xFE70, 0xFE74}, {0xFE76, 0xFEFC},
	{0xFEFF, 0xFEFF}, {0xFFF9, 0xFFFC}, {0x10000, 0x1000B},
	{0x1000D, 0x10026}, {0x10028, 0x1003A}, {0x1003C, 0x1003D},
	{0x1003F, 0x1004D}, {0x10050, 0x1005D}, {0x10080, 0x100FA},
	{0x10100, 0x10102}, {0x10107, 0x10133}, {0x10137, 0x1018E},
	{0x10190, 0x1019C}, {0x101A0, 0x101A0}, {0x101D0, 0x101FD},
	{0x10280, 0x1029C}, {0x102A0, 0x102D0}, {0x102E0, 0x102FB},
	{0x10300, 0x10323}, {0x1032D, 0x1034A}, {0x10350, 0x1037A},
	{0x10380, 0x1039D}, {0x1039F, 0x103C3}, {0x103C8, 0x103D5},
	{0x10400, 0x1049D}, {0x104A0, 0x104A9}, {0x104B0, 0x104D3},
	{0x104D8, 0x104FB}, {0x10500, 0x10527}, {0x10530, 0x10563},
	{0x1056F, 0x1056F}, {0x10600, 0x10736}, {0x10740, 0x10755},
	{0x10760, 0x10767}, {0x10800, 0x10805}, {0x10808, 0x10808},
	{0x1080A, 0x10835}, {0x10837, 0x10838}, {0x1083C, 0x1083C},
	{0x1083F, 0x10855}, {0x10857, 0x1089E}, {0x108A7, 0x108AF},
	{0x108E0, 0x108F2}, {0x108F4, 0x108F5}, {0x108FB, 0x1091B},
	{0x1091F, 0x10939}, {0x1093F, 0x1093F}, {0x10980, 0x109B7},
	{0x109BC, 0x109CF}, {0x109D2, 0x10A03}, {0x10A05, 0x10A06},
	{0x10A0C, 0x10A13}, {0x10A15, 0x10A17}, {0x10A19, 0x10A35},
	{0x10A38, 0x10A3A}, {0x10A3F, 0x10A48}, {0x10A50, 0x10A58},
	{0x10A60, 0x10A9F}, {0x10AC0, 0x10AE6}, {0x10AEB, 0x10AF6},
	{0x10B00, 0x10B35}, {0x10B39, 0x10B55}, {0x10B58, 0x10B72},
	{0x10B78, 0x10B91}, {0x10B99, 0x10B9C}, {0x10BA9, 0x10BAF},
	{0x10C00, 0x10C48}, {0x10C80, 0x10CB2}, {0x10CC0, 0x10CF2},
	{0x10CFA, 0x10D27}, {0x10D30, 0x10D39}, {0x10E60, 0x10E7E},
	{0x10E80, 0x10EA9}, {0x10EAB, 0x10EAD}, {0x10EB0, 0x10EB1},
	{0x10F00, 0x10F27}, {0x10F30, 0x10F59}, {0x10FB0, 0x10FCB},
	{0x10FE0, 0x10FF6}, {0x11000, 0x1104D}, {0x11052, 0x1106F},
	{0x1107F, 0x110C1}, {0x110CD, 0x110CD}, {0x110D0, 0x110E8},
	{0x110F0, 0x110F9}, {0x11100, 0x11134}, {0x11136, 0x11147},
	{0x11150, 0x11176}, {0x11180, 0x111DF}, {0x111E1, 0x111F4},
	{0x11200, 0x11211}, {0x11213, 0x1123E}, {0x11280, 0x11286},
	{0x11288, 0x11288}, {0x1128A, 0x1128D}, {0x1128F, 0x1129D},
	{0x1129F, 0x112A9}, {0x112B0, 0x112EA}, {0x112F0, 0x112F9},
	{0x11300, 0x11303}, {0x11305, 0x1130C}, {0x1130F, 0x11310},
	{0x11313, 0x11328}, {0x1132A, 0x11330}, {0x11332, 0x11333},
	{0x11335, 0x11339}, {0x1133B, 0x11344}, {0x11347, 0x11348},
	{0x1134B, 0x1134D}, {0x11350, 0x11350}, {0x11357, 0x11357},
	{0x1135D, 0x11363}, {0x11366, 0x1136C}, {0x11370, 0x11374},
	{0x11400, 0x1145B}, {0x1145D, 0x11461}, {0x11480, 0x114C7},
	{0x114D0, 0x114D9}, {0x11580, 0x115B5}, {0x115B8, 0x115DD},
	{0x11600, 0x11644}, {0x11650, 0x11659}, {0x11660, 0x1166C},
	{0x11680, 0x116B8}, {0x116C0, 0x116C9}, {0x11700, 0x1171A},
	{0x1171D, 0x1172B}, {0x11730, 0x1173F}, {0x11800, 0x1183B},
	{0x118A0, 0x118F2}, {0x118FF, 0x11906}, {0x11909, 0x11909},
	{0x1190C, 0x11913}, {0x11915, 0x11916}, {0x11918, 0x11935},
	{0x11937, 0x11938}, {0x1193B, 0x11946}, {0x11950, 0x11959},
	{0x119A0, 0x119A7}, {0x119AA, 0x119D7}, {0x119DA, 0x119E4},
	{0x11A00, 0x11A47}, {0x11A50, 0x11AA2}, {0x11AC0, 0x11AF8},
	{0x11C00, 0x11C08}, {0x11C0A, 0x11C36}, {0x11C38, 0x11C45},
	{0x11C50, 0x11C6C}, {0x11C70, 0x11C8F}, {0x11C92, 0x11CA7},
	{0x11CA9, 0x11CB6}, {0x11D00, 0x11D06}, {0x11D08, 0x11D09},
	{0x11D0B, 0x11D36}, {0x11D3A, 0x11D3A}, {0x11D3C, 0x11D3D},
	{0x11D3F, 0x11D47}, {0x11D50, 0x11D59}, {0x11D60, 0x11D65},
	{0x11D67, 0x11D68}, {0x11D6A, 0x11D8E}, {0x11D90, 0x11D91},
	{0x11D93, 0x11D98}, {0x11DA0, 0x11DA9}, {0x11EE0, 0x11EF8},
	{0x11FB0, 0x11FB0}, {0x11FC0, 0x11FF1}, {0x11FFF, 0x12399},
	{0x12400, 0x1246E}, {0x12470, 0x12474}, {0x12480, 0x12543},
	{0x13000, 0x1342E}, {0x13430, 0x13438}, {0x14400, 0x14646},
	{0x16800, 0x16A38}, {0x16A40, 0x16A5E}, {0x16A60, 0x16A69},
	{0x16A6E, 0x16A6F}, {0x16AD0, 0x16AED}, {0x16AF0, 0x16AF5},
	{0x16B00, 0x16B45}, {0x16B50, 0x16B59}, {0x16B5B, 0x16B61},
	{0x16B63, 0x16B77}, {0x16B7D, 0x16B8F}, {0x16E40, 0x16E9A},
	{0x16F00, 0x16F4A}, {0x16F4F, 0x16F87}, {0x16F8F, 0x16F9F},
	{0x1BC00, 0x1BC6A}, {0x1BC70, 0x1BC7C}, {0x1BC80, 0x1BC88},
	{0x1BC90, 0x1BC99}, {0x1BC9C, 0x1BCA3}, {0x1D000, 0x1D0F5},
	{0x1D100, 0x1D126}, {0x1D129, 0x1D1E8}, {0x1D200, 0x1D245},
	{0x1D2E0, 0x1D2F3}, {0x1D300, 0x1D356}, {0x1D360, 0x1D378},
	{0x1D400, 0x1D454}, {0x1D456, 0x1D49C}, {0x1D49E, 0x1D49F},
	{0x1D4A2, 0x1D4A2}, {0x1D4A5, 0x1D4A6}, {0x1D4A9, 0x1D4AC},
	{0x1D4AE, 0x1D4B9}, {0x1D4BB, 0x1D4BB}, {0x1D4BD, 0x1D4C3},
	{0x1D4C5, 0x1D505}, {0x1D507, 0x1D50A}, {0x1D50D, 0x1D514},
	{0x1D516, 0x1D51C}, {0x1D51E, 0x1D539}, {0x1D53B, 0x1D53E},
	{0x1D540, 0x1D544}, {0x1D546, 0x1D546}, {0x1D54A, 0x1D550},
	{0x1D552, 0x1D6A5}, {0x1D6A8, 0x1D7CB}, {0x1D7CE, 0x1DA8B},
	{0x1DA9B, 0x1DA9F}, {0x1DAA1, 0x1DAAF}, {0x1E000, 0x1E006},
	{0x1E008, 0x1E018}, {0x1E01B, 0x1E021}, {0x1E023, 0x1E024},
	{0x1E026, 0x1E02A}, {0x1E100, 0x1E12C}, {0x1E130, 0x1E13D},
	{0x1E140, 0x1E149}, {0x1E14E, 0x1E14F}, {0x1E2C0, 0x1E2F9},
	{0x1E2FF, 0x1E2FF}, {0x1E800, 0x1E8C4}, {0x1E8C7, 0x1E8D6},
	{0x1E900, 0x1E94B}, {0x1E950, 0x1E959}, {0x1E95E, 0x1E95F},
	{0x1EC71, 0x1ECB4}, {0x1ED01, 0x1ED3D}, {0x1EE00, 0x1EE03},
	{0x1EE05, 0x1EE1F}, {0x1EE21, 0x1EE22}, {0x1EE24, 0x1EE24},
	{0x1EE27, 0x1EE27}, {0x1EE29, 0x1EE32}, {0x1EE34, 0x1EE37},
	{0x1EE39, 0x1EE39}, {0x1EE3B, 0x1EE3B}, {0x1EE42, 0x1EE42},
	{0x1EE47, 0x1EE47}, {0x1EE49, 0x1EE49}, {0x1EE4B, 0x1EE4B},
	{0x1EE4D, 0x1EE4F}, {0x1EE51, 0x1EE52}, {0x1EE54, 0x1EE54},
	{0x1EE57, 0x1EE57}, {0x1EE59, 0x1EE59}, {0x1EE5B, 0x1EE5B},
	{0x1EE5D, 0x1EE5D}, {0x1EE5F, 0x1EE5F}, {0x1EE61, 0x1EE62},
	{0x1EE64, 0x1EE64}, {0x1EE67, 0x1EE6A}, {0x1EE6C, 0x1EE72},
	{0x1EE74, 0x1EE77}, {0x1EE79, 0x1EE7C}, {0x1EE7E, 0x1EE7E},
	{0x1EE80, 0x1EE89}, {0x1EE8B, 0x1EE9B}, {0x1EEA1, 0x1EEA3},
	{0x1EEA5, 0x1EEA9}, {0x1EEAB, 0x1EEBB}, {0x1EEF0, 0x1EEF1},
	{0x1F000, 0x1F003}, {0x1F005, 0x1F02B}, {0x1F030, 0x1F093},
	{0x1F0A0, 0x1F0AE}, {0x1F0B1, 0x1F0BF}, {0x1F0C1, 0x1F0CE},
	{0x1F0D1, 0x1F0F5}, {0x1F10B, 0x1F10F}, {0x1F12E, 0x1F12F},
	{0x1F16A, 0x1F16F}, {0x1F1AD, 0x1F1AD}, {0x1F1E6, 0x1F1FF},
	{0x1F321, 0x1F32C}, {0x1F336, 0x1F336}, {0x1F37D, 0x1F37D},
	{0x1F394, 0x1F39F}, {0x1F3CB, 0x1F3CE}, {0x1F3D4, 0x1F3DF},
	{0x1F3F1, 0x1F3F3}, {0x1F3F5, 0x1F3F7}, {0x1F43F, 0x1F43F},
	{0x1F441, 0x1F441}, {0x1F4FD, 0x1F4FE}, {0x1F53E, 0x1F54A},
	{0x1F54F, 0x1F54F}, {0x1F568, 0x1F579}, {0x1F57B, 0x1F594},
	{0x1F597, 0x1F5A3}, {0x1F5A5, 0x1F5FA}, {0x1F650, 0x1F67F},
	{0x1F6C6, 0x1F6CB}, {0x1F6CD, 0x1F6CF}, {0x1F6D3, 0x1F6D4},
	{0x1F6E0, 0x1F6EA}, {0x1F6F0, 0x1F6F3}, {0x1F700, 0x1F773},
	{0x1F780, 0x1F7D8}, {0x1F800, 0x1F80B}, {0x1F810, 0x1F847},
	{0x1F850, 0x1F859}, {0x1F860, 0x1F887}, {0x1F890, 0x1F8AD},
	{0x1F8B0, 0x1F8B1}, {0x1F900, 0x1F90B}, {0x1F93B, 0x1F93B},
	{0x1F946, 0x1F946}, {0x1FA00, 0x1FA53}, {0x1FA60, 0x1FA6D},
	{0x1FB00, 0x1FB92}, {0x1FB94, 0x1FBCA}, {0x1FBF0, 0x1FBF9},
	{0xE0001, 0xE0001}, {0xE0020, 0xE007F},
}

var emoji = table{
	{0x203C, 0x203C}, {0x2049, 0x2049}, {0x2122, 0x2122},
	{0x2139, 0x2139}, {0x2194, 0x2199}, {0x21A9, 0x21AA},
	{0x231A, 0x231B}, {0x2328, 0x2328}, {0x2388, 0x2388},
	{0x23CF, 0x23CF}, {0x23E9, 0x23F3}, {0x23F8, 0x23FA},
	{0x24C2, 0x24C2}, {0x25AA, 0x25AB}, {0x25B6, 0x25B6},
	{0x25C0, 0x25C0}, {0x25FB, 0x25FE}, {0x2600, 0x2605},
	{0x2607, 0x2612}, {0x2614, 0x2685}, {0x2690, 0x2705},
	{0x2708, 0x2712}, {0x2714, 0x2714}, {0x2716, 0x2716},
	{0x271D, 0x271D}, {0x2721, 0x2721}, {0x2728, 0x2728},
	{0x2733, 0x2734}, {0x2744, 0x2744}, {0x2747, 0x2747},
	{0x274C, 0x274C}, {0x274E, 0x274E}, {0x2753, 0x2755},
	{0x2757, 0x2757}, {0x2763, 0x2767}, {0x2795, 0x2797},
	{0x27A1, 0x27A1}, {0x27B0, 0x27B0}, {0x27BF, 0x27BF},
	{0x2934, 0x2935}, {0x2B05, 0x2B07}, {0x2B1B, 0x2B1C},
	{0x2B50, 0x2B50}, {0x2B55, 0x2B55}, {0x3030, 0x3030},
	{0x303D, 0x303D}, {0x3297, 0x3297}, {0x3299, 0x3299},
	{0x1F000, 0x1F0FF}, {0x1F10D, 0x1F10F}, {0x1F12F, 0x1F12F},
	{0x1F16C, 0x1F171}, {0x1F17E, 0x1F17F}, {0x1F18E, 0x1F18E},
	{0x1F191, 0x1F19A}, {0x1F1AD, 0x1F1FF}, {0x1F201, 0x1F20F},
	{0x1F21A, 0x1F21A}, {0x1F22F, 0x1F22F}, {0x1F232, 0x1F23A},
	{0x1F23C, 0x1F23F}, {0x1F249, 0x1F3FA}, {0x1F400, 0x1F53D},
	{0x1F546, 0x1F64F}, {0x1F680, 0x1F6FF}, {0x1F774, 0x1F77F},
	{0x1F7D5, 0x1F7FF}, {0x1F80C, 0x1F80F}, {0x1F848, 0x1F84F},
	{0x1F85A, 0x1F85F}, {0x1F888, 0x1F88F}, {0x1F8AE, 0x1F8FF},
	{0x1F90C, 0x1F93A}, {0x1F93C, 0x1F945}, {0x1F947, 0x1FAFF},
	{0x1FC00, 0x1FFFD},
}
