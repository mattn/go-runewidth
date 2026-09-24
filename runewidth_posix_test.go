//go:build !windows && !js && !appengine
// +build !windows,!js,!appengine

package runewidth

import (
	"os"
	"testing"
)

type envVars struct {
	lang    string
	lcall   string
	lcctype string
}

func saveEnv() envVars {
	return envVars{
		lang:    os.Getenv("LANG"),
		lcall:   os.Getenv("LC_ALL"),
		lcctype: os.Getenv("LC_CTYPE"),
	}
}
func restoreEnv(env *envVars) {
	os.Setenv("LANG", env.lang)
	os.Setenv("LC_ALL", env.lcall)
	os.Setenv("LC_CTYPE", env.lcctype)
}

func TestIsEastAsian(t *testing.T) {
	testcases := []struct {
		locale string
		want   bool
	}{
		{"foo@cjk_narrow", false},
		{"foo@cjk", false},
		{"utf-8@cjk", false},
		{"ja_JP.CP932", true},
		{"ja_JP.eucjp", true},
		{"ja_JP.EUC-JP", true},
		{"ja_JP", true},
		{"ja", true},
		{"ja_JP@cjk_narrow", false},
		{"ja_JP.UTF-8@cjk_narrow", false},
		{"zh_CN.gbk", true},
		{"zh_CN.gb18030", true},
		{"zh_CN.GB18030", true},
		{"zh_CN", true},
		{"ko_KR.euckr", true},
		{"ko_KR.EUC-KR", true},
		{"ko_KR", true},
		{"zh_TW.big5", true},
		{"zh_TW.euctw", true},
		{"zh_TW.EUC-TW", true},
		{"zh_HK.big5hkscs", true},
		{"zh_HK.Big5HKSCS", true},
		{"zh_HK.BIG5-HKSCS", true},
		{"ja_JP.Shift_JIS", true},
		{"en_US.SJIS", true},
		{"en_US", false},
		{"en_US.UTF-8", false},
		{"jv_ID.UTF-8", false},
		{"ja_JP.ISO-8859-1", false},
	}

	for _, tt := range testcases {
		got := isEastAsian(tt.locale)
		if got != tt.want {
			t.Fatalf("isEastAsian(%q) should be %v", tt.locale, tt.want)
		}
	}
}

func TestIsEastAsianLCCTYPE(t *testing.T) {
	env := saveEnv()
	defer restoreEnv(&env)
	os.Setenv("LC_ALL", "")

	testcases := []struct {
		lcctype string
		want    bool
	}{
		{"ja_JP.UTF-8", true},
		{"C", false},
		{"POSIX", false},
		{"en_US.UTF-8", false},
	}

	for _, tt := range testcases {
		os.Setenv("LC_CTYPE", tt.lcctype)
		got := IsEastAsian()
		if got != tt.want {
			t.Fatalf("IsEastAsian() for LC_CTYPE=%v should be %v", tt.lcctype, tt.want)
		}
	}
}

func TestIsEastAsianLANG(t *testing.T) {
	env := saveEnv()
	defer restoreEnv(&env)
	os.Setenv("LC_ALL", "")
	os.Setenv("LC_CTYPE", "")

	testcases := []struct {
		lcctype string
		want    bool
	}{
		{"ja_JP.UTF-8", true},
		{"C", false},
		{"POSIX", false},
		{"en_US.UTF-8", false},
		{"C.UTF-8", false},
	}

	for _, tt := range testcases {
		os.Setenv("LANG", tt.lcctype)
		got := IsEastAsian()
		if got != tt.want {
			t.Fatalf("IsEastAsian() for LANG=%v should be %v", tt.lcctype, tt.want)
		}
	}
}
