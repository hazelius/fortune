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

	n, x, y := readInt(), readInt(), readInt()
	uvmap := make(map[int][]int)
	for i := 0; i < n-1; i++ {
		u, v := readInt()-1, readInt()-1
		uvmap[u] = append(uvmap[u], v)
		uvmap[v] = append(uvmap[v], u)
	}

	used := make([]bool, n)
	ans := make([]int, n)

	used[x-1] = true
	cnt := f(x-1, y-1, 0, used, uvmap, ans)

	s := fmt.Sprintf("%v", ans[0:cnt+1])
	fmt.Fprint(out, s[1:len(s)-1])
}

func f(v, goal, count int, used []bool, uvmap map[int][]int, ans []int) int {
	ans[count] = v + 1
	count++
	for _, p := range uvmap[v] {
		if used[p] {
			continue
		}
		if p == goal {
			ans[count] = p + 1
			return count
		}

		used[p] = true
		tmp := f(p, goal, count, used, uvmap, ans)
		if tmp > 0 {
			return tmp
		}
		used[p] = false
	}

	return -1
}

func main() {
	run(os.Stdin, os.Stdout)
}
