// https://atcoder.jp/contests/abc268/tasks/abc268_d

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
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
	sc.Split(bufio.ScanWords)

	n, m := readInt(), readInt()
	ss := make([]string, n)
	for i := range ss {
		ss[i] = readString()
	}
	tmap := make(map[string]bool)
	for i := 0; i < m; i++ {
		t := readString()
		tmap[t] = true
	}

	sort.Strings(ss)
	for {
		j := strings.Join(ss, "_")
		yo := 16 - len(j)
		ans, ok := dfs(ss[0], ss[1:], yo, tmap)
		if ok {
			fmt.Fprint(out, ans)
			return
		}

		if !NextPermutation(sort.StringSlice(ss)) {
			fmt.Fprint(out, -1)
			return
		}
	}
}

func dfs(name string, ss []string, yo int, tmap map[string]bool) (string, bool) {
	if len(ss) == 0 {
		if len(name) > 2 && !tmap[name] {
			return name, true
		}
		return "", false
	}

	join := "_"
	for i := 0; i <= yo; i++ {
		newname := fmt.Sprintf("%v%v%v", name, join, ss[0])
		ans, ok := dfs(newname, ss[1:], yo-i, tmap)
		if ok {
			return ans, ok
		}
		join += "_"
	}

	return "", false
}

func NextPermutation(x sort.Interface) bool {
	n := x.Len() - 1
	if n < 1 {
		return false
	}
	j := n - 1
	for ; !x.Less(j, j+1); j-- {
		if j == 0 {
			return false
		}
	}
	l := n
	for !x.Less(j, l) {
		l--
	}
	x.Swap(j, l)
	for k, l := j+1, n; k < l; {
		x.Swap(k, l)
		k++
		l--
	}
	return true
}

func main() {
	run(os.Stdin, os.Stdout)
}
