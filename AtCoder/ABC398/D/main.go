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

type point struct {
	x, y int
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	_, r, c := readInt(), readInt(), readInt()
	s := readString()
	kmap := make(map[point]bool)
	cur := point{r, c}
	head := point{0, 0}
	kmap[head] = true
	for _, ch := range s {
		dir := point{0, 0}
		switch ch {
		case 'N':
			dir = point{1, 0}
		case 'W':
			dir = point{0, 1}
		case 'S':
			dir = point{-1, 0}
		case 'E':
			dir = point{0, -1}
		}

		cur.x += dir.x
		cur.y += dir.y
		head.x += dir.x
		head.y += dir.y
		kmap[head] = true

		ans := "0"
		if kmap[cur] {
			ans = "1"
		}
		fmt.Fprint(out, ans)
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
