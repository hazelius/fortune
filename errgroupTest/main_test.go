package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	args = []string{"1", "2", "3"}
	main()

	args = []string{"1", "a", "b"}
	main()
}
