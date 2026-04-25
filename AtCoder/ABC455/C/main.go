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
	buf := make([]byte, 1<<20)
	sc.Buffer(buf, len(buf))
	sc.Split(bufio.ScanWords)

	n, k := readInt(), readInt()
	amap := make(map[int]int)
	for i := 0; i < n; i++ {
		a := readInt()
		amap[a] += a
	}
	ans := 0
	as := make([]int, len(amap))
	idx := 0
	for _, v := range amap {
		as[idx] = v
		ans += v
		idx++
	}
	sort.Sort(sort.Reverse(sort.IntSlice(as)))
	for i := 0; i < k && i < len(as); i++ {
		ans -= as[i]
	}
	fmt.Fprint(out, ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
