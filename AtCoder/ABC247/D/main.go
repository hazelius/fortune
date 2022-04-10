package main

import (
	"bufio"
	"container/list"
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

func run(r io.Reader) string {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	q := readInt()
	deque := list.New()
	ans := make([]int, 0, q)
	for i := 0; i < q; i++ {
		if readInt() == 1 {
			x, c := readInt(), readInt()
			deque.PushBack([]int{x, c})
		} else {
			c := readInt()
			a := 0
			for i := c; i > 0; {
				xc := deque.Front().Value.([]int)
				x := xc[0]
				tmpc := xc[1]
				if i >= tmpc {
					a += x * tmpc
					i -= tmpc
					deque.Remove(deque.Front())
				} else {
					a += x * i
					deque.Remove(deque.Front())
					deque.PushFront([]int{x, tmpc - i})
					i = 0
				}
			}
			ans = append(ans, a)
		}
	}

	ansstr := fmt.Sprintf("%v", ans)
	return ansstr[1 : len(ansstr)-1]
}

func main() {
	fmt.Println(run(os.Stdin))
}
