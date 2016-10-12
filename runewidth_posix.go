// +build !windows,!js

package runewidth

import (
	"os"
	"regexp"
	"strings"
)

var reLoc = regexp.MustCompile(`^[a-z][a-z][a-z]?(?:_[A-Z][A-Z])?\.(.+)`)

func isEastAsian(locale string) int {
	charset := strings.ToLower(locale)
	r := reLoc.FindStringSubmatch(locale)
	if len(r) == 2 {
		charset = strings.ToLower(r[1])
	}

	if strings.HasSuffix(charset, "@cjk_narrow") {
		return false
	}

	for pos, b := range []byte(charset) {
		if b == '@' {
			charset = charset[:pos]
			break
		}
	}
	max := 1
	switch charset {
	case "utf-8", "utf8":
		max = 6
	case "jis":
		max = 8
	case "eucjp":
		max = 3
	case "euckr", "euccn":
		max = 2
	case "sjis", "cp932", "cp51932", "cp936", "cp949", "cp950":
		max = 2
	case "big5":
		max = 2
	case "gbk", "gb2312":
		max = 2
	}

	if max > 1 && (charset[0] != 'u' ||
		strings.HasPrefix(locale, "ja") ||
		strings.HasPrefix(locale, "ko") ||
		strings.HasPrefix(locale, "zh")) {
		return true
	}
	return false
}

// IsEastAsian return true if the current locale is CJK
func IsEastAsian() bool {
	locale := os.Getenv("LC_CTYPE")
	if locale == "" {
		locale = os.Getenv("LANG")
	}

	// ignore C locale
	if locale == "POSIX" || locale == "C" {
		return false
	}
	if len(locale) > 1 && locale[0] == 'C' && (locale[1] == '.' || locale[1] == '-') {
		return false
	}

	return isEastAsian(locale)
}
