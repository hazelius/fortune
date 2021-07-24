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
		{name: "1", args: args{r: strings.NewReader(`4 5
		2 4
		1 2
		2 3
		1 3
		3 4`)}, want: 2},
		{name: "2", args: args{r: strings.NewReader(`4 3
		1 3
		2 3
		2 4`)}, want: 1},
		{name: "3", args: args{r: strings.NewReader("2 0")}, want: 0},
		{name: "4", args: args{r: strings.NewReader(`7 8
		1 3
		1 4
		2 3
		2 4
		2 5
		2 6
		5 7
		6 7`)}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
