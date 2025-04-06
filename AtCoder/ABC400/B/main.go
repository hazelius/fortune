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
	ans := 0
	n2 := 1
	for i := 0; i <= m; i++ {
		if n2 > 1e9 {
			fmt.Fprint(out, "inf")
			return
		}
		ans += n2
		if ans > 1e9 {
			fmt.Fprint(out, "inf")
			return
		}
		n2 *= n
	}
	fmt.Fprint(out, ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
