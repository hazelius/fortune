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

func run(r io.Reader) string {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	n, m := readInt(), readInt()
	nmap := make(map[int][]int)
	for i := 0; i < m; i++ {
		a, b := readInt()-1, readInt()-1
		nmap[a] = append(nmap[a], b)
		nmap[b] = append(nmap[b], a)
	}

	used := make([]bool, n)
	for i := 0; i < n; i++ {
		if used[i] {
			continue
		}
		if !f(-1, i, nmap, used) {
			return "No"
		}
	}

	return "Yes"
}

func f(pre, i int, nmap map[int][]int, used []bool) bool {
	if used[i] {
		return false
	}
	if len(nmap[i]) > 2 {
		return false
	}

	used[i] = true
	for _, l := range nmap[i] {
		if l != pre {
			if !f(i, l, nmap, used) {
				return false
			}
		}
	}
	return true
}

func main() {
	fmt.Println(run(os.Stdin))
}
