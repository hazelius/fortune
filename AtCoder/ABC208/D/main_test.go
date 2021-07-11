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
		{name: "1", args: args{r: strings.NewReader(`3 2
		1 2 3
		2 3 2`)}, want: 25},
		{name: "2", args: args{r: strings.NewReader("3 0")}, want: 0},
		{name: "3", args: args{r: strings.NewReader(`5 20
		1 2 6
		1 3 10
		1 4 4
		1 5 1
		2 1 5
		2 3 9
		2 4 8
		2 5 6
		3 1 5
		3 2 1
		3 4 7
		3 5 9
		4 1 4
		4 2 6
		4 3 4
		4 5 8
		5 1 2
		5 2 5
		5 3 6
		5 4 5`)}, want: 517},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
