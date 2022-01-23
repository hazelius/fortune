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

	s := readString()
	a, b := readInt(), readInt()
	ans := make([]rune, len(s))
	for i, v := range s {
		if i+1 == a {
			ans[b-1] = v
		} else if i+1 == b {
			ans[a-1] = v
		} else {
			ans[i] = v
		}
	}
	return string(ans)
}

func main() {
	fmt.Println(run(os.Stdin))
}
