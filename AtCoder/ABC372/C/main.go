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
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, q := readInt(), readInt()
	s := readString()
	ans := 0
	i := 0
	for {
		next := strings.Index(s[i:], "ABC")
		if next < 0 {
			break
		}
		i = i + next + 1
		ans++
	}

	bytes := []byte(s)
	for i := 0; i < q; i++ {
		x, c := readInt()-1, readString()
		if bytes[x] == c[0] {
			fmt.Fprintln(out, ans)
			continue
		}
		for idx := x - 2; idx <= x; idx++ {
			if idx < 0 {
				continue
			}
			if idx+2 >= n {
				break
			}
			if isABC(bytes, idx) {
				ans--
				break
			}
		}
		bytes[x] = c[0]
		for idx := x - 2; idx <= x; idx++ {
			if idx < 0 {
				continue
			}
			if idx+2 >= n {
				break
			}
			if isABC(bytes, idx) {
				ans++
				break
			}
		}

		fmt.Fprintln(out, ans)
	}

}

func isABC(bytes []byte, idx int) bool {
	return bytes[idx] == 'A' && bytes[idx+1] == 'B' && bytes[idx+2] == 'C'
}

func main() {
	run(os.Stdin, os.Stdout)
}
