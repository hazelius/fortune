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
	yaku := make(map[int]bool)
	for i := 1; i < 10; i++ {
		if n%i == 0 {
			yaku[i] = true
		}
	}

	ans := ""
	for i := 0; i <= n; i++ {
		a := 10
		for v := range yaku {
			if i%(n/v) == 0 {
				if a > v {
					a = v
				}
			}
		}
		if a > 9 {
			ans += "-"
		} else {
			ans = fmt.Sprintf("%v%v", ans, a)
		}
	}
	fmt.Fprint(out, ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
