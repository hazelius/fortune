package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
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

func run(r io.Reader) string {
	sc = bufio.NewScanner(r)
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := readInt()
	xys := make([][]int, n)
	for i := range xys {
		xys[i] = []int{readInt(), readInt()}
	}
	s := readString()

	ymap := make(map[int][][]int)
	for i, xy := range xys {
		x, y := xy[0], xy[1]
		c := s[i]
		ym, ok := ymap[y]
		if !ok {
			ym = make([][]int, 2)
		}
		if c == 'R' {
			ym[0] = append(ym[0], x)
		} else {
			ym[1] = append(ym[1], x)
		}
		ymap[y] = ym
	}

	for _, ym := range ymap {
		rn, ln := ym[0], ym[1]
		if len(rn) > 0 && len(ln) > 0 {
			sort.Ints(rn)
			sort.Sort(sort.Reverse(sort.IntSlice(ln)))
			if rn[0] < ln[0] {
				return "Yes"
			}
		}
	}

	return "No"
}

func main() {
	fmt.Println(run(os.Stdin))
}
