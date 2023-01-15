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

	n := readInt()
	namemap := make(map[string]string)
	for i := 0; i < n; i++ {
		s, t := readString(), readString()
		namemap[s] = t
	}

	for s := range namemap {
		used := make(map[string]bool)
		if !f(s, namemap, used) {
			fmt.Fprint(out, "No")
			return
		}
	}

	fmt.Fprint(out, "Yes")
}

func f(s string, namemap map[string]string, used map[string]bool) bool {
	if used[s] {
		return false
	}
	used[s] = true

	s2, ok := namemap[s]
	if !ok {
		delete(namemap, s)
		return true
	}
	if f(s2, namemap, used) {
		delete(namemap, s)
		return true
	}
	return false
}

func main() {
	run(os.Stdin, os.Stdout)
}
