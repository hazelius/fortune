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

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func readString() string {
	sc.Scan()
	return sc.Text()
}

type person struct {
	s string
	c int
}

type persons []person

func (a persons) Len() int           { return len(a) }
func (a persons) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a persons) Less(i, j int) bool { return a[i].s < a[j].s }

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	n := readInt()
	ps := make(persons, n)
	sm := 0
	for i := range ps {
		ps[i] = person{readString(), readInt()}
		sm += ps[i].c
	}
	sort.Sort(ps)
	sm %= n

	fmt.Fprint(out, ps[sm].s)
}

func main() {
	run(os.Stdin, os.Stdout)
}
