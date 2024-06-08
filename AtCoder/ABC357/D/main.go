package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

var sc *bufio.Scanner
var mod = 998244353

func readString() string {
	sc.Scan()
	return sc.Text()
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	s := readString()
	keta := len(s)
	n, _ := strconv.Atoi(s)

	ans := n
	ans = ans * (modPow(modPow(10, keta), n) - 1) % mod
	ans = ans * modInv(modPow(10, keta)-1) % mod
	fmt.Fprint(out, ans)
}

func modInv(x int) int {
	return modPow(x, mod-2)
}

func modPow(x, n int) int {
	res := 1
	for n > 0 {
		if n&1 == 1 {
			res = res * x % mod
		}
		x = x * x % mod
		n >>= 1
	}
	return res % mod
}

func main() {
	run(os.Stdin, os.Stdout)
}
