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

	h, w, k := readInt(), readInt(), readInt()
	ss := make([]string, h)
	for i := range ss {
		ss[i] = readString()
	}

	ans := -1
	for _, s := range ss {
		d, m := 0, 0
		for i, c := range s {
			if c == 'x' {
				d, m = 0, 0
				continue
			}
			if c == 'o' {
				m++
			} else {
				d++
			}
			sum := d + m
			if sum < k {
				continue
			}

			if sum > k {
				switch s[i-k] {
				case 'o':
					m--
				case '.':
					d--
				}
			}
			if ans < 0 || ans > d {
				ans = d
			}
		}
	}

	for i := 0; i < w; i++ {
		d, m := 0, 0
		for j := 0; j < h; j++ {
			c := ss[j][i]
			if c == 'x' {
				d, m = 0, 0
				continue
			}
			if c == 'o' {
				m++
			} else {
				d++
			}
			sum := d + m
			if sum < k {
				continue
			}

			if sum > k {
				switch ss[j-k][i] {
				case 'o':
					m--
				case '.':
					d--
				}
			}
			if ans < 0 || ans > d {
				ans = d
			}
		}
	}

	fmt.Fprint(out, ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
