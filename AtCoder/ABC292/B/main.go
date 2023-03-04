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
	sc.Split(bufio.ScanWords)

	_, q := readInt(), readInt()

	mem := make(map[int]int)
	for i := 0; i < q; i++ {
		c, x := readInt(), readInt()
		switch c {
		case 1:
			mem[x] += 1
		case 2:
			mem[x] += 2
		case 3:
			cnt, ok := mem[x]
			if ok && cnt > 1 {
				fmt.Fprintln(out, "Yes")
			} else {
				fmt.Fprintln(out, "No")
			}
		}
	}

}

func main() {
	run(os.Stdin, os.Stdout)
}
