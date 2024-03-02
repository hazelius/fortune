package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
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

	ms := sort.Search(1000000, func(i int) bool {
		return i*i*i > n
	})

	for i := ms - 1; i > 0; i-- {
		t := fmt.Sprintf("%v", i*i*i)
		flg := true
		for idx := 0; idx < len(t)/2; idx++ {
			if t[idx] != t[len(t)-1-idx] {
				flg = false
				break
			}
		}
		if flg {
			fmt.Fprint(out, t)
			return
		}
	}

}

func main() {
	run(os.Stdin, os.Stdout)
}
