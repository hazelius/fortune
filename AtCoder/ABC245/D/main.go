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
	as := make([]int, n+1)
	for i := range as {
		as[i] = readInt()
	}
	cs := make([]int, n+m+1)
	for i := range cs {
		cs[i] = readInt()
	}

	bs := make([]int, m+1)
	for i := m; i >= 0; i-- {
		c := cs[i+n]
		for j := 0; j < n; j++ {
			if i+n-j > m {
				continue
			}
			c -= as[j] * bs[i+n-j]
		}
		bs[i] = c / as[n]
	}

	ansstr := fmt.Sprintf("%v", bs)
	return ansstr[1 : len(ansstr)-1]
}

func main() {
	fmt.Println(run(os.Stdin))
}
