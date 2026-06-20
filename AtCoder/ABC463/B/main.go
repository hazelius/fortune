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
	buf := make([]byte, 1<<20)
	sc.Buffer(buf, len(buf))
	sc.Split(bufio.ScanWords)

	n, x := readInt(), readString()
	idx := 0
	switch x {
	case "B":
		idx = 1
	case "C":
		idx = 2
	case "D":
		idx = 3
	case "E":
		idx = 4
	}
	ans := "No"
	for i := 0; i < n; i++ {
		s := readString()
		if s[idx] == 'o' {
			ans = "Yes"
		}
	}
	fmt.Fprint(out, ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
