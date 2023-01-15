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

	s := readString()

	abc := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	ans := 0
	kisu := 1
	for i := len(s) - 1; i >= 0; i-- {
		for j := 0; j < len(abc); j++ {
			if s[i] == abc[j] {
				ans += (j + 1) * kisu
			}
		}
		kisu *= len(abc)
	}

	fmt.Fprint(out, ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
