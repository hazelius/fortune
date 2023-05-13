package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var sc *bufio.Scanner

func readString() string {
	sc.Scan()
	return sc.Text()
}
func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	s, t := readString(), readString()
	smap := make(map[rune]int)
	tmap := make(map[rune]int)
	allkeys := make(map[rune]bool)
	for _, c := range s {
		smap[c]++
		allkeys[c] = true
	}
	for _, c := range t {
		tmap[c]++
		allkeys[c] = true
	}

	for k := range allkeys {
		if k == '@' {
			continue
		}
		cnt := smap[k]

		tcnt, ok := tmap[k]
		if ok {
			cnt -= tcnt
		}
		if cnt == 0 {
			continue
		}

		switch k {
		case 'a', 't', 'c', 'o', 'd', 'e', 'r':
		default:
			fmt.Fprint(out, "No")
			return
		}

		if cnt > 0 {
			atcnt, ok := tmap['@']
			if !ok || atcnt < cnt {
				fmt.Fprint(out, "No")
				return
			}
			tmap['@'] = atcnt - cnt
		} else {
			cnt *= -1
			atcnt, ok := smap['@']
			if !ok || atcnt < cnt {
				fmt.Fprint(out, "No")
				return
			}
			smap['@'] = atcnt - cnt
		}
	}

	fmt.Fprint(out, "Yes")
}

func main() {
	run(os.Stdin, os.Stdout)
}
