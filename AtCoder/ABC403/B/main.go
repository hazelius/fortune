package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

var sc *bufio.Scanner

func readString() string {
	sc.Scan()
	return sc.Text()
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	t, u := readString(), readString()
	t = strings.ReplaceAll(t, "?", ".")
	for i := 0; i <= len(t)-len(u); i++ {
		if regexp.MustCompile(t[i : i+len(u)]).MatchString(u) {
			fmt.Fprint(out, "Yes")
			return
		}
	}
	fmt.Fprint(out, "No")
}

func main() {
	run(os.Stdin, os.Stdout)
}
