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

	abc := []int{readInt(), readInt(), readInt()}
	sort.Ints(abc)
	fmt.Fprintf(out, "%v%v%v", abc[2], abc[1], abc[0])
}

func main() {
	run(os.Stdin, os.Stdout)
}
