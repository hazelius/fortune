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

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	m := readInt()
	ss := []string{readString(), readString(), readString()}
	smap := make([]map[rune][]int, 3)
	for i := range smap {
		smap[i] = make(map[rune][]int)
	}
	for i, s := range ss {
		for j, c := range s {
			smap[i][c] = append(smap[i][c], j)
		}
	}

	ans := -1
	for key, vs1 := range smap[0] {
		vs2, ok := smap[1][key]
		if !ok {
			continue
		}
		vs3, ok := smap[2][key]
		if !ok {
			continue
		}

		for i, v1 := range vs1 {
			if i > 2 {
				break
			}
			for j, v2 := range vs2 {
				if j > 2 {
					break
				}
				if v1 == v2 {
					v2 += m
				}

				for k, v3 := range vs3 {
					if k > 2 {
						break
					}
					if v1 == v3 || v2 == v3 {
						v3 += m
						if v1 == v3 || v2 == v3 {
							v3 += m
						}
					}

					t := v1
					if t < v2 {
						t = v2
					}
					if t < v3 {
						t = v3
					}
					if ans < 0 || ans > t {
						ans = t
					}
				}
			}
		}
	}

	fmt.Fprint(out, ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
