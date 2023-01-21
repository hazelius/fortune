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

	n, x := readInt(), readInt()
	as := make([][]int, n)
	for i := range as {
		as[i] = []int{readInt(), readInt()}
	}

	pmap := make(map[int]bool)
	pmap[0] = true

	for _, ab := range as {
		a, b := ab[0], ab[1]
		nextMap := make(map[int]bool)
		for p := range pmap {
			if p > x {
				continue
			}
			for i := 0; i <= b; i++ {
				prise := p + a*i
				if prise == x {
					fmt.Fprint(out, "Yes")
					return
				}
				nextMap[prise] = true
			}
		}
		pmap = nextMap
	}

	if pmap[x] {
		fmt.Fprint(out, "Yes")
	} else {
		fmt.Fprint(out, "No")
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
