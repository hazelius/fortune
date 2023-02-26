// https://atcoder.jp/contests/abc291/tasks/abc291_e
// トポロジカルソート
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

	n, m := readInt(), readInt()
	xymap := make(map[int][]int)
	inputs := make([]int, n)
	for i := 0; i < m; i++ {
		x, y := readInt()-1, readInt()-1
		xymap[x] = append(xymap[x], y)
		inputs[y]++
	}

	ans := make([]int, n)
	idx := 1

	que := make([]int, 0)
	for y, cnt := range inputs {
		if cnt == 0 {
			que = append(que, y)
		}
	}

	for len(que) > 0 {
		if len(que) > 1 {
			fmt.Fprint(out, "No")
			return
		}

		s := que[0]
		que = que[1:]

		ans[s] = idx
		idx++

		ys, ok := xymap[s]
		if !ok {
			continue
		}
		for _, y := range ys {
			inputs[y]--
			if inputs[y] == 0 {
				que = append(que, y)
			}
		}
	}

	fmt.Fprintln(out, "Yes")
	ansstr := fmt.Sprintf("%v", ans)
	fmt.Fprint(out, ansstr[1:len(ansstr)-1])
}

func main() {
	run(os.Stdin, os.Stdout)
}
