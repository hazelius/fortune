package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

var sc *bufio.Scanner
var cls []int

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	n, m := readInt(), readInt()
	uvs := make([][]int, m)
	for i := range uvs {
		uvs[i] = []int{readInt(), readInt()}
	}

	ans := -1
	for i := 0; i < 1<<n; i++ {
		cnt := 0
		for _, uv := range uvs {
			if (i>>uv[0])&1 == (i>>uv[1])&1 {
				cnt++
			}
		}
		if ans < 0 || ans > cnt {
			ans = cnt
		}
	}

	fmt.Fprint(out, ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
