//go:build ignore
// +build ignore

// Generate runewidth_table.go from data at https://unicode.org/

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"go/format"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strings"
)

type rrange struct {
	lo rune
	hi rune
}

func generate(out io.Writer, v string, arr []rrange) {
	fmt.Fprintf(out, "var %s = table{\n\t", v)
	for i := 0; i < len(arr); i++ {
		fmt.Fprintf(out, "{0x%04X, 0x%04X},", arr[i].lo, arr[i].hi)
		if i < len(arr)-1 {
			if i%3 == 2 {
				fmt.Fprint(out, "\n\t")
			} else {
				fmt.Fprint(out, " ")
			}
		}
	}
	fmt.Fprintln(out, "\n}")
}

func shapeup(p *[]rrange) {
	arr := *p
	for i := 0; i < len(arr)-1; i++ {
		if arr[i].hi+1 == arr[i+1].lo {
			lo := arr[i].lo
			arr = append(arr[:i], arr[i+1:]...)
			arr[i].lo = lo
			i--
		}
	}
	*p = arr
}

func eastasian(out io.Writer, in io.Reader) error {
	scanner := bufio.NewScanner(in)

	dbl := []rrange{}
	amb := []rrange{}
	cmb := []rrange{}
	na := []rrange{}
	nu := []rrange{}
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#") {
			continue
		}
		var r1, r2 rune
		var ss string
		n, err := fmt.Sscanf(line, "%x..%x ; %s", &r1, &r2, &ss)
		if err != nil || n == 2 {
			n, err = fmt.Sscanf(line, "%x ; %s", &r1, &ss)
			if err != nil || n != 2 {
				continue
			}
			r2 = r1
		}

		// Nonspacing (Mn) and enclosing (Me) marks are zero width.
		// Spacing marks (Mc) occupy their own cell, like glibc's
		// wcwidth, so they are excluded here.
		isCombo := strings.Index(line, "COMBINING") != -1 ||
			strings.Contains(line, "# Mn") ||
			strings.Contains(line, "# Me")
		isVS := strings.Contains(line, "VARIATION SELECTOR")

		if isCombo || isVS {
			cmb = append(cmb, rrange{lo: r1, hi: r2})
		}
		if isVS {
			continue
		}

		switch ss {
		case "W", "F":
			dbl = append(dbl, rrange{
				lo: r1,
				hi: r2,
			})
		case "A":
			amb = append(amb, rrange{
				lo: r1,
				hi: r2,
			})
		case "Na":
			na = append(na, rrange{
				lo: r1,
				hi: r2,
			})
		case "N":
			nu = append(nu, rrange{
				lo: r1,
				hi: r2,
			})
		}
	}

	shapeup(&cmb)
	generate(out, "combining", cmb)
	fmt.Fprintln(out)

	shapeup(&dbl)
	generate(out, "doublewidth", dbl)
	fmt.Fprintln(out)

	shapeup(&amb)
	generate(out, "ambiguous", amb)
	fmt.Fprint(out)

	shapeup(&na)
	generate(out, "narrow", na)
	fmt.Fprintln(out)

	shapeup(&nu)
	generate(out, "neutral", nu)
	fmt.Fprintln(out)

	return nil
}

// joiner emits the runes that can join with a neighbour to form a
// multi-rune grapheme cluster: the GraphemeBreakProperty classes Extend,
// ZWJ, SpacingMark, Prepend and Regional_Indicator, plus the Hangul classes
// L, V, T, LV and LVT. CR is left out because it is below 0x300, which
// isJoiner tests on its own.
//
// It takes every version it is given and emits the union, because the table
// is only safe when it covers what the segmenter package joins as well as
// what the current data says: a rune the segmenter joins but the table
// clears would send a multi-rune cluster down the fast path, while the
// other direction only costs a needless trip through the segmenter.
// TestIsJoinerAgainstSegmentation checks that the union is wide enough.
func joiner(out io.Writer, in ...io.Reader) error {
	joining := map[string]bool{
		"Extend": true, "ZWJ": true, "SpacingMark": true, "Prepend": true,
		"Regional_Indicator": true,
		"L":                  true, "V": true, "T": true, "LV": true, "LVT": true,
	}

	arr := []rrange{}
	for _, r := range in {
		scanner := bufio.NewScanner(r)
		for scanner.Scan() {
			line := scanner.Text()
			if i := strings.IndexByte(line, '#'); i >= 0 {
				line = line[:i]
			}
			field := strings.Split(line, ";")
			if len(field) != 2 {
				continue
			}
			if !joining[strings.TrimSpace(field[1])] {
				continue
			}
			var r1, r2 rune
			n, err := fmt.Sscanf(field[0], "%x..%x", &r1, &r2)
			if err != nil || n == 1 {
				n, err = fmt.Sscanf(field[0], "%x", &r1)
				if err != nil || n != 1 {
					continue
				}
				r2 = r1
			}
			if r2 < 0x300 {
				continue
			}
			arr = append(arr, rrange{lo: r1, hi: r2})
		}
		if scanner.Err() != nil {
			return scanner.Err()
		}
	}

	coalesce(&arr)
	generate(out, "joiner", arr)

	return nil
}

// coalesce sorts the ranges and merges the ones that overlap or touch,
// which shapeup does not do because the tables it is used on come already
// sorted and disjoint.
func coalesce(p *[]rrange) {
	arr := *p
	sort.Slice(arr, func(i, j int) bool { return arr[i].lo < arr[j].lo })
	out := arr[:0]
	for _, r := range arr {
		if len(out) > 0 && r.lo <= out[len(out)-1].hi+1 {
			if r.hi > out[len(out)-1].hi {
				out[len(out)-1].hi = r.hi
			}
			continue
		}
		out = append(out, r)
	}
	*p = out
}

func emoji(out io.Writer, in io.Reader) error {
	scanner := bufio.NewScanner(in)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Index(line, "Extended_Pictographic=No") != -1 {
			break
		}
	}

	if scanner.Err() != nil {
		return scanner.Err()
	}

	arr := []rrange{}
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#") {
			continue
		}
		var r1, r2 rune
		n, err := fmt.Sscanf(line, "%x..%x ", &r1, &r2)
		if err != nil || n == 1 {
			n, err = fmt.Sscanf(line, "%x ", &r1)
			if err != nil || n != 1 {
				continue
			}
			r2 = r1
		}
		if r2 < 0xFF {
			continue
		}

		arr = append(arr, rrange{
			lo: r1,
			hi: r2,
		})
	}

	shapeup(&arr)
	generate(out, "emoji", arr)

	return nil
}

func main() {
	var buf bytes.Buffer
	f := &buf
	fmt.Fprint(f, "// Code generated by script/generate.go. DO NOT EDIT.\n\n")

	fmt.Fprint(f, "package runewidth\n\n")

	resp, err := http.Get("https://unicode.org/Public/17.0.0/ucd/EastAsianWidth.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	eastasian(f, resp.Body)

	resp, err = http.Get("https://unicode.org/Public/17.0.0/ucd/emoji/emoji-data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	emoji(f, resp.Body)

	// 15.0.0 is the version github.com/clipperhouse/uax29 built its own
	// tables from, see the header of its graphemes/trie.go.
	var gbp []io.Reader
	for _, u := range []string{
		"https://unicode.org/Public/17.0.0/ucd/auxiliary/GraphemeBreakProperty.txt",
		"https://unicode.org/Public/15.0.0/ucd/auxiliary/GraphemeBreakProperty.txt",
	} {
		resp, err = http.Get(u)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		gbp = append(gbp, resp.Body)
	}

	fmt.Fprintln(f)
	joiner(f, gbp...)

	out, err := format.Source(f.Bytes())
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile("runewidth_table.go", out, 0666)
	if err != nil {
		log.Fatal(err)
	}
}
