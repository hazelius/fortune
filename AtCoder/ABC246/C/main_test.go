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
		{name: "1", args: args{r: strings.NewReader(`5 4 7
		8 3 10 5 13`)}, want: 12},
		{name: "2", args: args{r: strings.NewReader(`5 100 7
		8 3 10 5 13`)}, want: 0},
		{name: "3", args: args{r: strings.NewReader(`20 815 60
		2066 3193 2325 4030 3725 1669 1969 763 1653 159 5311 5341 4671 2374 4513 285 810 742 2981 202`)}, want: 112},
		{name: "4", args: args{r: strings.NewReader(`5 15 10
		50 40 30 20 11`)}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
