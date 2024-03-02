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
	for i := 0; i < n; i++ {
		as := make([]int, 0)
		for j := 0; j < n; j++ {
			if readInt() == 1 {
				as = append(as, j+1)
			}
		}
		if len(as) > 0 {
			ans := fmt.Sprintf("%v", as)
			fmt.Fprintln(out, ans[1:len(ans)-1])
		}
	}

}

func main() {
	run(os.Stdin, os.Stdout)
}
