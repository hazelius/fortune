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
	if s[0] == '1' {
		fmt.Fprint(out, "No")
		return
	}

	r := make([]int, 7)
	for i, c := range s {
		if c == '0' {
			continue
		}
		switch i + 1 {
		case 7:
			r[0] = 1
		case 4:
			r[1] = 1
		case 8, 2:
			r[2] = 1
		case 5, 1:
			r[3] = 1
		case 9, 3:
			r[4] = 1
		case 6:
			r[5] = 1
		case 10:
			r[6] = 1
		}
	}

	flg1, flg2 := false, false
	for _, v := range r {
		if v == 1 {
			if flg1 && flg2 {
				fmt.Fprint(out, "Yes")
				return
			}
			flg1 = true
		} else if flg1 {
			flg2 = true
		}
	}

	fmt.Fprint(out, "No")
}

func main() {
	run(os.Stdin, os.Stdout)
}
