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

	readInt()
	s := readString()

	to := make(map[string]string)
	fr := make(map[string][]string)
	q := readInt()
	for i := 0; i < q; i++ {
		c, d := readString(), readString()
		if c == d {
			continue
		}
		fvals, ok := fr[c]
		if ok {
			for _, fval := range fvals {
				to[fval] = d
				fr[d] = append(fr[d], fval)
			}
			delete(fr, c)
		}
		if _, ok := to[c]; !ok {
			to[c] = d
			fr[d] = append(fr[d], c)
		}
	}

	ans := ""
	for i := range s {
		tmp := s[i : i+1]
		v, ok := to[tmp]
		if ok {
			ans += v
		} else {
			ans += tmp
		}
	}
	fmt.Fprint(out, ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
