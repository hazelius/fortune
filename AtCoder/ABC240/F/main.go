// F - Sum Sum Max
// https://atcoder.jp/contests/abc240/tasks/abc240_f
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

	t := readInt()
	ans := make([]int, t)
	for i := 0; i < t; i++ {
		n, _ := readInt(), readInt()

		maxc := 0
		b, c := 0, 0
		for j := 0; j < n; j++ {
			x, y := readInt(), readInt()
			if j == 0 {
				maxc = x
			}

			if b > 0 && x < 0 {
				cnt := b / (x * -1)
				if cnt < y {
					tc := c + b*cnt + x*cnt*(cnt+1)/2
					if maxc < tc {
						maxc = tc
					}
				}
			}

			c += b*y + x*y*(y+1)/2
			b += x * y
			if maxc < c {
				maxc = c
			}
		}

		ans[i] = maxc
	}

	ansstr := fmt.Sprintf("%v", ans)
	return ansstr[1 : len(ansstr)-1]
}

func main() {
	fmt.Println(run(os.Stdin))
}
