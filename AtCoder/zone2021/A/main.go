// HELLO SPACE
//

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func readString() string {
	sc.Scan()
	return sc.Text()
}

func run(s string) int {
	return strings.Count(s, "ZONe")
}

func main() {
	sc.Split(bufio.ScanWords)
	s := readString()
	fmt.Println(run(s))
}
