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

	n := readInt()
	s, t := readString(), readString()

	ans := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			flg := false
			for k := 0; k < n; k++ {
				ic := (i + k) % n
				jc := (j + k) % n
				if s[ic] < t[jc] {
					ans++
					flg = true
					break
				}
				if s[ic] > t[jc] {
					flg = true
					break
				}
			}
			if !flg {
				ans++
			}
		}
	}

	fmt.Fprint(out, ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
