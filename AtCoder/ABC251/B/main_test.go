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
		{name: "1", args: args{r: strings.NewReader(`2 10
		1 3`)}, want: 3},
		{name: "2", args: args{r: strings.NewReader(`2 1
		2 3`)}, want: 0},
		{name: "3", args: args{r: strings.NewReader(`4 12
		3 3 3 3`)}, want: 3},
		{name: "4", args: args{r: strings.NewReader(`7 251
		202 20 5 1 4 2 100`)}, want: 48},
		{name: "5", args: args{r: strings.NewReader(`5 8
		10 5 2 8 2`)}, want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
