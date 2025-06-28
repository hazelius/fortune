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

	t := readInt()
	for i := 0; i < t; i++ {
		n := readInt()
		s1 := readInt()
		as := make([]int, n-2)
		for i := range as {
			as[i] = readInt()
		}
		sn := readInt()

		if s1*2 >= sn {
			fmt.Fprintln(out, 2)
			continue
		}

		sort.Ints(as)
		ans := 0
		tmp := s1
		preIdx := -1
		for ; ans < n; ans++ {
			idx := sort.Search(n-2, func(j int) bool {
				return tmp*2 < as[j]
			})
			idx--
			if preIdx == idx {
				ans = -1
				break
			}
			tmp = as[idx]
			preIdx = idx
			if tmp*2 >= sn {
				ans += 3
				break
			}
		}
		fmt.Fprintln(out, ans)
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
