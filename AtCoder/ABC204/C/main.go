package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(n, m int, ab [][]int) int {
	mab := make(map[int][]int)
	for _, v := range ab {
		a, b := v[0], v[1]
		if bs, ok := mab[a]; ok {
			bs = append(bs, b)
			mab[a] = bs
		} else {
			mab[a] = []int{b}
		}
	}
	ans := 0
	for i := 1; i <= n; i++ {
		rt := make(map[int]bool)
		f(rt, i, mab)
		ans += len(rt)
	}

	return ans
}

func f(rt map[int]bool, s int, mab map[int][]int) {
	if _, ok := rt[s]; ok {
		return
	}

	rt[s] = true

	bs, ok := mab[s]
	if !ok {
		return
	}

	for _, b := range bs {
		f(rt, b, mab)
	}
}

func main() {
	sc.Split(bufio.ScanWords)
	n, m := readInt(), readInt()
	ab := make([][]int, m)
	for i := range ab {
		ab[i] = []int{readInt(), readInt()}
	}
	fmt.Println(run(n, m, ab))
}
