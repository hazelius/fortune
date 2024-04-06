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

type point struct {
	x int
	y int
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	ah, aw := readInt(), readInt()
	mapa := readSheet(ah)
	bh, bw := readInt(), readInt()
	mapb := readSheet(bh)
	xh, xw := readInt(), readInt()
	mapx := readSheet(xh)

	for x := -ah + 1; x < xh; x++ {
		for y := -aw + 1; y < xw; y++ {
		loop:
			for x2 := -bh + 1; x2 < xh; x2++ {
				for y2 := -bw + 1; y2 < xw; y2++ {
					for k := range mapx {
						mapx[k] = false
					}

					flg := true
					for ap := range mapa {
						ax := ap.x + x
						ay := ap.y + y
						if _, ok := mapx[point{ax, ay}]; !ok {
							flg = false
							break
						}
						mapx[point{ax, ay}] = true
					}
					if !flg {
						break loop
					}

					for bp := range mapb {
						bx := bp.x + x2
						by := bp.y + y2
						if _, ok := mapx[point{bx, by}]; !ok {
							flg = false
							break
						}
						mapx[point{bx, by}] = true
					}
					if !flg {
						continue
					}

					flg = true
					for k := range mapx {
						if !mapx[k] {
							flg = false
							break
						}
					}
					if flg {
						fmt.Fprint(out, "Yes")
						return
					}
				}
			}
		}
	}

	fmt.Fprint(out, "No")
}

func readSheet(h int) map[point]bool {
	ret := make(map[point]bool)
	for i := 0; i < h; i++ {
		a := readString()
		for j, v := range a {
			if v == '#' {
				ret[point{i, j}] = true
			}
		}
	}
	return ret
}

func main() {
	run(os.Stdin, os.Stdout)
}
