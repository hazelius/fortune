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

	n, m := readInt(), readInt()
	if n != m+1 {
		fmt.Fprint(out, "No")
		return
	}
	uvmap := make(map[int][]int)
	for i := 0; i < m; i++ {
		u, v := readInt(), readInt()
		uvmap[u] = append(uvmap[u], v)
		uvmap[v] = append(uvmap[v], u)
	}

	used := make(map[int]bool)
	for u, vs := range uvmap {
		used[u] = true
		for _, v := range vs {
			if !f(v, uvmap, used) {
				fmt.Fprint(out, "No")
				return
			}
		}
		break
	}

	fmt.Fprint(out, "Yes")
}

func f(u int, uvmap map[int][]int, used map[int]bool) bool {
	if used[u] {
		return false
	}
	used[u] = true
	vs := uvmap[u]
	if len(vs) > 2 {
		return false
	}
	if len(vs) == 1 {
		return true
	}
	ret := false
	for _, v := range vs {
		if f(v, uvmap, used) {
			ret = true
		}
	}
	return ret
}

func main() {
	run(os.Stdin, os.Stdout)
}
