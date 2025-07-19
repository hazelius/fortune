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

	t := readInt()
	for i := 0; i < t; i++ {
		n := readInt()
		s := "0" + readString()
		ok := make([]bool, 1<<n)
		ok[0] = true
		for j := 0; j < (1 << n); j++ {
			if !ok[j] {
				continue
			}
			for k := 0; k < n; k++ {
				if j&(1<<k) > 0 {
					continue
				}
				next := j | (1 << k)
				if s[next] == '1' {
					continue
				}
				ok[next] = true
			}
		}
		if ok[(1<<n)-1] {
			fmt.Fprintln(out, "Yes")
		} else {
			fmt.Fprintln(out, "No")
		}
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
