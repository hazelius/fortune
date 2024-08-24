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

	n, k := readInt(), readInt()
	abmap := make(map[int]map[int]bool)
	for i := 1; i < n; i++ {
		a, b := readInt(), readInt()
		if _, ok := abmap[a]; !ok {
			abmap[a] = make(map[int]bool)
		}
		if _, ok := abmap[b]; !ok {
			abmap[b] = make(map[int]bool)
		}
		abmap[a][b] = true
		abmap[b][a] = true
	}

	mapvs := make(map[int]bool)
	for i := 0; i < k; i++ {
		v := readInt()
		mapvs[v] = true
	}

	ans := n
	for a := range abmap {
		ans = delA(a, ans, mapvs, abmap)
	}

	fmt.Fprint(out, ans)
}

func delA(a, ans int, mapvs map[int]bool, abmap map[int]map[int]bool) int {
	if mapvs[a] {
		return ans
	}
	bs := abmap[a]
	if len(bs) > 1 {
		return ans
	}
	delete(abmap, a)
	ans--

	b := 0
	for k := range bs {
		b = k
		break
	}
	delete(abmap[b], a)

	return delA(b, ans, mapvs, abmap)
}

func main() {
	run(os.Stdin, os.Stdout)
}
