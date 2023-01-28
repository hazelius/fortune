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

	n, m := readInt(), readInt()
	as := make([]string, n)
	for i := range as {
		as[i] = readString()
	}
	tmap := make(map[string]bool)
	for i := 0; i < m; i++ {
		t := readString()
		tmap[t] = true
	}
	ans := 0
	for _, a := range as {
		if tmap[a[3:]] {
			ans++
		}
	}

	fmt.Fprint(out, ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
