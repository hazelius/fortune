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

func readString() string {
	sc.Scan()
	return sc.Text()
}

func run(r io.Reader) string {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	n := readInt()
	ss := readFigure(n)
	ts := readFigure(n)

	flg := true
	if len(ss) == len(ts) && len(ss[0]) == len(ts[0]) {
		for i := 0; i < len(ss); i++ {
			if ss[i] != ts[i] {
				flg = false
				break
			}
		}
	} else {
		flg = false
	}
	if flg {
		return "Yes"
	}

	flg = true
	if len(ss) == len(ts) && len(ss[0]) == len(ts[0]) {
		for i := 0; i < len(ss); i++ {
			for j := 0; j < len(ss[0]); j++ {
				i2 := len(ts) - 1 - i
				j2 := len(ts[0]) - 1 - j
				if i2 < 0 || j2 < 0 || ss[i][j] != ts[i2][j2] {
					flg = false
					break
				}
			}
			if !flg {
				break
			}
		}
	} else {
		flg = false
	}

	if flg {
		return "Yes"
	}

	flg = true
	if len(ss) == len(ts[0]) && len(ss[0]) == len(ts) {
		for i := 0; i < len(ss); i++ {
			for j := 0; j < len(ss[0]); j++ {
				i2 := len(ts) - 1 - j
				j2 := i
				if i2 < 0 || j2 >= len(ts[0]) || ss[i][j] != ts[i2][j2] {
					flg = false
					break
				}
			}
			if !flg {
				break
			}
		}
	} else {
		flg = false
	}
	if flg {
		return "Yes"
	}

	flg = true
	if len(ss) == len(ts[0]) && len(ss[0]) == len(ts) {
		for i := 0; i < len(ss); i++ {
			for j := 0; j < len(ss[0]); j++ {
				i2 := j
				j2 := len(ts[0]) - 1 - i
				// fmt.Printf("%v %v, %v %v, %v %v\n", i, j, i2, j2, ss[i][j], ts[i2][j2])
				if i2 >= len(ts) || j2 < 0 || ss[i][j] != ts[i2][j2] {
					flg = false
					break
				}
			}
			if !flg {
				break
			}
		}
	} else {
		flg = false
	}
	if flg {
		return "Yes"
	}

	return "No"
}

func readFigure(n int) []string {
	ret := make([]string, 0, n)
	rgt, lft := n, 0
	for i := 0; i < n; i++ {
		str := readString()

		flg := false
		for i, c := range str {
			if c == '#' {
				flg = true
				if rgt > i {
					rgt = i
				}
				if lft < i {
					lft = i
				}
			}
		}
		if flg {
			ret = append(ret, str)
		}
	}

	for i, s := range ret {
		ret[i] = s[rgt : lft+1]
	}
	return ret
}

func main() {
	fmt.Println(run(os.Stdin))
}
