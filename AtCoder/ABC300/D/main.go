// https://atcoder.jp/contests/abc300/tasks/abc300_d
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

	sosu := make([]int, 0)
	mx := 5
	for mx*mx*12 <= n {
		mx++
	}
	mx--
	s := make([]bool, mx+1)
	for i := 2; i <= mx; i++ {
		if !s[i] {
			sosu = append(sosu, i)
			for j := 2 * i; j <= mx; j += i {
				s[j] = true
			}
		}
	}

	sosucnt := len(sosu)
	ans := 0
	for i := 0; i < sosucnt; i++ {
		k := sosucnt - 1
		for j := i + 1; j < k && j < sosucnt; j++ {
			for j < k {
				v := sosu[i] * sosu[i] * sosu[j]
				if v > n {
					k--
					continue
				}
				v *= sosu[k]
				if v > n {
					k--
					continue
				}
				v *= sosu[k]
				if v > n {
					k--
					continue
				}
				break
			}
			ans += k - j
		}
	}

	fmt.Fprint(out, ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
