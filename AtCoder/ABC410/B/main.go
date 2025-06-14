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

	n, q := readInt(), readInt()
	xs := make([]int, q)
	for i := range xs {
		xs[i] = readInt()
	}

	anss := make([]int, q)
	boxes := make([]int, n)
	for i, x := range xs {
		if x == 0 {
			mn := q
			idx := 0
			for j, cnt := range boxes {
				if mn > cnt {
					mn = cnt
					idx = j
				}
			}
			anss[i] = idx + 1
			boxes[idx]++
		} else {
			anss[i] = x
			boxes[x-1]++
		}
	}
	str := fmt.Sprintf("%v", anss)
	fmt.Fprint(out, str[1:len(str)-1])
}

func main() {
	run(os.Stdin, os.Stdout)
}
