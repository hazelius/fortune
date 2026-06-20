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
	lrs := make([][]int, n)
	for i := range lrs {
		lrs[i] = []int{readInt(), readInt()}
	}

	sort.Slice(lrs, func(i, j int) bool {
		return lrs[i][1] < lrs[j][1]
	})

	ans := sort.Search(999999999, func(x int) bool {
		cnt := 1
		pre := lrs[0]
		for i, lr := range lrs {
			if i == 0 {
				continue
			}
			if pre[1]+x < lr[0] {
				pre = lr
				cnt++
			}
		}
		return cnt < k
	})

	if ans <= 0 {
		ans = -1
	}

	fmt.Fprint(out, ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
