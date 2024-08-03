package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
)

var sc *bufio.Scanner

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

type num struct {
	a   int
	idx int
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	n := readInt()
	as := make([]num, n)
	for i := range as {
		as[i] = num{readInt(), i}
	}

	sort.Slice(as, func(i, j int) bool {
		return as[i].a > as[j].a
	})

	fmt.Fprint(out, as[1].idx+1)
}

func main() {
	run(os.Stdin, os.Stdout)
}
