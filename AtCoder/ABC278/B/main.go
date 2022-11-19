package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

var sc *bufio.Scanner

func readString() string {
	sc.Scan()
	return sc.Text()
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	h, m := readString(), readString()
	t, _ := time.Parse("15:4", fmt.Sprintf("%v:%v", h, m))

	ha, ma := 0, 0
	for {
		ha = t.Hour()
		ma = t.Minute()
		h2 := ha - ha%10 + ma/10
		m2 := (ha%10)*10 + ma%10
		if h2 < 24 && m2 < 60 {
			break
		}
		t = t.Add(1 * time.Minute)
	}

	fmt.Fprintf(out, "%v %v", ha, ma)
}

func main() {
	run(os.Stdin, os.Stdout)
}
