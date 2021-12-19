package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var sc *bufio.Scanner

func readString() string {
	sc.Scan()
	return sc.Text()
}

func run(r io.Reader) string {
	sc = bufio.NewScanner(r)
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	s, t := readString(), readString()

	abc := "abcdefghijklmnopqrstuvwxyz"
	abcm := make(map[byte]int)
	for i := 0; i < len(abc); i++ {
		abcm[abc[i]] = i
	}

	s1 := abcm[s[0]]
	t1 := abcm[t[0]]
	if s1 > t1 {
		t1 += len(abc)
	}
	d := t1 - s1
	for i := 0; i < len(s); i++ {
		ai := abcm[s[i]]
		bi := abcm[t[i]]
		if bi != (ai+d)%len(abc) {
			return "No"
		}
	}

	return "Yes"
}

func main() {
	fmt.Println(run(os.Stdin))
}
