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

	n := readInt()
	as := make([]int, n)
	for i := range as {
		as[i] = readInt()
	}

	bs := make([]int, n)
	idx := 0
	cnt := 0
	last := -1
	for _, a := range as {
		if a != last {
			last = a
			bs[idx] = a
			idx++
			cnt = 1
			continue
		}
		if cnt == 3 {
			idx -= 3
			last = -1
			cnt = 0
			if idx > 0 {
				last = bs[idx-1]
				cnt++
				for j := idx - 2; j >= 0; j-- {
					if last == bs[j] {
						cnt++
					} else {
						break
					}
				}
			}
			continue
		}
		cnt++
		bs[idx] = a
		idx++
	}

	fmt.Fprint(out, idx)
}

func main() {
	run(os.Stdin, os.Stdout)
}
