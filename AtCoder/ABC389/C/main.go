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

	q := readInt()
	ans := make([]int, 0, q)
	long, gap := 0, 0
	for i := 0; i < q; i++ {
		switch readInt() {
		case 1:
			long += readInt()
			ans = append(ans, long)
		case 2:
			gap += ans[0] - gap
			ans = ans[1:]
		case 3:
			k := readInt() - 1
			if k == 0 {
				fmt.Fprintln(out, 0)
			} else {
				fmt.Fprintln(out, ans[k-1]-gap)
			}
		}
	}

}

func main() {
	run(os.Stdin, os.Stdout)
}
