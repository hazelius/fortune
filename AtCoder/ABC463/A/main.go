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

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	buf := make([]byte, 1<<20)
	sc.Buffer(buf, len(buf))
	sc.Split(bufio.ScanWords)

	x, y := readInt(), readInt()
	if x*9 == y*16 {
		fmt.Fprint(out, "Yes")
		return
	}
	fmt.Fprint(out, "No")
}

func main() {
	run(os.Stdin, os.Stdout)
}
