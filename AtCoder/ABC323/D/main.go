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

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	n := readInt()
	scmap := make(map[int]int)
	qu := make([]int, 0, n)
	for i := 0; i < n; i++ {
		s, c := readInt(), readInt()
		scmap[s] = c
		if c > 1 {
			qu = append(qu, s)
		}
	}

	for len(qu) > 0 {
		quNext := make([]int, 0, len(qu))

		for _, v := range qu {
			cnt := scmap[v]

			dis := 2 * v
			cnt2 := cnt / 2
			cntDis, ok := scmap[dis]
			if !ok {
				scmap[dis] = cnt2
			} else {
				scmap[dis] = cntDis + cnt2
			}
			if scmap[dis] > 1 {
				quNext = append(quNext, dis)
			}

			if cnt%2 == 1 {
				scmap[v] = 1
			} else {
				delete(scmap, v)
			}
		}
		qu = quNext
	}

	fmt.Fprint(out, len(scmap))
}

func main() {
	run(os.Stdin, os.Stdout)
}
