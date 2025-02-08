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

	n, m := readInt(), readInt()
	amap := make(map[int]bool)
	for i := 0; i < m; i++ {
		a := readInt()
		amap[a] = true
	}
	ans := make([]int, n)
	idx := 0
	for i := 1; i <= n; i++ {
		if amap[i] {
			continue
		}
		ans[idx] = i
		idx++
	}
	ans = ans[:idx]
	str := fmt.Sprintf("%v", ans)
	fmt.Fprint(out, len(ans))
	if len(ans) > 0 {
		fmt.Fprintf(out, "\n%v", str[1:len(str)-1])
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
