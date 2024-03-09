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

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	t := readString()
	n := readInt()

	dpMap := make(map[int]int)
	dpMap[0] = 0
	for i := 0; i < n; i++ {
		a := readInt()
		nextMap := make(map[int]int)
		for j := 0; j < a; j++ {
			s := readString()
			for idx, cost := range dpMap {
				last := idx + len(s)
				if last > len(t) {
					continue
				}
				if s == t[idx:last] {
					old, ok := nextMap[last]
					if !ok || old > cost+1 {
						nextMap[last] = cost + 1
					}
				}
			}
		}
		for k, next := range nextMap {
			old, ok := dpMap[k]
			if !ok || old > next {
				dpMap[k] = next
			}
		}
	}

	ans, ok := dpMap[len(t)]
	if !ok {
		fmt.Fprint(out, -1)
	} else {
		fmt.Fprint(out, ans)
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
