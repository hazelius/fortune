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
	buf := make([]byte, 1<<20)
	sc.Buffer(buf, len(buf))
	sc.Split(bufio.ScanWords)

	h, w := readInt(), readInt()
	ss := make([]string, h)
	for i := 0; i < h; i++ {
		ss[i] = readString()
	}
	ans := 0
	for h1 := 0; h1 < h; h1++ {
		for h2 := h1; h2 < h; h2++ {
			for w1 := 0; w1 < w; w1++ {
				for w2 := w1; w2 < w; w2++ {
					flg := true
					for i := h1; i <= h2; i++ {
						for j := w1; j <= w2; j++ {
							if ss[i][j] != ss[h1+h2-i][w1+w2-j] {
								flg = false
								break
							}
						}
						if !flg {
							break
						}
					}
					if flg {
						ans++
					}
				}
			}
		}
	}
	fmt.Fprint(out, ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
