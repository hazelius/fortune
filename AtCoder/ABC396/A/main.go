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
	pre := -1
	cnt := 0
	ans := "No"
	for i := 0; i < n; i++ {
		a := readInt()
		if pre == a {
			cnt++
			if cnt >= 2 {
				ans = "Yes"
			}
		} else {
			pre = a
			cnt = 0
		}
	}
	fmt.Fprint(out, ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
