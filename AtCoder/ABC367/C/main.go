package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

var sc *bufio.Scanner
var memo []map[int][][]int

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	n, k := readInt(), readInt()
	rs := make([]int, n)
	for i := range rs {
		rs[i] = readInt()
	}

	memo = make([]map[int][][]int, n)
	for i := range memo {
		memo[i] = make(map[int][][]int)
	}

	os := make([]int, n)
	ans := f(0, 0, k, os, rs)

	for _, v := range ans {
		st := fmt.Sprintf("%v", v)
		fmt.Fprintln(out, st[1:len(st)-1])
	}
}

func f(i, sum, k int, os, rs []int) [][]int {
	m, ok := memo[i][sum]
	if ok {
		return m
	}

	ret := make([][]int, 0)
	for cnt := 1; cnt <= rs[i]; cnt++ {
		os[i] = cnt
		if i < len(rs)-1 {
			result := f(i+1, sum+cnt, k, os, rs)
			for _, v := range result {
				ret = append(ret, append([]int{cnt}, v...))
			}
		} else {
			if (sum+cnt)%k == 0 {
				ret = append(ret, []int{cnt})
			}
		}
	}

	memo[i][sum] = ret
	return ret
}

func main() {
	run(os.Stdin, os.Stdout)
}
