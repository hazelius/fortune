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

type num struct {
	pre  int
	next int
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	n := readInt()
	numMap := make(map[int]num)
	pre := -1
	for i := 0; i < n; i++ {
		a := readInt()
		numMap[a] = num{pre, -1}
		if i > 0 {
			preNum := numMap[pre]
			preNum.next = a
			numMap[pre] = preNum
		}
		pre = a
	}

	q := readInt()
	for i := 0; i < q; i++ {
		switch readInt() {
		case 1:
			x, y := readInt(), readInt()
			xnum := numMap[x]
			numMap[y] = num{x, xnum.next}
			if xnum.next >= 0 {
				nextNum := numMap[xnum.next]
				nextNum.pre = y
				numMap[xnum.next] = nextNum
			}
			xnum.next = y
			numMap[x] = xnum
		case 2:
			x := readInt()
			xnum := numMap[x]
			if xnum.pre >= 0 {
				preNum := numMap[xnum.pre]
				preNum.next = xnum.next
				numMap[xnum.pre] = preNum
			}
			if xnum.next >= 0 {
				nextNum := numMap[xnum.next]
				nextNum.pre = xnum.pre
				numMap[xnum.next] = nextNum
			}
			delete(numMap, x)
		}
	}

	var tnum num
	for k, val := range numMap {
		if val.pre == -1 {
			fmt.Fprint(out, k, " ")
			tnum = val
			break
		}
	}

	for tnum.next != -1 {
		fmt.Fprint(out, tnum.next, " ")
		tnum = numMap[tnum.next]
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
