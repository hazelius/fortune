package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

var sc *bufio.Scanner

type res struct {
	sum   int
	start int
	end   int
}

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	n, k := readInt(), readInt()
	abs := make([][]int, n)
	for i := range abs {
		abs[i] = []int{readInt(), readInt()}
	}

	ans := 0
	sas := make([]int, n)
	for i := range sas {
		ans += abs[i][0]
		sas[i] = abs[i][1] - abs[i][0]
	}

	for i := 0; i < k; i++ {
		rlt := find(sas)
		if rlt.sum <= 0 {
			break
		}
		ans += rlt.sum
		for j := rlt.start; j <= rlt.end; j++ {
			sas[j] = -sas[j]
		}
	}

	fmt.Fprint(out, ans)
}

func find(sas []int) res {
	sum := sas[0]
	start := 0
	end := 1

	tmpMax := sas[0]
	tmpStart := 0

	for i := 1; i < len(sas); i++ {
		if sas[i] > tmpMax+sas[i] {
			tmpMax = sas[i]
			tmpStart = i
		} else {
			tmpMax += sas[i]
		}

		if tmpMax > sum {
			sum = tmpMax
			start = tmpStart
			end = i
		}
	}

	return res{sum: sum, start: start, end: end}
}

func main() {
	run(os.Stdin, os.Stdout)
}
