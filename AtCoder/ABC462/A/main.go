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
	buf := make([]byte, 1<<20)
	sc.Buffer(buf, len(buf))
	sc.Split(bufio.ScanWords)

	s := readString()
	rex := regexp.MustCompile("[0-9]+")
	ans := rex.FindAllString(s, -1)
	fmt.Fprint(out, strings.Join(ans, ""))
}

func main() {
	run(os.Stdin, os.Stdout)
}
