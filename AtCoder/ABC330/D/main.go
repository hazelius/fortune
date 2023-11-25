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
	sc.Split(bufio.ScanWords)

	n := readInt()

	xymap := make(map[int][]int)
	yxmap := make(map[int][]int)
	for i := 0; i < n; i++ {
		s := readString()
		for j, c := range s {
			if c == 'o' {
				v, ok := xymap[i]
				if !ok {
					v = make([]int, 0)
				}
				v = append(v, j)
				xymap[i] = v

				y, ok := yxmap[j]
				if !ok {
					y = make([]int, 0)
				}
				y = append(y, i)
				yxmap[j] = y
			}
		}
	}

	ans := 0
	for _, ys := range xymap {
		for _, y := range ys {
			xs := yxmap[y]
			ans += (len(xs) - 1) * (len(ys) - 1)
		}
	}

	fmt.Fprint(out, ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
