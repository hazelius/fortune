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
		{name: "1", args: args{r: strings.NewReader(`4 3 1
		1 2 3 1`)}, want: 4},
		{name: "2", args: args{r: strings.NewReader(`5 2 1
		1 3 2 4 1`)}, want: 0},
		{name: "3", args: args{r: strings.NewReader(`5 1 1
		1 1 1 1 1`)}, want: 15},
		{name: "4", args: args{r: strings.NewReader(`10 8 1
		2 7 1 8 2 8 1 8 2 8`)}, want: 36},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
