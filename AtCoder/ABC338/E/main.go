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

	n := readInt()
	n2 := n * 2
	abs := make([]int, n2)
	for i := 0; i < n; i++ {
		a, b := readInt()-1, readInt()-1
		abs[a] = b
		abs[b] = a
	}

	que := make([]int, 0)
	for i := 0; i < n2; i++ {
		a := abs[i]
		if i < a {
			que = append(que, i)
		} else {
			b := que[len(que)-1]
			if b != a {
				fmt.Fprint(out, "Yes")
				return
			}
			que = que[:len(que)-1]
		}
	}
	fmt.Fprint(out, "No")
}

func main() {
	run(os.Stdin, os.Stdout)
}
