package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

var sc *bufio.Scanner

func readString() string {
	sc.Scan()
	return sc.Text()
}

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	buf := make([]byte, 1<<20)
	sc.Buffer(buf, len(buf))
	sc.Split(bufio.ScanWords)

	q := readInt()
	s := readString()
	t := readString()

	hits := make([]int, 0)
	for i := 0; i >= 0; i++ {
		idx := strings.Index(s[i:], t)
		if idx < 0 {
			break
		}
		i += idx
		hits = append(hits, i)
	}

	for i := 0; i < q; i++ {
		l, r := readInt()-1, readInt()
		if r-l < len(t) {
			fmt.Fprintln(out, "No")
			continue
		}

		v := sort.Search(len(hits), func(j int) bool {
			return l <= hits[j]
		})
		if v < len(hits) {
			idx := hits[v]
			if idx+len(t) <= r {
				fmt.Fprintln(out, "Yes")
				continue
			}
		}
		fmt.Fprintln(out, "No")
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
