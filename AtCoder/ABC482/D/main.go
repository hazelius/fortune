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
	as := make([]int, n)
	for i := range as {
		as[i] = readInt()
	}

	for i := 0; i < n; i++ {
		if i == 0 || i == n-1 {
			fmt.Fprintln(out, (as[n-1]-as[0])/2)
			continue
		}
		ans1 := (as[n-1] - as[i]) / 2
		ans2 := (as[i+1] - as[0]) / 2
		if ans1 < ans2 {
			ans1 = ans2
		}

		ans3 := (as[i] - as[0]) / 2
		ans4 := (as[n-1] - as[i-1]) / 2
		if ans3 < ans4 {
			ans3 = ans4
		}

		if ans1 > ans3 {
			ans1 = ans3
		}
		fmt.Fprintln(out, ans1)
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
