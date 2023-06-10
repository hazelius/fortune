package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var sc *bufio.Scanner

func readString() string {
	sc.Scan()
	return sc.Text()
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	p, q := readString(), readString()
	almap := map[string]int{
		"A": 0,
		"B": 1,
		"C": 2,
		"D": 3,
		"E": 4,
		"F": 5,
		"G": 6,
	}

	pi := almap[p]
	qi := almap[q]
	if pi > qi {
		pi, qi = qi, pi
	}

	dis := []int{3, 1, 4, 1, 5, 9, 0}
	ans := 0
	for i := pi; i < qi; i++ {
		ans += dis[i]
	}

	fmt.Fprint(out, ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
