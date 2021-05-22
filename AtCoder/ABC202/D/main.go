package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

var c map[int]map[int]int

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(a, b, k int) string {
	c = make(map[int]map[int]int)
	return f("", a, b, k)
}

func combi(a, l int) int {
	if v, ok := c[a][l]; ok {
		return v
	}

	var ret float64 = 1
	for i := l; i > 1; i-- {
		if i > (l - a) {
			ret *= float64(i)
		}
		if i <= a {
			ret /= float64(i)
		}
	}
	if _, ok := c[a]; !ok {
		c[a] = make(map[int]int)
	}
	c[a][l] = int(ret)
	return c[a][l]
}

func f(x string, a, b, k int) string {
	if a == 0 && b == 0 {
		return x
	}
	if a == 0 {
		x += "b"
		b--
		return f(x, a, b, k)
	}
	if b == 0 {
		x += "a"
		a--
		return f(x, a, b, k)
	}

	tmp := combi(a-1, a-1+b)
	if tmp < k {
		x += "b"
		b--
		k -= tmp
	} else {
		x += "a"
		a--
	}
	return f(x, a, b, k)
}

func main() {
	sc.Split(bufio.ScanWords)
	a, b, k := readInt(), readInt(), readInt()
	fmt.Println(run(a, b, k))
}
