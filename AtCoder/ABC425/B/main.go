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
	sc.Split(bufio.ScanWords)

	n := readInt()
	as := make([]int, n)
	for i := range as {
		as[i] = readInt()
	}

	amap := make(map[int]bool)
	for i := 1; i <= n; i++ {
		amap[i] = true
	}
	for _, a := range as {
		if a < 0 {
			continue
		}
		if !amap[a] {
			fmt.Fprint(out, "No")
			return
		}
		delete(amap, a)
	}

	as2 := make([]int, len(amap))
	idx := 0
	for a := range amap {
		as2[idx] = a
		idx++
	}
	sort.Ints(as2)

	fmt.Fprintln(out, "Yes")
	idx = 0
	for _, a := range as {
		if a < 0 {
			fmt.Fprint(out, as2[idx], " ")
			idx++
		} else {
			fmt.Fprint(out, a, " ")
		}
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
