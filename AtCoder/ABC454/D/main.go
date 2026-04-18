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
	buf := make([]byte, 10*1024*1024)
	sc.Buffer(buf, len(buf))
	sc.Split(bufio.ScanWords)

	t := readInt()
	// 各テストケースを処理
	for i := 0; i < t; i++ {
		a, b := readString(), readString()
		a = rep([]byte(a))
		b = rep([]byte(b))
		if a == b {
			fmt.Fprintln(out, "Yes")
		} else {
			fmt.Fprintln(out, "No")
		}
	}
}

func rep(s []byte) string {
	// スタックを用いて (xx) を見つけ次第 xx に置換していく
	ret := make([]byte, 0, len(s))
	for _, b := range s {
		ret = append(ret, b)
		n := len(ret)
		// 末尾が (xx) になったら即座に xx に置き換える
		if n >= 4 && ret[n-1] == ')' && ret[n-2] == 'x' && ret[n-3] == 'x' && ret[n-4] == '(' {
			ret = ret[:n-4]
			ret = append(ret, 'x', 'x')
		}
	}
	return string(ret)
}

func main() {
	run(os.Stdin, os.Stdout)
}
