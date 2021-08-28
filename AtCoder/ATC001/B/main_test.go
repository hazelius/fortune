// B - Union Find
// https://atcoder.jp/contests/atc001/tasks/unionfind_a
package main

import (
	"io"
	"strings"
	"testing"
)

func Test_run(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{r: strings.NewReader(`8 9
		0 1 2
		0 3 2
		1 1 3
		1 1 4
		0 2 4
		1 4 1
		0 4 2
		0 0 0
		1 0 0`)}, want: "Yes\nNo\nYes\nYes"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
