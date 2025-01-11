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

type person struct {
	name  string
	point int
}

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	pmap := make(map[byte]int)
	s := "ABCDE"
	for i := range s {
		pmap[s[i]] = readInt()
	}
	ans := make([]person, 31)
	for i := range ans {
		for j := 0; j < 5; j++ {
			if (1<<j)&(i+1) > 0 {
				ans[i].name += s[j : j+1]
				ans[i].point += pmap[s[j]]
			}
		}
	}

	sort.Slice(ans, func(a, b int) bool {
		if ans[a].point > ans[b].point {
			return true
		}
		if ans[a].point == ans[b].point {
			return ans[a].name < ans[b].name
		}
		return false
	})

	for _, p := range ans {
		fmt.Fprintln(out, p.name)
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
