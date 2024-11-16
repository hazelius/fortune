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

	_, k := readInt(), readInt()
	s := readString()
	sbyte := []byte(s)

	flg := false
	cnt := 0
	start, end := 0, 0
	e2 := 0
	for i, v := range sbyte {
		if !flg && v == '1' {
			cnt++
			flg = true
			start = i
		} else if flg && v == '0' {
			flg = false
			end = i
			if cnt == k-1 {
				e2 = end
			}
			if cnt == k {
				break
			}
		}
	}

	if start < end {
		fmt.Fprint(out, s[:e2]+s[start:end]+s[e2:start]+s[end:])
	} else {
		fmt.Fprint(out, s[:e2]+s[start:]+s[e2:start])
	}

}

func main() {
	run(os.Stdin, os.Stdout)
}
