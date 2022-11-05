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
	ans := 0
	mind, mint, other := 0, 0, 0
	for i := 0; i < n; i++ {
		a := readInt()
		d, t, o := fcount(a)
		if i == 0 {
			mind = d
			mint = t
			other = o
			continue
		}

		if d > mind {
			ans += d - mind
		} else if d < mind {
			ans += (mind - d) * i
			mind = d
		}

		if t > mint {
			ans += t - mint
		} else if t < mint {
			ans += (mint - t) * i
			mint = t
		}

		if o != other {
			fmt.Fprint(out, -1)
			return
		}
	}

	fmt.Fprint(out, ans)
}

func fcount(a int) (int, int, int) {
	d, t := 0, 0

	for {
		if a%2 > 0 {
			break
		}
		a /= 2
		d++
	}
	for {
		if a%3 > 0 {
			break
		}
		a /= 3
		t++
	}

	return d, t, a
}

func main() {
	run(os.Stdin, os.Stdout)
}
