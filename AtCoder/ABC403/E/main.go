package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

var sc *bufio.Scanner

// Node is a node of tree.
type Node struct {
	key      string
	delFlg   bool
	count    int
	children map[rune]*Node
}

// NewTrie is create a root node.
func NewTrie() *Node {
	return &Node{
		key:      "",
		delFlg:   false,
		count:    0,
		children: make(map[rune]*Node),
	}
}

func (n *Node) Insert(word string) bool {
	if n.delFlg {
		return false
	}
	curNode := n
	for _, r := range word {
		if _, ok := curNode.children[r]; !ok {
			curNode.children[r] = NewTrie()
		}
		curNode = curNode.children[r]
		if curNode.delFlg {
			return false
		}
	}
	curNode.key = word
	curNode.count++
	return true
}

func (n *Node) Delete(word string) int {
	if n.delFlg {
		return 0
	}
	curNode := n
	for _, r := range word {
		if _, ok := curNode.children[r]; !ok {
			curNode.children[r] = NewTrie()
		}
		curNode = curNode.children[r]
		if curNode.delFlg {
			return 0
		}
	}
	cnt := curNode.Count()
	curNode.key = word
	curNode.delFlg = true
	curNode.count = 0
	return cnt
}

func (n *Node) Count() int {
	cnt := 0
	if !n.delFlg && n.key != "" {
		cnt += n.count
	}
	for _, child := range n.children {
		if child.delFlg {
			continue
		}
		cnt += child.Count()
	}
	return cnt
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

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	q := readInt()
	nodes := NewTrie()
	cnt := 0
	for i := 0; i < q; i++ {
		t := readInt()
		s := readString()
		if t == 2 {
			if nodes.Insert(s) {
				cnt++
			}
		} else {
			cnt -= nodes.Delete(s)
		}
		fmt.Fprintln(out, cnt)
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
