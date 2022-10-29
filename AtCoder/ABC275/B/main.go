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

	mod := 998244353
	a, b, c, d, e, f := readInt()%mod, readInt()%mod, readInt()%mod, readInt()%mod, readInt()%mod, readInt()%mod
	aa := a * b % mod
	aa = aa * c % mod
	ad := d * e % mod
	ad = ad * f % mod
	ans := (aa + mod) - ad
	fmt.Fprint(out, ans%mod)
}

func main() {
	run(os.Stdin, os.Stdout)
}
