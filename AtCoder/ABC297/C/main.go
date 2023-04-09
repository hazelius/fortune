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

	h, _ := readInt(), readInt()
	hws := make([]string, h)
	for i := range hws {
		hws[i] = readString()
	}
	for _, s := range hws {
		ans := strings.ReplaceAll(s, "TT", "PC")
		fmt.Fprintln(out, ans)
	}

}

func main() {
	run(os.Stdin, os.Stdout)
}
