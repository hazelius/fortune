package main

import (
	"fmt"
	"strconv"

	"golang.org/x/sync/errgroup"
)

var args = []string{}

func main() {
	fmt.Println("main start")
	eg := errgroup.Group{}
	for _, str := range args {
		str := str
		eg.Go(func() error {
			return sub(str)
		})
	}

	if err := eg.Wait(); err != nil {
		fmt.Println(err)
	}
	fmt.Println("main end")
}

func sub(str string) error {
	num, err := strconv.Atoi(str)
	fmt.Println(num)
	return err
}
