package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

var sc *bufio.Scanner

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	n, q := readInt(), readInt()
	as := make([]int, q)
	for i := range as {
		as[i] = readInt() - 1
	}
	cnt := 0
	mas := make([]int, n)
	for _, a := range as {
		mas[a] = 1 - mas[a]
		if n == 1 {
			if mas[a] == 1 {
				cnt = 1
			} else {
				cnt = 0
			}
		} else if a == 0 {
			if mas[a] == 1 && mas[a+1] == 0 {
				cnt++
			} else if mas[a] == 0 && mas[a+1] == 0 {
				cnt--
			}
		} else if a == n-1 {
			if mas[a] == 1 && mas[a-1] == 0 {
				cnt++
			} else if mas[a] == 0 && mas[a-1] == 0 {
				cnt--
			}
		} else {
			l := mas[a-1]
			r := mas[a+1]
			if mas[a] == 1 && l == 0 && r == 0 {
				cnt++
			}
			if mas[a] == 0 && l == 0 && r == 0 {
				cnt--
			}
			if mas[a] == 1 && l == 1 && r == 1 {
				cnt--
			}
			if mas[a] == 0 && l == 1 && r == 1 {
				cnt++
			}
		}
		fmt.Fprintln(out, cnt)
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
