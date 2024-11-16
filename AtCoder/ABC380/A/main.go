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

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	n := readString()
	nmap := make(map[rune]int)
	for _, c := range n {
		nmap[c]++
	}
	v, ok := nmap['1']
	if !ok || v != 1 {
		fmt.Fprint(out, "No")
		return
	}
	v, ok = nmap['2']
	if !ok || v != 2 {
		fmt.Fprint(out, "No")
		return
	}
	v, ok = nmap['3']
	if !ok || v != 3 {
		fmt.Fprint(out, "No")
		return
	}
	fmt.Fprint(out, "Yes")
}

func main() {
	run(os.Stdin, os.Stdout)
}
