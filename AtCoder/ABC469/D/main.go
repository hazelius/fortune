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

type pair struct {
	a int
	b int
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	n, m := readInt(), readInt()
	f1, f2 := readInt(), readInt()
	pairMap := make(map[pair]bool)

	for i := 0; i < m-1; i++ {
		a, b := readInt(), readInt()
		if f1 > 0 && f1 != a && f1 != b {
			p := pair{f1, a}
			if p.a > p.b {
				p = pair{p.b, p.a}
			}
			pairMap[p] = true
			p = pair{f1, b}
			if p.a > p.b {
				p = pair{p.b, p.a}
			}
			pairMap[p] = true
			f1 = -1
		}
		if f2 > 0 && f2 != a && f2 != b {
			p := pair{f2, a}
			if p.a > p.b {
				p = pair{p.b, p.a}
			}
			pairMap[p] = true
			p = pair{f2, b}
			if p.a > p.b {
				p = pair{p.b, p.a}
			}
			pairMap[p] = true
			f2 = -1
		}
		for v := range pairMap {
			if v.a != a && v.a != b && v.b != b && v.b != a {
				delete(pairMap, v)
			}
		}
	}

	ans := 0
	if f1 != -1 {
		ans += n - 1
		for v := range pairMap {
			if v.a == f1 || v.b == f1 {
				delete(pairMap, v)
			}
		}
	}
	if f2 != -1 {
		ans += n - 1
		for v := range pairMap {
			if v.a == f2 || v.b == f2 {
				delete(pairMap, v)
			}
		}
	}
	if f1 != -1 && f2 != -1 {
		ans -= 1
	}
	ans += len(pairMap)

	fmt.Fprint(out, ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
