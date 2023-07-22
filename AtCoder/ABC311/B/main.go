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
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, d := readInt(), readInt()
	as := make([]string, n)
	for i := range as {
		as[i] = readString()
	}
	os := make([]bool, d)
	for i := range os {
		os[i] = true
	}

	for i := range os {
		for _, a := range as {
			if a[i] == 'x' {
				os[i] = false
				break
			}
		}
	}

	ans := 0
	tmp := 0
	for _, v := range os {
		if v {
			tmp++
			if ans < tmp {
				ans = tmp
			}
		} else {
			tmp = 0
		}
	}

	fmt.Fprint(out, ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
