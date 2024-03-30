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
	ans := make([]int, n)

	cur := 0
	xmap := make(map[int]bool)
	for i := 0; i < q; i++ {
		x := readInt() - 1
		if xmap[x] {
			ans[x] += cur
			delete(xmap, x)
		} else {
			ans[x] -= cur
			xmap[x] = true
		}
		cur += len(xmap)
	}

	for k := range xmap {
		ans[k] += cur
	}

	ansstr := fmt.Sprintf("%v", ans)
	fmt.Fprint(out, ansstr[1:len(ansstr)-1])
}

func main() {
	run(os.Stdin, os.Stdout)
}
