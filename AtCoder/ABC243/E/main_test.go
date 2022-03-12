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
		want int
	}{
		{name: "1", args: args{r: strings.NewReader(`3 3
		1 2 2
		2 3 3
		1 3 6`)}, want: 1},
		{name: "2", args: args{r: strings.NewReader(`5 4
		1 3 3
		2 3 9
		3 5 3
		4 5 3`)}, want: 0},
		{name: "3", args: args{r: strings.NewReader(`5 10
		1 2 71
		1 3 9
		1 4 82
		1 5 64
		2 3 22
		2 4 99
		2 5 1
		3 4 24
		3 5 18
		4 5 10`)}, want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
