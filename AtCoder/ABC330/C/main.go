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

	d := readInt()

	ans := d
	y := 2000000
	for i := 0; i <= 2000000; i++ {
		x2 := i * i
		for ; y > 0 && x2+y*y > d; y-- {

		}
		ans = min(ans, abs(x2+(y*y)-d))
		ans = min(ans, abs(x2+((y+1)*(y+1))-d))

		if d-x2 < 0 {
			break
		}
	}

	fmt.Fprint(out, ans)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	run(os.Stdin, os.Stdout)
}
