package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var sc *bufio.Scanner

func readString() string {
	sc.Scan()
	return sc.Text()
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	s := readString()

	startIdxs := make([]int, 0)
	kakko := [][]rune{{'(', ')'}, {'[', ']'}, {'<', '>'}}
	for i, c := range s {
		for _, ks := range kakko {
			if c == ks[0] {
				startIdxs = append(startIdxs, i)
				break
			}
			if c == ks[1] {
				if len(startIdxs) == 0 {
					fmt.Fprint(out, "No")
					return
				}
				lastIdx := startIdxs[len(startIdxs)-1]
				if ks[0] == rune(s[lastIdx]) {
					startIdxs = startIdxs[:len(startIdxs)-1]
					break
				}
			}
		}
	}

	if len(startIdxs) == 0 {
		fmt.Fprint(out, "Yes")
	} else {
		fmt.Fprint(out, "No")
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
