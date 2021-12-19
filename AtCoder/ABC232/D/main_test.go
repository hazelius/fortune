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
		{name: "1", args: args{r: strings.NewReader(`3 4
		.#..
		..#.
		..##`)}, want: 4},
		{name: "2", args: args{r: strings.NewReader(`1 1
		.`)}, want: 1},
		{name: "3", args: args{r: strings.NewReader(`5 5
		.....
		.....
		.....
		.....
		.....`)}, want: 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
