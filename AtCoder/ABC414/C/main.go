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

	a, n := readInt(), readInt()
	ans := 0
	for i := 1; i <= n; i++ {
		num := i
		tmp := i
		for tmp/10 > 0 {
			tmp /= 10
			num = num*10 + (tmp % 10)
		}
		if num > n {
			break
		}
		if kaibun(num, a) {
			ans += num
		}
	}

	for i := 1; i <= n; i++ {
		num := i
		tmp := i
		for {
			num = num*10 + (tmp % 10)
			if tmp < 10 {
				break
			}
			tmp /= 10
		}
		if num > n {
			break
		}
		if kaibun(num, a) {
			ans += num
		}
	}

	fmt.Fprint(out, ans)
}

func kaibun(num, a int) bool {
	reg := make([]int, 0)
	for {
		reg = append(reg, num%a)
		if num < a {
			break
		}
		num /= a
	}

	for i := 0; i*2 < len(reg); i++ {
		if reg[i] != reg[(len(reg)-i-1)] {
			return false
		}
	}
	return true
}

func main() {
	run(os.Stdin, os.Stdout)
}
