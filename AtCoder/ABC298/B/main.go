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

	n := readInt()
	as := make([][]int, n)
	for i := range as {
		as[i] = make([]int, n)
		for j := range as[i] {
			as[i][j] = readInt()
		}
	}
	bs := make([][]int, n)
	for i := range bs {
		bs[i] = make([]int, n)
		for j := range bs[i] {
			bs[i][j] = readInt()
		}
	}

	for cnt := 0; cnt < 4; cnt++ {
		flg := true
		for i := range as {
			for j := range as[i] {
				x, y := i, j
				for k := 0; k < cnt; k++ {
					x, y = next(x, y, n)
				}
				if as[i][j] == 1 && bs[x][y] == 0 {
					flg = false
					break
				}
			}
			if !flg {
				break
			}
		}
		if flg {
			fmt.Fprint(out, "Yes")
			return
		}
	}

	fmt.Fprint(out, "No")
}

func next(x, y, n int) (int, int) {
	return n - y - 1, x
}

func main() {
	run(os.Stdin, os.Stdout)
}
