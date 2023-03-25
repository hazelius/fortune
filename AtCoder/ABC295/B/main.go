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

func readString() []byte {
	sc.Scan()
	return sc.Bytes()
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	r, c := readInt(), readInt()
	bs := make([][]byte, r)
	for i := range bs {
		bs[i] = readString()
	}

	ans := make([][]byte, r)
	for i := range ans {
		ans[i] = make([]byte, c)
		for j := range ans[i] {
			ans[i][j] = byte('#')
		}
	}

	for h, b := range bs {
		for i := 0; i < c; i++ {
			m := string(b[i])
			if m == "." {
				ans[h][i] = byte('.')
			}
			bomb, err := strconv.Atoi(m)
			if err != nil {
				continue
			}
			ans[h][i] = byte('.')

			used := make([][]bool, r)
			for i := range used {
				used[i] = make([]bool, c)
			}

			qs := [][]int{{h, i + 1, 1}, {h, i - 1, 1}, {h + 1, i, 1}, {h - 1, i, 1}}
			for len(qs) > 0 {
				q := qs[0]
				qs = qs[1:]
				x := q[0]
				y := q[1]
				cnt := q[2]
				if x < 0 || x >= r || y < 0 || y >= c {
					continue
				}
				if used[x][y] {
					continue
				}
				used[x][y] = true
				ans[x][y] = byte('.')

				cnt++
				if cnt <= bomb {
					qs = append(qs, []int{x, y + 1, cnt}, []int{x, y - 1, cnt}, []int{x + 1, y, cnt}, []int{x - 1, y, cnt})
				}
			}
		}
	}

	for _, b := range ans {
		fmt.Fprintln(out, string(b))
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
