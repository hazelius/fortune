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

	ab := []int{readInt(), readInt()}
	ansArr := make([]int, 3)
	for _, v := range ab {
		if v%2 > 0 {
			ansArr[0] = 1
		}
		if v == 2 || v == 3 || v >= 6 {
			ansArr[1] = 2
		}
		if v >= 4 {
			ansArr[2] = 4
		}
	}

	ans := 0
	for _, v := range ansArr {
		ans += v
	}

	fmt.Fprint(out, ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
