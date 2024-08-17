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
	rs := make([]int, n)
	ls := make([]int, n)
	sum := 0
	for i := range rs {
		l, r := readInt(), readInt()
		sum += l
		ls[i] = l
		rs[i] = r
	}

	if sum > 0 {
		fmt.Fprint(out, "No")
		return
	}

	for i, r := range rs {
		ss := r - ls[i]
		if sum+ss > 0 {
			ss = -sum
		}
		ls[i] += ss
		sum += ss
		if sum == 0 {
			break
		}
	}

	if sum < 0 {
		fmt.Fprint(out, "No")
		return
	}

	fmt.Fprintln(out, "Yes")
	ansstr := fmt.Sprintf("%v", ls)
	fmt.Fprint(out, ansstr[1:len(ansstr)-1])
}

func main() {
	run(os.Stdin, os.Stdout)
}
