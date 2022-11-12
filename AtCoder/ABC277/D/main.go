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

var ans int

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	n, m := readInt(), readInt()
	amaps := make(map[int]int)
	sum := 0
	for i := 0; i < n; i++ {
		a := readInt()
		sum += a
		amaps[a]++
	}

	as := make([]int, len(amaps))
	idx := 0
	for k := range amaps {
		as[idx] = k
		idx++
	}

	sort.Ints(as)

	ans = sum
	for i, a := range as {
		if i > 0 {
			if (as[i-1]+1)%m == a {
				continue
			}
		}
		dfs(a, m, sum, amaps)
	}

	fmt.Fprint(out, ans)
}

func dfs(now, m, sum int, amaps map[int]int) {
	cnt, ok := amaps[now]
	if !ok || cnt == 0 {
		return
	}

	amaps[now] = 0
	sum -= now * cnt

	if sum < ans {
		ans = sum
	}

	next := (now + 1) % m
	dfs(next, m, sum, amaps)

	amaps[now] = cnt
	sum += now * cnt

}

func main() {
	run(os.Stdin, os.Stdout)
}
