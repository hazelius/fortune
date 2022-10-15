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

type Node struct {
	value    int
	previous *Node
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	q := readInt()
	nodemap := make(map[int]*Node)

	preNode := &Node{value: -1}

	for i := 0; i < q; i++ {
		switch readString() {
		case "ADD":
			x := readInt()
			nn := &Node{value: x, previous: preNode}
			preNode = nn
		case "DELETE":
			if preNode.previous != nil {
				preNode = preNode.previous
			}
		case "SAVE":
			y := readInt()
			nodemap[y] = preNode
		case "LOAD":
			z := readInt()
			savedQ, ok := nodemap[z]
			if ok {
				preNode = savedQ
			} else {
				preNode = &Node{value: -1}
			}
		}
		fmt.Fprint(out, preNode.value)
		fmt.Fprint(out, " ")
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
