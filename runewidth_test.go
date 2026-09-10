//go:build !js && !appengine
// +build !js,!appengine

package runewidth

import (
	"crypto/sha256"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strings"
	"sync"
	"testing"
	"unicode/utf8"

	"github.com/clipperhouse/uax29/v2/graphemes"
)

var _ sort.Interface = (*table)(nil) // ensure that type "table" does implement sort.Interface

func init() {
	os.Setenv("RUNEWIDTH_EASTASIAN", "")
	handleEnv()
}

func (t table) Len() int {
	return len(t)
}

func (t table) Less(i, j int) bool {
	return t[i].first < t[j].first
}

func (t *table) Swap(i, j int) {
	(*t)[i], (*t)[j] = (*t)[j], (*t)[i]
}

type tableInfo struct {
	tbl     table
	name    string
	wantN   int
	wantSHA string
}

var tables = []tableInfo{
	{private, "private", 137468, "a4a641206dc8c5de80bd9f03515a54a706a5a4904c7684dc6a33d65c967a51b2"},
	{nonprint, "nonprint", 2143, "288904683eb225e7c4c0bd3ee481b53e8dace404ec31d443afdbc4d13729fe95"},
	{combining, "combining", 2081, "bbb88427a84cf23bd601d560b32ffd88b1d0c1aeb365b54af37f1ad7d7e6944e"},
	{doublewidth, "doublewidth", 182876, "55dcb1b999d6356d1a083085bb053bdeafc6dda05dec002617d85fda2a82d496"},
	{ambiguous, "ambiguous", 138483, "f4ed2dd733c0821cf6297fc24be0baea527ec7cad10d23b0ac7f57dcdf344cdb"},
	{emoji, "emoji", 2846, "09914b87febaa5493f2420a58f03dd6b026fa665b7c811abc7423a26a9b442c3"},
	{narrow, "narrow", 111, "fa897699c5e3cd9141c638d539331b0bdd508b874e22996c5e929767d455fc5a"},
	{neutral, "neutral", 33695, "f6af4edbdfe84d0c4c4419da8e2f86e49248010bad19c4ed7bf1897696db9084"},
}

func TestTableChecksums(t *testing.T) {
	for _, ti := range tables {
		gotN := 0
		buf := make([]byte, utf8.MaxRune+1)
		for r := rune(0); r <= utf8.MaxRune; r++ {
			if inTable(r, ti.tbl) {
				gotN++
				buf[r] = 1
			}
		}
		gotSHA := fmt.Sprintf("%x", sha256.Sum256(buf))
		if gotN != ti.wantN || gotSHA != ti.wantSHA {
			t.Errorf("table = %s,\n\tn = %d want %d,\n\tsha256 = %s want %s", ti.name, gotN, ti.wantN, gotSHA, ti.wantSHA)
		}
	}
}

func TestRuneWidthChecksums(t *testing.T) {
	var testcases = []struct {
		name           string
		eastAsianWidth bool
		wantSHA        string
	}{
		{"ea-no", false, "b166b7c41c9231ce5a3f05d94b0da79c245714d5f708b3002ccef55bda38152d"},
		{"ea-yes", true, "d46d9c64c35351d6daa041d69092fb1181d63096ca089149ecc4b922bc0e2cf8"},
	}

	for _, testcase := range testcases {
		c := NewCondition()
		c.EastAsianWidth = testcase.eastAsianWidth
		buf := make([]byte, utf8.MaxRune+1)
		for r := rune(0); r <= utf8.MaxRune; r++ {
			buf[r] = byte(c.RuneWidth(r))
		}
		gotSHA := fmt.Sprintf("%x", sha256.Sum256(buf))
		if gotSHA != testcase.wantSHA {
			t.Errorf("TestRuneWidthChecksums = %s,\n\tsha256 = %s want %s",
				testcase.name, gotSHA, testcase.wantSHA)
		}

		// Test with LUT
		c.CreateLUT()
		for r := rune(0); r <= utf8.MaxRune; r++ {
			buf[r] = byte(c.RuneWidth(r))
		}
		gotSHA = fmt.Sprintf("%x", sha256.Sum256(buf))
		if gotSHA != testcase.wantSHA {
			t.Errorf("TestRuneWidthChecksums = %s,\n\tsha256 = %s want %s",
				testcase.name, gotSHA, testcase.wantSHA)
		}
	}
}

func TestStrictWidthLUT(t *testing.T) {
	buildStrictWidthLUT()
	lut := &strictWidthLUT
	for r := rune(0); r <= utf8.MaxRune; r++ {
		if got, want := int(lut[0][r]), runeWidthNoLUT(r, false, true); got != want {
			t.Errorf("strictWidthLUT[0][%U] = %d, want %d", r, got, want)
		}
		if got, want := int(lut[1][r]), runeWidthNoLUT(r, true, true); got != want {
			t.Errorf("strictWidthLUT[1][%U] = %d, want %d", r, got, want)
		}
	}
}

// TestRuneWidthConcurrent exercises concurrent first use of RuneWidth, so
// that running it alone under -race checks readers against the lazy LUT
// build. Run before other tests have built the LUT to be effective.
func TestRuneWidthConcurrent(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 8; i++ {
		wg.Add(1)
		go func(ea bool) {
			defer wg.Done()
			c := NewCondition()
			c.EastAsianWidth = ea
			for r := rune(0); r <= utf8.MaxRune; r += 7 {
				if w := c.RuneWidth(r); w < 0 || w > 2 {
					t.Errorf("RuneWidth(%U) = %d", r, w)
					return
				}
			}
		}(i%2 == 0)
	}
	wg.Wait()
}

func TestDefaultLUT(t *testing.T) {
	var testcases = []struct {
		name           string
		eastAsianWidth bool
		wantSHA        string
	}{
		{"ea-no", false, "b166b7c41c9231ce5a3f05d94b0da79c245714d5f708b3002ccef55bda38152d"},
		{"ea-yes", true, "d46d9c64c35351d6daa041d69092fb1181d63096ca089149ecc4b922bc0e2cf8"},
	}

	old := os.Getenv("RUNEWIDTH_EASTASIAN")
	defer os.Setenv("RUNEWIDTH_EASTASIAN", old)

	CreateLUT()
	for _, testcase := range testcases {
		c := DefaultCondition

		if testcase.eastAsianWidth {
			os.Setenv("RUNEWIDTH_EASTASIAN", "1")
		} else {
			os.Setenv("RUNEWIDTH_EASTASIAN", "0")
		}
		handleEnv()

		buf := make([]byte, utf8.MaxRune+1)
		for r := rune(0); r <= utf8.MaxRune; r++ {
			buf[r] = byte(c.RuneWidth(r))
		}
		gotSHA := fmt.Sprintf("%x", sha256.Sum256(buf))
		if gotSHA != testcase.wantSHA {
			t.Errorf("TestRuneWidthChecksums = %s,\n\tsha256 = %s want %s",
				testcase.name, gotSHA, testcase.wantSHA)
		}
	}
	// Remove for other tests.
	DefaultCondition.combinedLut = nil
}

func checkInterval(first, last rune) bool {
	return first >= 0 && first <= utf8.MaxRune &&
		last >= 0 && last <= utf8.MaxRune &&
		first <= last
}

func isCompact(t *testing.T, ti *tableInfo) bool {
	tbl := ti.tbl
	for i := range tbl {
		e := tbl[i]
		if !checkInterval(e.first, e.last) { // sanity check
			t.Errorf("table invalid: table = %s index = %d %v", ti.name, i, e)
			return false
		}
		if i+1 < len(tbl) && e.last+1 >= tbl[i+1].first { // can be combined into one entry
			t.Errorf("table not compact: table = %s index = %d %v %v", ti.name, i, e, tbl[i+1])
			return false
		}
	}
	return true
}

func TestSorted(t *testing.T) {
	for _, ti := range tables {
		if !sort.IsSorted(&ti.tbl) {
			t.Errorf("table not sorted: %s", ti.name)
		}
		if !isCompact(t, &ti) {
			t.Errorf("table not compact: %s", ti.name)
		}
	}
}

var runewidthtests = []struct {
	in     rune
	out    int
	eaout  int
	nseout int
}{
	{'世', 2, 2, 2},
	{'界', 2, 2, 2},
	{'ｾ', 1, 1, 1},
	{'ｶ', 1, 1, 1},
	{'ｲ', 1, 1, 1},
	{'☆', 1, 2, 2}, // double width in ambiguous
	{'☺', 1, 1, 2},
	{'☻', 1, 1, 1},
	{'♥', 1, 2, 2},
	{'♦', 1, 1, 2},
	{'♣', 1, 2, 2},
	{'♠', 1, 2, 2},
	{'♂', 1, 2, 2},
	{'♀', 1, 2, 2},
	{'♪', 1, 2, 2},
	{'♫', 1, 1, 1},
	{'☼', 1, 1, 1},
	{'↕', 1, 2, 2},
	{'‼', 1, 1, 2},
	{'↔', 1, 2, 2},
	{'\x00', 0, 0, 0},
	{'\x01', 0, 0, 0},
	{'\u0300', 0, 0, 0},
	{'\u2028', 0, 0, 0},
	{'\u2029', 0, 0, 0},
	{'a', 1, 1, 1}, // ASCII classified as "na" (narrow)
	{'⟦', 1, 1, 1}, // non-ASCII classified as "na" (narrow)
	{'👁', 1, 1, 2},
	{'\u093E', 1, 1, 1}, // DEVANAGARI VOWEL SIGN AA (Mc) - spacing mark occupies its own cell
	{'\u0941', 0, 0, 0}, // DEVANAGARI VOWEL SIGN U (Mn) - nonspacing mark
	{'\u094D', 0, 0, 0}, // DEVANAGARI SIGN VIRAMA (Mn)
	{'\u0915', 1, 1, 1}, // DEVANAGARI LETTER KA (Lo) - base consonant, width 1
}

func TestRuneWidth(t *testing.T) {
	c := NewCondition()
	c.EastAsianWidth = false
	for _, tt := range runewidthtests {
		if out := c.RuneWidth(tt.in); out != tt.out {
			t.Errorf("RuneWidth(%q) = %d, want %d (EastAsianWidth=false)", tt.in, out, tt.out)
		}
	}
	c.EastAsianWidth = true
	for _, tt := range runewidthtests {
		if out := c.RuneWidth(tt.in); out != tt.eaout {
			t.Errorf("RuneWidth(%q) = %d, want %d (EastAsianWidth=true)", tt.in, out, tt.eaout)
		}
	}
	c.StrictEmojiNeutral = false
	for _, tt := range runewidthtests {
		if out := c.RuneWidth(tt.in); out != tt.nseout {
			t.Errorf("RuneWidth(%q) = %d, want %d (StrictEmojiNeutral=false)", tt.in, out, tt.eaout)
		}
	}
}

var isambiguouswidthtests = []struct {
	in  rune
	out bool
}{
	{'世', false},
	{'■', true},
	{'界', false},
	{'○', true},
	{'㈱', false},
	{'①', true},
	{'②', true},
	{'③', true},
	{'④', true},
	{'⑤', true},
	{'⑥', true},
	{'⑦', true},
	{'⑧', true},
	{'⑨', true},
	{'⑩', true},
	{'⑪', true},
	{'⑫', true},
	{'⑬', true},
	{'⑭', true},
	{'⑮', true},
	{'⑯', true},
	{'⑰', true},
	{'⑱', true},
	{'⑲', true},
	{'⑳', true},
	{'☆', true},
}

func TestIsAmbiguousWidth(t *testing.T) {
	for _, tt := range isambiguouswidthtests {
		if out := IsAmbiguousWidth(tt.in); out != tt.out {
			t.Errorf("IsAmbiguousWidth(%q) = %v, want %v", tt.in, out, tt.out)
		}
	}
}

var stringwidthtests = []struct {
	in    string
	out   int
	eaout int
}{
	{"■㈱の世界①", 10, 12},
	{"スター☆", 7, 8},
	{"つのだ☆HIRO", 11, 12},
	{"खा", 2, 2},    // base consonant + spacing mark (Mc), one cluster, two cells
	{"हिन्द", 4, 4}, // spacing marks count, nonspacing marks (Mn) do not
	{"फ़ल", 2, 2},   // nukta (Mn) is zero width
	{"🇩🇰", 2, 2},    // regional indicator pair renders as one two-cell flag
}

func TestStringWidth(t *testing.T) {
	c := NewCondition()
	c.EastAsianWidth = false
	for _, tt := range stringwidthtests {
		if out := c.StringWidth(tt.in); out != tt.out {
			t.Errorf("StringWidth(%q) = %d, want %d", tt.in, out, tt.out)
		}
	}
	c.EastAsianWidth = true
	for _, tt := range stringwidthtests {
		if out := c.StringWidth(tt.in); out != tt.eaout {
			t.Errorf("StringWidth(%q) = %d, want %d (EA)", tt.in, out, tt.eaout)
		}
	}
}

func TestStringWidthInvalid(t *testing.T) {
	s := "こんにちわ\x00世界"
	if out := StringWidth(s); out != 14 {
		t.Errorf("StringWidth(%q) = %d, want %d", s, out, 14)
	}
}

func TestTruncateSmaller(t *testing.T) {
	s := "あいうえお"
	expected := "あいうえお"

	if out := Truncate(s, 10, "..."); out != expected {
		t.Errorf("Truncate(%q) = %q, want %q", s, out, expected)
	}
}

func TestTruncate(t *testing.T) {
	s := "あいうえおあいうえおえおおおおおおおおおおおおおおおおおおおおおおおおおおおおおお"
	expected := "あいうえおあいうえおえおおおおおおおおおおおおおおおおおおおおおおおおおおお..."
	out := Truncate(s, 80, "...")
	if out != expected {
		t.Errorf("Truncate(%q) = %q, want %q", s, out, expected)
	}
	width := StringWidth(out)
	if width != 79 {
		t.Errorf("width of Truncate(%q) should be %d, but %d", s, 79, width)
	}
}

func TestTruncateFit(t *testing.T) {
	s := "aあいうえおあいうえおえおおおおおおおおおおおおおおおおおおおおおおおおおおおおおお"
	expected := "aあいうえおあいうえおえおおおおおおおおおおおおおおおおおおおおおおおおおおお..."

	out := Truncate(s, 80, "...")
	if out != expected {
		t.Errorf("Truncate(%q) = %q, want %q", s, out, expected)
	}
	width := StringWidth(out)
	if width != 80 {
		t.Errorf("width of Truncate(%q) should be %d, but %d", s, 80, width)
	}
}

func TestTruncateJustFit(t *testing.T) {
	s := "あいうえおあいうえおえおおおおおおおおおおおおおおおおおおおおおおおおおおおおお"
	expected := "あいうえおあいうえおえおおおおおおおおおおおおおおおおおおおおおおおおおおおおお"

	out := Truncate(s, 80, "...")
	if out != expected {
		t.Errorf("Truncate(%q) = %q, want %q", s, out, expected)
	}
	width := StringWidth(out)
	if width != 80 {
		t.Errorf("width of Truncate(%q) should be %d, but %d", s, 80, width)
	}
}

func TestWrap(t *testing.T) {
	s := `東京特許許可局局長はよく柿喰う客だ/東京特許許可局局長はよく柿喰う客だ
123456789012345678901234567890

END`
	expected := `東京特許許可局局長はよく柿喰う
客だ/東京特許許可局局長はよく
柿喰う客だ
123456789012345678901234567890

END`

	if out := Wrap(s, 30); out != expected {
		t.Errorf("Wrap(%q) = %q, want %q", s, out, expected)
	}
}

func TestTruncateNoNeeded(t *testing.T) {
	s := "あいうえおあい"
	expected := "あいうえおあい"

	if out := Truncate(s, 80, "..."); out != expected {
		t.Errorf("Truncate(%q) = %q, want %q", s, out, expected)
	}
}

var truncatelefttests = []struct {
	s      string
	w      int
	prefix string
	out    string
}{
	{"source", 4, "", "ce"},
	{"source", 4, "...", "...ce"},
	{"あいうえお", 6, "", "えお"},
	{"あいうえお", 6, "...", "...えお"},
	{"あいうえお", 10, "", ""},
	{"あいうえお", 10, "...", "..."},
	{"あいうえお", 5, "", " えお"},
	{"Aあいうえお", 5, "", "うえお"},
}

func TestTruncateLeft(t *testing.T) {
	t.Parallel()

	for _, tt := range truncatelefttests {
		if out := TruncateLeft(tt.s, tt.w, tt.prefix); out != tt.out {
			t.Errorf("TruncateLeft(%q) = %q, want %q", tt.s, out, tt.out)
		}
	}
}

var truncateprefixtests = []struct {
	s      string
	w      int
	prefix string
	out    string
}{
	{"source", 4, "*", "*rce"},
	{"source", 6, "*", "source"},
	{"あいうえお", 4, "*", "*お"},
	{"あいうえお", 10, "*", "あいうえお"},
	{"Aあいうえお", 5, "*", "*えお"},

	{"source", 4, "", "urce"},
	{"source", 4, "...", "...e"},
	{"あいうえお", 6, "", "うえお"},
	{"あいうえお", 6, "...", "...お"},
	{"あいうえお", 10, "", "あいうえお"},
	{"あいうえお", 10, "...", "あいうえお"},
	{"あいうえお", 5, "", "えお"},

	{"source", 1, "*", "*"},
	{"source", 0, "*", "*"}, // same as Truncate

	// multi-rune grapheme clusters must be dropped or kept as a whole
	{"👨‍👩‍👧cde", 4, "*", "*cde"},
	{"खाabc", 3, "*", "*bc"},
	{"🇩🇰abc", 3, "", "abc"},
	{"あ🏳️‍🌈い", 4, "*", "*い"},
}

func TestTruncatePrefix(t *testing.T) {
	t.Parallel()

	for _, tt := range truncateprefixtests {
		if out := TruncatePrefix(tt.s, tt.w, tt.prefix); out != tt.out {
			t.Errorf("TruncatePrefix(%q) = %q, want %q", tt.s, out, tt.out)
		}
	}
}

var isneutralwidthtests = []struct {
	in  rune
	out bool
}{
	{'→', false},
	{'┊', false},
	{'┈', false},
	{'～', false},
	{'└', false},
	{'⣀', true},
	{'⣀', true},
}

func TestIsNeutralWidth(t *testing.T) {
	for _, tt := range isneutralwidthtests {
		if out := IsNeutralWidth(tt.in); out != tt.out {
			t.Errorf("IsNeutralWidth(%q) = %v, want %v", tt.in, out, tt.out)
		}
	}
}

var iscombiningwidthtests = []struct {
	in  rune
	out bool
}{
	{'\u0300', true},     // COMBINING GRAVE ACCENT
	{'\uFE0F', true},     // VARIATION SELECTOR-16
	{'\uFE0E', true},     // VARIATION SELECTOR-15
	{'\u180B', true},     // MONGOLIAN FREE VARIATION SELECTOR ONE
	{'\U000E0100', true}, // VARIATION SELECTOR-17
	{'A', false},
	{'☆', false},
}

func TestIsCombiningWidth(t *testing.T) {
	for _, tt := range iscombiningwidthtests {
		if out := IsCombiningWidth(tt.in); out != tt.out {
			t.Errorf("IsCombiningWidth(%q) = %v, want %v", tt.in, out, tt.out)
		}
	}
}

func TestFillLeft(t *testing.T) {
	s := "あxいうえお"
	expected := "    あxいうえお"

	if out := FillLeft(s, 15); out != expected {
		t.Errorf("FillLeft(%q) = %q, want %q", s, out, expected)
	}
}

func TestFillLeftFit(t *testing.T) {
	s := "あいうえお"
	expected := "あいうえお"

	if out := FillLeft(s, 10); out != expected {
		t.Errorf("FillLeft(%q) = %q, want %q", s, out, expected)
	}
}

func TestFillRight(t *testing.T) {
	s := "あxいうえお"
	expected := "あxいうえお    "

	if out := FillRight(s, 15); out != expected {
		t.Errorf("FillRight(%q) = %q, want %q", s, out, expected)
	}
}

func TestFillRightFit(t *testing.T) {
	s := "あいうえお"
	expected := "あいうえお"

	if out := FillRight(s, 10); out != expected {
		t.Errorf("FillRight(%q) = %q, want %q", s, out, expected)
	}
}

func TestEnv(t *testing.T) {
	old := os.Getenv("RUNEWIDTH_EASTASIAN")
	defer os.Setenv("RUNEWIDTH_EASTASIAN", old)

	os.Setenv("RUNEWIDTH_EASTASIAN", "0")
	handleEnv()

	if w := RuneWidth('│'); w != 1 {
		t.Errorf("RuneWidth('│') = %d, want %d", w, 1)
	}
}

func TestZeroWidthJoiner(t *testing.T) {
	c := NewCondition()

	var tests = []struct {
		in   string
		want int
	}{
		{"👩", 2},
		{"👩\u200d", 2},
		{"👩\u200d🍳", 2},
		{"\u200d🍳", 2},
		{"👨\u200d👨", 2},
		{"👨\u200d👨\u200d👧", 2},
		{"🏳️\u200d🌈", 2},
		{"あ👩\u200d🍳い", 6},
		{"あ\u200d🍳い", 6},
		{"あ\u200dい", 4},
	}

	for _, tt := range tests {
		if got := c.StringWidth(tt.in); got != tt.want {
			t.Errorf("StringWidth(%q) = %d, want %d", tt.in, got, tt.want)
		}
	}
}

func TestZeroWidthJoinerFlag(t *testing.T) {
	// The deprecated flag must compile and must not change any result.
	for _, zwj := range []bool{true, false} {
		c := NewCondition()
		c.ZeroWidthJoiner = zwj
		if got := c.StringWidth("👨‍👨‍👧"); got != 2 {
			t.Errorf("StringWidth with ZeroWidthJoiner=%v = %d, want 2", zwj, got)
		}
	}
}

func TestWrapNonPositiveWidth(t *testing.T) {
	// Wrap only inserts newlines, so stripping them must give the input
	// back for any width; w == 0 used to panic on the capacity hint.
	for _, w := range []int{0, -1} {
		got := strings.ReplaceAll(Wrap("hello", w), "\n", "")
		if got != "hello" {
			t.Errorf("Wrap(%q, %d) stripped = %q, want %q", "hello", w, got, "hello")
		}
	}
}

func TestCreateLUTRebuildsAfterFlagChange(t *testing.T) {
	savedEA, savedLut := DefaultCondition.EastAsianWidth, DefaultCondition.combinedLut
	defer func() {
		DefaultCondition.EastAsianWidth, DefaultCondition.combinedLut = savedEA, savedLut
	}()
	for _, ea := range []bool{false, true} {
		DefaultCondition.EastAsianWidth = !ea
		CreateLUT()
		DefaultCondition.EastAsianWidth = ea
		CreateLUT()
		want := (&Condition{EastAsianWidth: ea}).RuneWidth('±')
		if got := RuneWidth('±'); got != want {
			t.Errorf("EastAsianWidth=%v after rebuild: RuneWidth(U+00B1) = %d, want %d", ea, got, want)
		}
	}
}

func TestCreateLUTSkipsRebuildWhenFlagsUnchanged(t *testing.T) {
	savedEA, savedSEN := DefaultCondition.EastAsianWidth, DefaultCondition.StrictEmojiNeutral
	savedLut := DefaultCondition.combinedLut
	defer func() {
		DefaultCondition.EastAsianWidth, DefaultCondition.StrictEmojiNeutral = savedEA, savedSEN
		DefaultCondition.combinedLut = savedLut
	}()

	DefaultCondition.combinedLut = nil
	CreateLUT()
	// A rebuild overwrites every entry, so a poisoned byte surviving the
	// second call is what shows the table was left alone.
	DefaultCondition.combinedLut[0] = 0xff
	CreateLUT()
	if DefaultCondition.combinedLut[0] != 0xff {
		t.Error("CreateLUT rebuilt the table although no flag changed")
	}
}

func TestWrapGraphemeCluster(t *testing.T) {
	var tests = []struct {
		s        string
		w        int
		expected string
	}{
		// a multi-rune cluster is one glyph of at most two cells, so it fits
		// in two columns and must not be split across lines
		{"👨‍👩‍👧‍👦", 2, "👨‍👩‍👧‍👦"},
		{"👩🏽", 2, "👩🏽"},
		{"👩🏽あ", 2, "👩🏽\nあ"},
		// breaks still land between clusters
		{"あい", 3, "あ\nい"},
		{"あい", 4, "あい"},
		// a newline resets the column, CRLF included
		{"あい\nうえ", 4, "あい\nうえ"},
		{"あい\r\nうえ", 4, "あい\r\nうえ"},
		{"abcdef", 4, "abcd\nef"},
		{"abcd\nef", 4, "abcd\nef"},
	}
	for _, tt := range tests {
		if out := Wrap(tt.s, tt.w); out != tt.expected {
			t.Errorf("Wrap(%q, %d) = %q, want %q", tt.s, tt.w, out, tt.expected)
		}
	}
}

// clusterWidth and clusterWrap are StringWidth and Wrap with the fast path
// taken out, the behaviour the fast path has to reproduce exactly.
func clusterWidth(c *Condition, s string) int {
	width := 0
	g := graphemes.FromString(s)
	for g.Next() {
		width += c.graphemeWidth(g.Value())
	}
	return width
}

func clusterWrap(c *Condition, s string, w int) string {
	width := 0
	var out strings.Builder
	g := graphemes.FromString(s)
	for g.Next() {
		cluster := g.Value()
		if strings.HasSuffix(cluster, "\n") {
			out.WriteString(cluster)
			width = 0
			continue
		}
		cw := c.graphemeWidth(cluster)
		if width+cw > w {
			out.WriteByte('\n')
			width = 0
		}
		out.WriteString(cluster)
		width += cw
	}
	return out.String()
}

// TestIsJoinerAgainstSegmentation checks the joiner table against the
// segmenter for every rune: a rune the table clears has to break on both
// sides, or the fast path would glue two clusters together.
func TestIsJoinerAgainstSegmentation(t *testing.T) {
	count := func(s string) int {
		n := 0
		g := graphemes.FromString(s)
		for g.Next() {
			n++
		}
		return n
	}
	for r := rune(0); r <= utf8.MaxRune; r++ {
		if r >= 0xD800 && r <= 0xDFFF { // not encodable
			continue
		}
		if isJoiner(r) {
			continue
		}
		if n := count("a" + string(r)); n != 2 {
			t.Fatalf("isJoiner(%#U) = false but %q is %d cluster(s)", r, "a"+string(r), n)
		}
		if n := count(string(r) + "a"); n != 2 {
			t.Fatalf("isJoiner(%#U) = false but %q is %d cluster(s)", r, string(r)+"a", n)
		}
		// Regional indicators break against a letter but pair with each
		// other, so a rune has to break against itself as well.
		if n := count(string(r) + string(r)); n != 2 {
			t.Fatalf("isJoiner(%#U) = false but %q is %d cluster(s)", r, string(r)+string(r), n)
		}
	}
}

var fastPathPieces = []string{
	"a", "Z", " ", "\t", "\n", "\r\n", "\r", "\x00", "\x7f",
	"あ", "漢", "ｱ", "。", "±", "é", "é", "́", "゛", "゙",
	"👩🏽", "👨‍👩‍👧‍👦", "🇯🇵", "🇯", "👍", "‍",
	"한", "가", "ᄀ", "ᅡ", "क्क", "︀", "󠀁",
}

func fastPathStrings(seed int64, n int) []string {
	r := rand.New(rand.NewSource(seed))
	ss := make([]string, n)
	for i := range ss {
		var b strings.Builder
		for j := r.Intn(8); j > 0; j-- {
			b.WriteString(fastPathPieces[r.Intn(len(fastPathPieces))])
		}
		ss[i] = b.String()
	}
	return ss
}

func TestStringWidthMatchesClusterLoop(t *testing.T) {
	for _, c := range []*Condition{{}, {EastAsianWidth: true, StrictEmojiNeutral: true}} {
		for _, s := range fastPathStrings(1, 20000) {
			if got, want := c.StringWidth(s), clusterWidth(c, s); got != want {
				t.Fatalf("StringWidth(%q) = %d, cluster loop = %d", s, got, want)
			}
		}
	}
}

func TestWrapMatchesClusterLoop(t *testing.T) {
	r := rand.New(rand.NewSource(2))
	for _, c := range []*Condition{{}, {EastAsianWidth: true, StrictEmojiNeutral: true}} {
		for _, s := range fastPathStrings(1, 20000) {
			w := 1 + r.Intn(6)
			if got, want := c.Wrap(s, w), clusterWrap(c, s, w); got != want {
				t.Fatalf("Wrap(%q, %d) = %q, cluster loop = %q", s, w, got, want)
			}
			for _, affix := range []string{"", "..."} {
				if got, want := c.Truncate(s, w, affix), clusterTruncate(c, s, w, affix); got != want {
					t.Fatalf("Truncate(%q, %d, %q) = %q, cluster loop = %q", s, w, affix, got, want)
				}
				if got, want := c.TruncateLeft(s, w, affix), clusterTruncateLeft(c, s, w, affix); got != want {
					t.Fatalf("TruncateLeft(%q, %d, %q) = %q, cluster loop = %q", s, w, affix, got, want)
				}
				if got, want := c.TruncatePrefix(s, w, affix), clusterTruncatePrefix(c, s, w, affix); got != want {
					t.Fatalf("TruncatePrefix(%q, %d, %q) = %q, cluster loop = %q", s, w, affix, got, want)
				}
			}
		}
	}
}

// invalidPieces mixes bytes that are not valid UTF-8 into text. A rune loop
// sees one U+FFFD per bad byte while the segmenter can gather a run of them
// into one cluster, so the fast paths have to leave these strings alone.
var invalidPieces = []string{
	"\xff", "\x80", "\xe3\x81", "\xf0\x9f", "\xc3",
	"a", "あ", "é", "\n", "👩🏽", "👨‍👩‍👧‍👦",
}

func TestInvalidUTF8MatchesClusterLoop(t *testing.T) {
	r := rand.New(rand.NewSource(9))
	for _, c := range []*Condition{{}, {EastAsianWidth: true, StrictEmojiNeutral: true}} {
		for i := 0; i < 20000; i++ {
			var b strings.Builder
			for j := 1 + r.Intn(6); j > 0; j-- {
				b.WriteString(invalidPieces[r.Intn(len(invalidPieces))])
			}
			s := b.String()
			// StringWidth answers a single byte from a shortcut of its
			// own, which has never agreed with the segmenter on a byte
			// that is not valid UTF-8.
			if len(s) > 1 {
				if got, want := c.StringWidth(s), clusterWidth(c, s); got != want {
					t.Fatalf("StringWidth(%q) = %d, cluster loop = %d", s, got, want)
				}
			}
			w := 1 + r.Intn(6)
			if got, want := c.Wrap(s, w), clusterWrap(c, s, w); got != want {
				t.Fatalf("Wrap(%q, %d) = %q, cluster loop = %q", s, w, got, want)
			}
			for _, affix := range []string{"", "..."} {
				if got, want := c.Truncate(s, w, affix), clusterTruncate(c, s, w, affix); got != want {
					t.Fatalf("Truncate(%q, %d, %q) = %q, cluster loop = %q", s, w, affix, got, want)
				}
				if got, want := c.TruncateLeft(s, w, affix), clusterTruncateLeft(c, s, w, affix); got != want {
					t.Fatalf("TruncateLeft(%q, %d, %q) = %q, cluster loop = %q", s, w, affix, got, want)
				}
				if got, want := c.TruncatePrefix(s, w, affix), clusterTruncatePrefix(c, s, w, affix); got != want {
					t.Fatalf("TruncatePrefix(%q, %d, %q) = %q, cluster loop = %q", s, w, affix, got, want)
				}
			}
		}
	}
}

// The three Truncate functions with the fast path taken out, the behaviour
// it has to reproduce exactly.
func clusterTruncate(c *Condition, s string, w int, tail string) string {
	if c.StringWidth(s) <= w {
		return s
	}
	w -= c.StringWidth(tail)
	var width int
	pos := len(s)
	g := graphemes.FromString(s)
	for g.Next() {
		chWidth := c.graphemeWidth(g.Value())
		if width+chWidth > w {
			pos = g.Start()
			break
		}
		width += chWidth
	}
	return s[:pos] + tail
}

func clusterTruncateLeft(c *Condition, s string, w int, prefix string) string {
	if c.StringWidth(s) <= w {
		return prefix
	}
	var width int
	pos := len(s)
	g := graphemes.FromString(s)
	for g.Next() {
		chWidth := c.graphemeWidth(g.Value())
		if width+chWidth > w {
			if width < w {
				pos = g.End()
				prefix += strings.Repeat(" ", width+chWidth-w)
			} else {
				pos = g.Start()
			}
			break
		}
		width += chWidth
	}
	return prefix + s[pos:]
}

func clusterTruncatePrefix(c *Condition, s string, w int, prefix string) string {
	if c.StringWidth(prefix) >= w {
		return prefix
	}
	sw := c.StringWidth(s)
	if sw <= w {
		return s
	}
	w -= c.StringWidth(prefix)
	var width int
	var pos int
	g := graphemes.FromString(s)
	for g.Next() {
		chWidth := c.graphemeWidth(g.Value())
		if sw-(width+chWidth) <= w {
			pos = g.End()
			break
		}
		width += chWidth
	}
	return prefix + s[pos:]
}

func TestTruncateMatchesClusterLoop(t *testing.T) {
	r := rand.New(rand.NewSource(4))
	for _, c := range []*Condition{{}, {EastAsianWidth: true, StrictEmojiNeutral: true}} {
		for _, s := range fastPathStrings(1, 20000) {
			w := r.Intn(8)
			for _, tail := range []string{"", "...", "…"} {
				if got, want := c.Truncate(s, w, tail), clusterTruncate(c, s, w, tail); got != want {
					t.Fatalf("Truncate(%q, %d, %q) = %q, cluster loop = %q", s, w, tail, got, want)
				}
			}
			for _, prefix := range []string{"", "...", "…"} {
				if got, want := c.TruncateLeft(s, w, prefix), clusterTruncateLeft(c, s, w, prefix); got != want {
					t.Fatalf("TruncateLeft(%q, %d, %q) = %q, cluster loop = %q", s, w, prefix, got, want)
				}
				if got, want := c.TruncatePrefix(s, w, prefix), clusterTruncatePrefix(c, s, w, prefix); got != want {
					t.Fatalf("TruncatePrefix(%q, %d, %q) = %q, cluster loop = %q", s, w, prefix, got, want)
				}
			}
		}
	}
}
