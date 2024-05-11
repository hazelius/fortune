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

	sum := 0
	ketas := make([]int, 11)
	for _, a := range as {
		sum = (sum + a) % mod
		keta := len(strconv.Itoa(a))
		ketas[keta]++
	}

	p10 := make([]int, 11)
	for i := range p10 {
		if i == 0 {
			p10[i] = 1
		} else {
			p10[i] = p10[i-1] * 10
		}
	}

	ans := 0
	for _, a := range as {
		sum -= a
		keta := len(strconv.Itoa(a))
		ketas[keta]--
		for k, v := range ketas {
			v = (p10[k] * v) % mod
			v = (a * v) % mod
			ans = (ans + v) % mod
		}
		ans = (ans + mod + sum) % mod
	}

	fmt.Fprint(out, (ans+mod)%mod)
}

func main() {
	run(os.Stdin, os.Stdout)
}
