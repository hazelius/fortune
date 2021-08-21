// https://atcoder.jp/contests/abc215/tasks/abc215_c
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
)

var sc *bufio.Scanner

type ByRune []rune

func (r ByRune) Len() int           { return len(r) }
func (r ByRune) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r ByRune) Less(i, j int) bool { return r[i] < r[j] }

func StringToRuneSlice(s string) []rune {
	var r []rune
	for _, runeValue := range s {
		r = append(r, runeValue)
	}
	return r
}

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

	s, k := readString(), readInt()
	ss := StringToRuneSlice(s)
	sort.Sort(ByRune(ss))

	_, ans := f(0, k, ss)
	return string(ans)
}

func f(cnt, goal int, s []rune) (int, []rune) {
	if len(s) == 1 {
		return cnt + 1, s
	}

	arg := make([]rune, len(s)-1)
	for i, v := range s {
		if i > 0 && s[i-1] == v {
			continue
		}
		for j, c := range s {
			if j < i {
				arg[j] = c
			} else if j > i {
				arg[j-1] = c
			}
		}
		var rs []rune
		cnt, rs = f(cnt, goal, arg)
		if cnt == goal {
			rs = append([]rune{v}, rs...)
			return cnt, rs
		}
	}
	return cnt, s
}

func main() {
	fmt.Println(run(os.Stdin))
}
