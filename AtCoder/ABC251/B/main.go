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

func run(r io.Reader) int {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	n, w := readInt(), readInt()
	as := make([]int, n)
	for i := range as {
		as[i] = readInt()
	}

	sort.Sort(sort.Reverse(sort.IntSlice(as)))
	ans := make(map[int]bool)
	for i, a := range as {
		if a > w {
			continue
		}
		ans[a] = true
		for j := i + 1; j < n; j++ {
			if a+as[j] > w {
				continue
			}
			ans[a+as[j]] = true
			for k := j + 1; k < n; k++ {
				if a+as[j]+as[k] > w {
					continue
				}
				ans[a+as[j]+as[k]] = true
			}
		}
	}

	return len(ans)
}

func main() {
	fmt.Println(run(os.Stdin))
}
