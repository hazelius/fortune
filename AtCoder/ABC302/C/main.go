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

func readString() string {
	sc.Scan()
	return sc.Text()
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	n, _ := readInt(), readInt()
	ss := make([]string, n)
	for i := range ss {
		ss[i] = readString()
	}

	dd := make([][]int, n)
	for i := range dd {
		dd[i] = make([]int, 0)
	}

	for i := 0; i < n-1; i++ {
		s1 := ss[i]
		for j := i + 1; j < n; j++ {
			s2 := ss[j]
			cnt := 0
			for k := range s1 {
				if s1[k] != s2[k] {
					cnt++
				}
			}
			if cnt == 1 {
				dd[i] = append(dd[i], j)
				dd[j] = append(dd[j], i)
			}
		}
	}

	for i := 0; i < n; i++ {
		used := make(map[int]bool)
		f(i, dd, used)
		if len(used) == n {
			fmt.Fprint(out, "Yes")
			return
		}
	}

	fmt.Fprint(out, "No")
}

func f(i int, dd [][]int, used map[int]bool) bool {
	for _, d := range dd[i] {
		if used[d] {
			continue
		}
		used[d] = true
		if f(d, dd, used) {
			return true
		}
		used[d] = false
	}
	return false
}

func main() {
	run(os.Stdin, os.Stdout)
}
