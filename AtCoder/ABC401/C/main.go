package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

var sc *bufio.Scanner

var mod = 1000000000

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	n, k := readInt(), readInt()
	as := make([]int, n+2)
	for i := range as {
		if i <= k {
			as[i] = i
		} else {
			as[i] = (as[i-1]*2 - as[i-k-1]) % mod
		}
	}
	fmt.Fprint(out, (as[n+1]+mod-as[n])%mod)
}

func main() {
	run(os.Stdin, os.Stdout)
}
