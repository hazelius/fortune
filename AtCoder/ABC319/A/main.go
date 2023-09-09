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

	rate := map[string]int{
		"tourist":    3858,
		"ksun48":     3679,
		"Benq":       3658,
		"Um_nik":     3648,
		"apiad":      3638,
		"Stonefeang": 3630,
		"ecnerwala":  3613,
		"mnbvmar":    3555,
		"newbiedmy":  3516,
		"semiexp":    3481,
	}

	fmt.Fprint(out, rate[s])
}

func main() {
	run(os.Stdin, os.Stdout)
}
