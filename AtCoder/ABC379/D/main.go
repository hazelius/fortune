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

	q := readInt()
	ts := make([]int, q)
	vmap := make(map[int]int)
	idx := 0

	for i := 0; i < q; i++ {
		switch readInt() {
		case 1:
			vmap[idx]++
		case 2:
			t := readInt()
			idx++
			ts[idx] = ts[idx-1] + t
		case 3:
			h := readInt()
			ans := 0
			for no, cnt := range vmap {
				if ts[idx]-ts[no] >= h {
					ans += cnt
					delete(vmap, no)
				}
			}
			fmt.Fprintln(out, ans)
		}
	}

}

func main() {
	run(os.Stdin, os.Stdout)
}
