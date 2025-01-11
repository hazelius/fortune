package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
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

	n := readInt()
	c1, c2 := readString(), readString()
	s := readString()
	for i := 0; i < n; i++ {
		if s[i:i+1] != c1 && s[i:i+1] != c2 {
			s = strings.ReplaceAll(s, s[i:i+1], c2)
		}
	}
	fmt.Fprint(out, s)
}

func main() {
	run(os.Stdin, os.Stdout)
}
