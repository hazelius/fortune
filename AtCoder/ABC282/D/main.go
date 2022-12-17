package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

var sc *bufio.Scanner
var bmap map[int]bool

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	n, m := readInt(), readInt()
	uvmap := make(map[int][]int)
	for i := 0; i < m; i++ {
		u, v := readInt(), readInt()
		uvmap[u] = append(uvmap[u], v)
		uvmap[v] = append(uvmap[v], u)
	}

	ans := (n * (n - 1) / 2) - m
	bmap = make(map[int]bool)
	for i := range uvmap {
		_, ok := bmap[i]
		if ok {
			continue
		}
		bmap[i] = true

		cnt1, cnt2, ok := dfs(i, false, uvmap)
		cnt1++
		if !ok {
			fmt.Fprint(out, 0)
			return
		}

		ans -= cnt1 * (cnt1 - 1) / 2
		ans -= cnt2 * (cnt2 - 1) / 2
	}

	fmt.Fprint(out, ans)
}

func dfs(i int, cl bool, uvmap map[int][]int) (int, int, bool) {
	uvs, ok := uvmap[i]
	if !ok {
		return 0, 0, true
	}

	cnt1, cnt2 := 0, 0
	for _, k := range uvs {
		vc, ok := bmap[k]
		if ok {
			if vc != cl {
				return cnt1, cnt2, false
			}
		} else {
			bmap[k] = cl
			if cl {
				cnt1++
			} else {
				cnt2++
			}
			c1, c2, ok := dfs(k, !cl, uvmap)
			if !ok {
				return cnt1, cnt2, ok
			}
			cnt1 += c1
			cnt2 += c2
		}
	}
	return cnt1, cnt2, true
}

func main() {
	run(os.Stdin, os.Stdout)
}
