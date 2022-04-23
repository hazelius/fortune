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
		{name: "1", args: args{r: strings.NewReader(`3
		6 2 3`)}, want: 2},
		{name: "2", args: args{r: strings.NewReader(`1
		2`)}, want: 0},
		{name: "3", args: args{r: strings.NewReader(`10
		1 3 2 4 6 8 2 2 3 7`)}, want: 62},
		{name: "4", args: args{r: strings.NewReader(`3
		1 1 1`)}, want: 27},
		{name: "5", args: args{r: strings.NewReader(`3
		1 3 3`)}, want: 9},
		{name: "6", args: args{r: strings.NewReader(`3
		2 2 4`)}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
