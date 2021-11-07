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
		3 0
		5 1 1
		7 1 1`)}, want: 10},
		{name: "2", args: args{r: strings.NewReader(`5
		1000000000 0
		1000000000 0
		1000000000 0
		1000000000 0
		1000000000 4 1 2 3 4`)}, want: 5000000000},
		{name: "3", args: args{r: strings.NewReader(`5
		1 0
		2 1 1
		3 2 1 2
		4 3 1 2 3
		5 4 1 2 3 4`)}, want: 15},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
