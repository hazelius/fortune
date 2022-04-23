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

func run(r io.Reader) int {
	sc = bufio.NewScanner(r)
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, k := readInt(), readInt()
	ss := make([]string, n)
	for i := range ss {
		ss[i] = readString()
	}

	ans := 0
	for i := 0; i < 1<<15; i++ {
		maps := make(map[rune]int)
		for j, s := range ss {
			if 1<<j&i == 0 {
				continue
			}
			for _, c := range s {
				maps[c]++
			}
		}
		kcnt := 0
		for _, cnt := range maps {
			if cnt == k {
				kcnt++
			}
		}
		if ans < kcnt {
			ans = kcnt
		}
	}

	return ans
}

func main() {
	fmt.Println(run(os.Stdin))
}
