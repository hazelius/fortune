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
	as := make([]int, n)
	for i := range as {
		as[i] = readInt()
	}

	used := make(map[int]int)
	pfa := make(map[int]int)
	zero := 0
	for _, a := range as {
		if a == 0 {
			zero++
			continue
		}

		v, ok := used[a]
		if ok {
			pfa[v]++
			continue
		}

		p := PrimeFactorization(a)
		p2 := 1
		for _, v := range p {
			if v[1]%2 > 0 {
				p2 *= v[0]
			}
		}
		pfa[p2]++
		used[a] = p2
	}

	ans := c2(n) - c2(n-zero)
	for _, v := range pfa {
		ans += c2(v)
	}

	fmt.Fprint(out, ans)
}

func c2(a int) int {
	return a * (a - 1) / 2
}

// 素因数分解
func PrimeFactorization(x int) [][2]int {
	var primes [][2]int

	for i := 2; i*i <= x; i++ {
		if x%i != 0 {
			continue
		}

		var e int

		for x%i == 0 {
			e++
			x /= i
		}

		primes = append(primes, [2]int{i, e})
	}

	if x != 1 {
		primes = append(primes, [2]int{x, 1})
	}

	return primes
}

func main() {
	run(os.Stdin, os.Stdout)
}
