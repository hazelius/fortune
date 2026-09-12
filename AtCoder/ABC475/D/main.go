package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

var sc *bufio.Scanner

func readString() string {
	sc.Scan()
	return sc.Text()
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	buf := make([]byte, 1<<20)
	sc.Buffer(buf, len(buf))
	sc.Split(bufio.ScanWords)

	s := readString()
	keta := 1
	for i := 1; i <= len(s); i++ {
		keta *= 10
	}
	sosu := make([]bool, keta)

	for i := 2; i < keta; i++ {
		if sosu[i] {
			continue
		}

		stri := strconv.Itoa(i)

		if len(stri) == len(s) {
			cmap := make(map[rune]byte)
			cmap2 := make(map[byte]rune)
			flg := true
			for j, c := range s {
				c2 := stri[j]
				v, ok := cmap[c]
				if ok && v != c2 {
					flg = false
					break
				}
				cmap[c] = c2

				v2, ok := cmap2[c2]
				if ok && v2 != c {
					flg = false
					break
				}
				cmap2[c2] = c
			}

			if flg {
				fmt.Fprint(out, i)
				return
			}
		}

		for j := 2; i*j < keta; j++ {
			sosu[i*j] = true
		}
	}

	fmt.Fprint(out, -1)
}

func main() {
	run(os.Stdin, os.Stdout)
}
