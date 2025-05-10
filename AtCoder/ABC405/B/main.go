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

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	n, m := readInt(), readInt()
	as := make([]int, n)
	amap := make(map[int]int)
	for i := range as {
		a := readInt()
		as[i] = a
		amap[a]++
	}
	if len(amap) < m {
		fmt.Fprint(out, 0)
		return
	}

	cnt := 0
	for i := n - 1; i >= 0; i-- {
		cnt++
		if amap[as[i]] == 1 {
			delete(amap, as[i])
			if len(amap) < m {
				break
			}
		} else {
			amap[as[i]]--
		}
	}
	fmt.Fprint(out, cnt)
}

func main() {
	run(os.Stdin, os.Stdout)
}
