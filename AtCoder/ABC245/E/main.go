// E - Wrapping Chocolate
// https://atcoder.jp/contests/abc245/tasks/abc245_e
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"

	rbt "github.com/emirpasic/gods/trees/redblacktree"
)

var sc *bufio.Scanner

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

type box struct {
	tate  int
	yoko  int
	isBox bool
}

type boxes []box

func (a boxes) Len() int      { return len(a) }
func (a boxes) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a boxes) Less(i, j int) bool {
	if a[i].tate == a[j].tate {
		return a[i].isBox
	}
	return a[i].tate > a[j].tate
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	n, m := readInt(), readInt()
	arr := make(boxes, n+m)
	for i := 0; i < n; i++ {
		arr[i].tate = readInt()
	}
	for i := 0; i < n; i++ {
		arr[i].yoko = readInt()
	}
	for i := 0; i < m; i++ {
		arr[i+n].tate = readInt()
		arr[i+n].isBox = true
	}
	for i := 0; i < m; i++ {
		arr[i+n].yoko = readInt()
		arr[i+n].isBox = true
	}

	sort.Sort(arr)
	tree := rbt.NewWithIntComparator()
	for _, b := range arr {
		if b.isBox {
			// 赤黒木に同じキーで複数登録はできないのでカウンターを保存する
			cnt := 1
			v, ok := tree.Get(b.yoko)
			if ok {
				cnt = v.(int) + 1
			}
			tree.Put(b.yoko, cnt)
			continue
		}

		v, ok := tree.Ceiling(b.yoko)
		if !ok {
			fmt.Fprint(out, "No")
			return
		}
		cnt := v.Value.(int)
		if cnt == 1 {
			tree.Remove(v.Key)
		} else {
			tree.Put(v.Key, cnt-1)
		}
	}

	fmt.Fprint(out, "Yes")
}

func main() {
	run(os.Stdin, os.Stdout)
}
