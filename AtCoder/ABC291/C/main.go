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
	x int
	y int
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := readInt()
	s := readString()

	pmap := make(map[point]bool)
	cur := point{0, 0}
	pmap[cur] = true

	for i := 0; i < n; i++ {
		c := s[i]
		switch c {
		case 'R':
			cur.x += 1
		case 'L':
			cur.x -= 1
		case 'U':
			cur.y += 1
		case 'D':
			cur.y -= 1
		}
		if pmap[cur] {
			fmt.Fprint(out, "Yes")
			return
		}
		pmap[cur] = true
	}
	fmt.Fprint(out, "No")
}

func main() {
	run(os.Stdin, os.Stdout)
}
