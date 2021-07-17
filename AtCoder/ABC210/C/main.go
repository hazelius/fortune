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
	sc.Split(bufio.ScanWords)

	n, k := readInt(), readInt()
	cs := make([]string, n)
	cmap := make(map[string]int, k)
	ans := 0
	for i := range cs {
		cs[i] = readString()
		if i >= k {
			delKey := cs[i-k]
			cnt := cmap[delKey] - 1
			if cnt == 0 {
				delete(cmap, delKey)
			} else {
				cmap[delKey] = cnt
			}
		}
		cnt := cmap[cs[i]]
		cmap[cs[i]] = cnt + 1
		if len(cmap) > ans {
			ans = len(cmap)
		}
	}

	return ans
}

func main() {
	fmt.Println(run(os.Stdin))
}
