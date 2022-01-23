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

func run(r io.Reader) string {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	n, _ := readInt(), readInt()
	ss := make([]string, n)
	for i := range ss {
		ss[i] = readString()
	}

	ans := make([]string, n)
	t := readString()

	for i, s := range ss {
		if s == t {
			ans[i] = "Yes"
			t = readString()
		} else {
			ans[i] = "No"
		}
	}

	ansstr := fmt.Sprintf("%v", ans)
	return ansstr[1 : len(ansstr)-1]
}

func main() {
	fmt.Println(run(os.Stdin))
}
