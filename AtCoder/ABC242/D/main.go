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
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	s := readString()
	q := readInt()
	ts := make([][]int, q)
	for i := range ts {
		t, k := readInt(), readInt()
		ts[i] = []int{t, k - 1}
	}

	// abc := make([][]string, 10)
	// for i := range abc {
	// 	if i == 0 {
	// 		abc[i] = []string{"A", "B", "C"}
	// 		continue
	// 	}
	// 	abc[i] = make([]string, 3)
	// 	for j, a := range abc[i-1] {
	// 		a = strings.ReplaceAll(a, "A", "1")
	// 		a = strings.ReplaceAll(a, "B", "2")
	// 		a = strings.ReplaceAll(a, "C", "AB")
	// 		a = strings.ReplaceAll(a, "1", "BC")
	// 		a = strings.ReplaceAll(a, "2", "CA")
	// 		abc[i][j] = a
	// 	}
	// 	p := abc[i][0]
	// 	fmt.Println(p)
	// }

	ans := make([]string, q)
	for i, tk := range ts {
		ans[i] = string(fq(tk[0], tk[1], s))
	}

	ansstr := fmt.Sprintf("%v", ans)
	return ansstr[1 : len(ansstr)-1]
}

func fq(t, k int, s string) byte {
	var ret byte
	if t == 0 {
		ret = s[k]
	} else if k == 0 {
		ret = f(s[0], t)
	} else {
		ret = f(fq(t-1, k/2, s), k%2+1)
	}
	return ret
}

func f(c byte, cnt int) byte {
	for i := 0; i < cnt%3; i++ {
		switch c {
		case 'A':
			c = 'B'
		case 'B':
			c = 'C'
		case 'C':
			c = 'A'
		}
	}
	return c
}

func main() {
	fmt.Println(run(os.Stdin))
}
