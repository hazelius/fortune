package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

type panel [3][3]int

var sc *bufio.Scanner

var memo map[panel]int

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	var p panel
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			p[i][j] = readInt()
		}
	}

	memo = map[panel]int{}
	var n panel
	ans := read(n, p, 1)

	if ans > 0 {
		fmt.Fprint(out, "Takahashi")
	} else {
		fmt.Fprint(out, "Aoki")
	}
}

func read(n, p panel, flg int) int {
	ret, ok := memo[n]
	if ok {
		return ret
	}

	ret = check(n, p)
	memo[n] = ret
	if ret != 0 {
		return ret
	}

	vic := false
	for i, row := range n {
		for j, clm := range row {
			if clm != 0 {
				continue
			}
			n[i][j] = flg
			ret = read(n, p, -flg)
			if ret == flg {
				vic = true
			}
			memo[n] = ret
			n[i][j] = 0
		}
	}

	if vic {
		return flg
	}
	return -flg
}

func check(n, p panel) int {
	for i := 0; i < 3; i++ {
		if n[i][0] == n[i][1] && n[i][1] == n[i][2] {
			if n[i][0] != 0 {
				return n[i][0]
			}
		}
		if n[0][i] == n[1][i] && n[1][i] == n[2][i] {
			if n[0][i] != 0 {
				return n[0][i]
			}
		}
	}
	if n[0][0] == n[1][1] && n[1][1] == n[2][2] {
		if n[1][1] != 0 {
			return n[1][1]
		}
	}
	if n[2][0] == n[1][1] && n[1][1] == n[0][2] {
		if n[1][1] != 0 {
			return n[1][1]
		}
	}

	t, a := 0, 0
	for i, row := range p {
		for j, score := range row {
			if n[i][j] > 0 {
				t += score
			} else if n[i][j] < 0 {
				a += score
			} else {
				return 0
			}
		}
	}

	if t > a {
		return 1
	}
	return -1
}

func main() {
	run(os.Stdin, os.Stdout)
}
