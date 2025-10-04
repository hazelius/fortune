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

	x, y := readString(), readString()
	switch x {
	case "Lynx":
		fmt.Fprint(out, "Yes")
		return
	case "Serval":
		if y == "Lynx" {
			fmt.Fprint(out, "No")
			return
		}
		fmt.Fprint(out, "Yes")
	default:
		if y == "Ocelot" {
			fmt.Fprint(out, "Yes")
			return
		}
		fmt.Fprint(out, "No")
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
