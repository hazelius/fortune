package main

import (
	"bufio"
	"fmt"
	"io"
	"math/big"
	"os"
	"strconv"
)

var sc *bufio.Scanner

func readInt() int64 {
	sc.Scan()
	i, _ := strconv.ParseInt(sc.Text(), 10, 64)
	return i
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, k := readInt(), readInt()
	k2 := big.NewInt(1)
	for i := int64(0); i < k; i++ {
		k2 = k2.Mul(k2, big.NewInt(int64(10)))
	}

	ans := big.NewInt(1)
	for i := int64(0); i < n; i++ {
		ans = ans.Mul(ans, big.NewInt(readInt()))
		if ans.Cmp(k2) >= 0 {
			ans = big.NewInt(1)
		}
	}
	fmt.Fprint(out, ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
