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

func run(r io.Reader) string {
	sc = bufio.NewScanner(r)
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, q := readInt(), readInt()
	s := readString()
	ans := make([]string, q)
	qidx := 0
	pos := 0
	for i := 0; i < q; i++ {
		t, x := readInt(), readInt()
		if t == 1 {
			pos = pos - x
			if pos < 0 {
				pos = n + pos
			}
		} else {
			idx := (pos + x - 1) % n
			ans[qidx] = s[idx : idx+1]
			qidx++
		}
	}
	ansstr := fmt.Sprintf("%v", ans[:qidx])
	return ansstr[1 : len(ansstr)-1]
}

func main() {
	fmt.Println(run(os.Stdin))
}
