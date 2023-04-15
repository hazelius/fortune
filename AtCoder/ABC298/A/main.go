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
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	readInt()
	s := readString()
	if !strings.Contains(s, "o") {
		fmt.Fprint(out, "No")
		return
	}
	if strings.Contains(s, "x") {
		fmt.Fprint(out, "No")
		return
	}
	fmt.Fprint(out, "Yes")
}

func main() {
	run(os.Stdin, os.Stdout)
}
