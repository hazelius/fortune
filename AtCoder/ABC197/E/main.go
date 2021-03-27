package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(xcs [][]int) int {
	c, r := 0, 0
	s := sortxc(xcs)
	fmt.Println(s)
	for i, val := range s {
		cmin := val[0]
		cmax := val[1]
		if c <= cmin && c <= cmax {
			r += abs(cmax - c)
			c = cmax
		} else if c >= cmin && c >= cmax {
			r += abs(cmin - c)
			c = cmin
		} else {
			nmin := 0
			nmax := 0
			if i+1 > len(s) {
				nmin = s[i+1][0]
				nmax = s[i+1][1]
			}

			rmin := rout(c, cmin, cmax, nmin, nmax)
			rmax := rout(c, cmax, cmin, nmin, nmax)
			if rmin > rmax {
				r += abs(c - cmax)
				c = cmin
			} else {
				r += abs(c - cmin)
				c = cmax
			}
			r += abs(cmin - cmax)
		}
		fmt.Printf("c:%v, r:%v\n", c, r)
	}
	return r + c
}

func rout(r1, r2, r3, r4, r5 int) int {
	a := abs(r2-r1) + abs(r3-r2) + abs(r4-r3) + abs(r5-r4)
	b := abs(r2-r1) + abs(r3-r2) + abs(r5-r3) + abs(r4-r5)
	return min(a, b)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}

func sortxc(xcs [][]int) [][]int {
	r := make([][]int, 0, len(xcs))
	for _, v := range xcs {
		flg := true
		for ri, rv := range r {
			if rv[2] > v[1] {
				tmp := make([][]int, len(r)+1)
				copy(tmp, r[:ri])
				tmp[ri] = []int{v[0], v[0], v[1]}
				copy(tmp[ri+1:], r[ri:])
				r = tmp
				flg = false
				break
			}

			if rv[2] == v[1] {
				if rv[0] > v[0] {
					rv[0] = v[0]
				} else if rv[1] < v[0] {
					rv[1] = v[0]
				}
				flg = false
				break
			}
		}
		if flg {
			r = append(r, []int{v[0], v[0], v[1]})
		}
	}
	return r
}

func main() {
	sc.Split(bufio.ScanWords)
	n := readInt()
	xcs := make([][]int, n)
	for i := 0; i < n; i++ {
		xcs[i] = []int{readInt(), readInt()}
	}
	fmt.Println(run(xcs))
}
