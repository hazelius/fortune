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

	sx, sy := readInt(), readInt()
	tx, ty := readInt(), readInt()
	wi := abs(tx - sx)
	hi := abs(ty - sy)
	if (sx+sy)%2 > 0 {
		if tx-sx < 0 {
			wi--
		}
	} else if tx-sx > 0 {
		wi--
	}

	if wi > hi {
		wi -= hi
		wi = (wi + 1) / 2
	} else {
		wi = 0
	}

	fmt.Fprint(out, hi+wi)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	run(os.Stdin, os.Stdout)
}
