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

	ss := make([]int, n)
	for num := range ss {
		if num == 0 {
			continue
		}
		if num == 1 {
			ss[num] = 1
			continue
		}
		if num == 4 {
			ss[num] = 3
			continue
		}
		s := 2
		for j := 2; j*j <= num; j++ {
			if num%j == 0 {
				h := num / j
				if j > h {
					break
				} else if j == h {
					s++
					break
				}
				s += 2
			}
		}
		ss[num] = s
	}

	ans := 0
	for i := 1; i <= n/2; i++ {
		a := ss[i] * ss[n-i]
		if i != n-i {
			a *= 2
		}
		ans += a
	}

	fmt.Fprint(out, ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
