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

func run(r io.Reader) string {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	_, m := readInt(), readInt()
	as := make([][]int, m)
	for i := range as {
		k := readInt()
		as[i] = make([]int, k)
		for j := range as[i] {
			as[i][j] = readInt()
		}
	}

	amap := make(map[int]int, m)
	idxes := make([]int, m)

	for i := 0; i < m; i++ {
		set(i, as, amap, idxes)
	}

	for t, idx := range idxes {
		if idx < len(as[t]) {
			return "No"
		}
	}

	return "Yes"
}

func set(t int, as [][]int, amap map[int]int, idxes []int) {
	idx := idxes[t]
	if idx >= len(as[t]) {
		return
	}

	a := as[t][idx]
	st, ok := amap[a]
	if ok {

		idxes[t]++
		set(t, as, amap, idxes)

		idxes[st]++
		set(st, as, amap, idxes)
	} else {
		amap[a] = t
	}
}

func main() {
	fmt.Println(run(os.Stdin))
}
