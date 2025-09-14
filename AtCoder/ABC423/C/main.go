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

	n, r := readInt(), readInt()
	ls := make([]int, n)
	for i := range ls {
		ls[i] = readInt()
	}

	ans := 0
	flg := false
	for i := 0; i < r; i++ {
		if ls[i] == 0 {
			flg = true
		}
		if flg {
			if ls[i] == 0 {
				ans++
			} else {
				ans += 2
			}
		}
	}

	flg = false
	for i := n - 1; i >= r; i-- {
		if ls[i] == 0 {
			flg = true
		}
		if flg {
			if ls[i] == 0 {
				ans++
			} else {
				ans += 2
			}
		}
	}

	fmt.Fprint(out, ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
