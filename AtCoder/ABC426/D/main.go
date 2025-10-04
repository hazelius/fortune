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
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	t := readInt()
	for i := 0; i < t; i++ {
		n := readInt()
		s := readString()
		renzoku := make([]int, 2)
		cost := make([]int, 2)
		cnt := 0
		for idx, c := range s {
			val := 0
			if c == '1' {
				val = 1
			}

			cnt++
			if idx == n-1 || s[idx+1] != s[idx] {
				if cnt > renzoku[val] {
					renzoku[val] = cnt
				}
				cnt = 0
			}

			cost[val] += 2
			cost[1-val]++
		}

		cost[0] -= renzoku[0] * 2
		cost[1] -= renzoku[1] * 2
		if cost[0] < cost[1] {
			fmt.Fprintln(out, cost[0])
		} else {
			fmt.Fprintln(out, cost[1])
		}
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
