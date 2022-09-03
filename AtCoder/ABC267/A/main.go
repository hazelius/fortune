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
	switch s {
	case "Monday":
		fmt.Fprint(out, 5)
	case "Tuesday":
		fmt.Fprint(out, 4)
	case "Wednesday":
		fmt.Fprint(out, 3)
	case "Thursday":
		fmt.Fprint(out, 2)
	case "Friday":
		fmt.Fprint(out, 1)
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
