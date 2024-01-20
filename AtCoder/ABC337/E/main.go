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

func readString() string {
	sc.Scan()
	return sc.Text()
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	n := readInt()

	b := 0
	for ; 1<<b < n; b++ {
	}
	fmt.Fprintln(out, b)

	for i := 0; i < b; i++ {
		bs := make([]int, 0)
		for j := 0; j < n; j++ {
			if j&(1<<i) > 0 {
				bs = append(bs, j+1)
			}
		}
		bsstr := fmt.Sprintf("%v", bs)
		fmt.Fprintln(out, len(bs), bsstr[1:len(bsstr)-1])
	}

	s := readString()
	ans := 0
	for i := 0; i < b; i++ {
		if s[i] == '1' {
			ans |= 1 << i
		}
	}
	fmt.Fprint(out, ans+1)
}

func main() {
	run(os.Stdin, os.Stdout)
}
