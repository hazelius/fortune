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
	amap := make(map[int]bool)

	for i := 1; i <= n; i++ {
		x := readInt()
		if !amap[i] {
			amap[x] = true
		}
	}

	ans := make([]int, 0)
	for i := 1; i <= n; i++ {
		if !amap[i] {
			ans = append(ans, i)
		}
	}

	ansstr := fmt.Sprintf("%v", ans)
	fmt.Fprintln(out, len(ans))
	fmt.Fprint(out, ansstr[1:len(ansstr)-1])
}

func main() {
	run(os.Stdin, os.Stdout)
}
