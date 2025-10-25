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

	n, m, c := readInt(), readInt(), readInt()
	amap := make(map[int]int)
	amap[0] = 0
	for i := 0; i < n; i++ {
		a := readInt()
		amap[a]++
	}

	ps := make([]int, len(amap))
	idx := 0
	for k := range amap {
		ps[idx] = k
		idx++
	}
	sort.Ints(ps)
	ps2 := make([]int, len(ps)*2)
	for i := range ps2 {
		ps2[i] = ps[i%len(ps)]
		if i >= len(ps) {
			ps2[i] += m
		}
	}

	ans := 0
	start := 0
	sv := 0
	for i, a := range ps2 {
		if i == 0 {
			continue
		}
		sv += amap[a%m]
		for sv >= c {
			ans += sv * (ps2[start+1] - ps2[start])
			start++
			if ps2[start] >= m {
				break
			}
			sv -= amap[ps2[start]]
		}
		if ps2[start] >= m {
			break
		}
	}

	fmt.Fprint(out, ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
