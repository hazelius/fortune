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
		1 2
		2 3
		3 1`)}, want: 6},
		{name: "2", args: args{r: strings.NewReader("3 0")}, want: 27},
		{name: "3", args: args{r: strings.NewReader(`4 6
		1 2
		2 3
		3 4
		2 4
		1 3
		1 4`)}, want: 0},
		{name: "4", args: args{r: strings.NewReader("20 0")}, want: 3486784401},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
