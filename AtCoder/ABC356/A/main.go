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

	n, l, r := readInt(), readInt(), readInt()
	ans := make([]int, n)
	for i := range ans {
		if i+1 < l {
			ans[i] = i + 1
		} else if i < r {
			ans[i] = r - (i + 1 - l)
		} else {
			ans[i] = i + 1
		}
	}
	anstr := fmt.Sprintf("%d", ans)
	fmt.Fprint(out, anstr[1:len(anstr)-1])
}

func main() {
	run(os.Stdin, os.Stdout)
}
