package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

var sc *bufio.Scanner

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func readString() string {
	sc.Scan()
	return sc.Text()
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, k := readInt(), readInt()
	s := readString()
	bs := make([]byte, n)
	maxcnt := 0
	hcnt := 0
	for i := range bs {
		if s[i] == 'o' {
			k--
			bs[i] = 'o'
			hcnt = 0
		} else if s[i] == '.' {
			bs[i] = '.'
			maxcnt += hcnt / 2
			if hcnt%2 > 0 {
				maxcnt++
				for j := 0; j < hcnt; j++ {
					bs[i-j-1] = '!'
				}
			}
			hcnt = 0
		} else {
			if i > 0 && s[i-1] == 'o' || i < n-1 && s[i+1] == 'o' {
				bs[i] = '.'
				maxcnt += hcnt / 2
				if hcnt%2 > 0 {
					maxcnt++
					for j := 0; j < hcnt; j++ {
						bs[i-j-1] = '!'
					}
				}
				hcnt = 0
			} else {
				bs[i] = '?'
				hcnt++
			}
		}
	}
	maxcnt += hcnt / 2
	if hcnt%2 > 0 {
		maxcnt++
		for j := 0; j < hcnt; j++ {
			bs[len(bs)-j-1] = '!'
		}
	}

	flg := true
	for i, v := range bs {
		if k == 0 {
			if v != 'o' {
				bs[i] = '.'
			}
			continue
		}
		if v != '!' {
			flg = true
			continue
		}

		if maxcnt > k {
			bs[i] = '?'
			continue
		}

		if flg {
			bs[i] = 'o'
			flg = false
		} else {
			bs[i] = '.'
			flg = true
		}
	}

	fmt.Fprint(out, string(bs))
}

func main() {
	run(os.Stdin, os.Stdout)
}
