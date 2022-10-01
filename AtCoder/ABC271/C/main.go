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

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	n := readInt()
	as := make([]int, n)
	for i := range as {
		as[i] = readInt()
	}
	mapa := make(map[int]int)
	for _, a := range as {
		mapa[a]++
	}
	dup := 0
	as = make([]int, len(mapa))
	idx := 0
	for a, cnt := range mapa {
		dup += cnt - 1
		as[idx] = a
		idx++
	}

	sort.Ints(as)

	ans := 0
	idx = 0
	last := len(as) - 1
	for {
		if idx <= last {
			a := as[idx]
			if a == ans {
				idx++
				continue
			}
			if a == ans+1 {
				ans = a
				idx++
				continue
			}
		}
		if dup > 1 {
			dup -= 2
			ans++
			continue
		}

		sac := 2
		if dup == 1 {
			sac--
			dup--
		}

		if last+1 < idx+sac {
			break
		}
		last -= sac
		ans++
	}

	fmt.Fprint(out, ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
