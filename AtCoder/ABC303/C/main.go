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

type position struct {
	x int
	y int
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, m, h, k := readInt(), readInt(), readInt(), readInt()
	s := readString()
	items := make(map[position]bool)
	for i := 0; i < m; i++ {
		p := position{readInt(), readInt()}
		items[p] = true
	}

	now := position{0, 0}
	for i, c := range s {
		if i+h >= n {
			fmt.Fprint(out, "Yes")
			return
		}
		switch c {
		case 'R':
			now.x++
		case 'L':
			now.x--
		case 'U':
			now.y++
		case 'D':
			now.y--
		}
		h--
		if h < 0 {
			break
		}

		if items[now] && h < k {
			h = k
			delete(items, now)
		}

	}
	fmt.Fprint(out, "No")
}

func main() {
	run(os.Stdin, os.Stdout)
}
