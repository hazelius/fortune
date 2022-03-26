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
	for i := range bs {
		c := cs[i]
		for j := n; j > 0; j-- {
			if j > i {
				continue
			}
			c -= as[j] * bs[i-j]
		}
		if as[0] == 0 {
			bs[i] = c
		} else {
			bs[i] = c / as[0]
		}
	}

	ansstr := fmt.Sprintf("%v", bs)
	return ansstr[1 : len(ansstr)-1]
}

func main() {
	fmt.Println(run(os.Stdin))
}
