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

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	n := readInt()
	asmap := make(map[int]bool)
	for i := 0; i < n; i++ {
		a := readInt()
		asmap[a] = true
	}
	as := make([]int, len(asmap))
	cnt := 0
	for a := range asmap {
		as[cnt] = a
		cnt++
	}
	sort.Ints(as)
	fmt.Fprintln(out, len(as))
	ansstr := fmt.Sprintf("%v", as)
	fmt.Fprint(out, ansstr[1:len(ansstr)-1])
}

func main() {
	run(os.Stdin, os.Stdout)
}
