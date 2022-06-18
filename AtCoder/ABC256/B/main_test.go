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
		{name: "1", args: args{r: strings.NewReader(`4
		1 1 3 2`)}, want: 3},
		{name: "2", args: args{r: strings.NewReader(`3
		1 1 1`)}, want: 0},
		{name: "3", args: args{r: strings.NewReader(`10
		2 2 4 1 1 1 4 2 2 1`)}, want: 8},
		{name: "4", args: args{r: strings.NewReader(`4
		1 1 4 4`)}, want: 4},
		{name: "5", args: args{r: strings.NewReader(`4
		1 1 1 3`)}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
