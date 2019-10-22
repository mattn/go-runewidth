package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
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
		if i%3 == 2 {
			if i < len(arr)-1 {
				fmt.Fprint(out, "\n\t")
			}
		} else {
			fmt.Fprint(out, " ")
		}
	}
	fmt.Fprintln(out, "\n}")
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
		n, err := fmt.Sscanf(line, "%x..%x;%s ", &r1, &r2, &ss)
		if err != nil || n == 2 {
			n, err = fmt.Sscanf(line, "%x;%s ", &r1, &ss)
			if err != nil || n != 2 {
				continue
			}
			r2 = r1
		}

		if strings.Index(line, "COMBINING") != -1 {
			cmb = append(cmb, rrange{
				lo: r1,
				hi: r2,
			})
		}

		switch ss {
		case "W":
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
			if r1 > 0xFF {
				na = append(na, rrange{
					lo: r1,
					hi: r2,
				})
			}
		case "N":
			nu = append(nu, rrange{
				lo: r1,
				hi: r2,
			})
		}
	}

	for i := 0; i < len(cmb)-1; i++ {
		if cmb[i].hi+1 == cmb[i+1].lo {
			lo := cmb[i].lo
			cmb = append(cmb[:i], cmb[i+1:]...)
			cmb[i].lo = lo
			i--
		}
	}
	for i := 0; i < len(dbl)-1; i++ {
		if dbl[i].hi+1 == dbl[i+1].lo {
			lo := dbl[i].lo
			dbl = append(dbl[:i], dbl[i+1:]...)
			dbl[i].lo = lo
			i--
		}
	}
	for i := 0; i < len(amb)-1; i++ {
		if amb[i].hi+1 == amb[i+1].lo {
			lo := amb[i].lo
			amb = append(amb[:i], amb[i+1:]...)
			amb[i].lo = lo
			i--
		}
	}
	for i := 0; i < len(na)-1; i++ {
		if na[i].hi+1 == na[i+1].lo {
			lo := na[i].lo
			na = append(na[:i], na[i+1:]...)
			na[i].lo = lo
			i--
		}
	}
	for i := 0; i < len(nu)-1; i++ {
		if nu[i].hi+1 == nu[i+1].lo {
			lo := nu[i].lo
			nu = append(nu[:i], nu[i+1:]...)
			nu[i].lo = lo
			i--
		}
	}

	generate(out, "combining", cmb)
	fmt.Fprintln(out)
	generate(out, "doublewidth", dbl)
	fmt.Fprintln(out)
	generate(out, "ambiguous", amb)
	fmt.Fprint(out)
	generate(out, "notassigned", na)
	fmt.Fprintln(out)
	generate(out, "neutral", nu)
	fmt.Fprintln(out)

	return nil
}

func emoji(out io.Writer, in io.Reader) error {
	scanner := bufio.NewScanner(in)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Index(line, "Extended_Pictographic ; No") != -1 {
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
	for i := 0; i < len(arr)-1; i++ {
		if arr[i].hi+1 == arr[i+1].lo {
			lo := arr[i].lo
			arr = append(arr[:i], arr[i+1:]...)
			arr[i].lo = lo
			i--
		}
	}

	generate(out, "emoji", arr)

	return nil
}

func main() {
	f, err := os.Create("runewidth_table.go")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	fmt.Fprint(f, "package runewidth\n\n")

	resp, err := http.Get("https://unicode.org/Public/12.1.0/ucd/EastAsianWidth.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	eastasian(f, resp.Body)

	resp, err = http.Get("https://unicode.org/Public/emoji/12.1/emoji-data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	emoji(f, resp.Body)
}
