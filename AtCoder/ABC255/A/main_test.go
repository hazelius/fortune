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
		{name: "1", args: args{r: strings.NewReader(`1 2
		1 0
		0 1`)}, want: 0},
		{name: "2", args: args{r: strings.NewReader(`2 2
		1 2
		3 4`)}, want: 4},
		{name: "3", args: args{r: strings.NewReader(`2 1
		90 80
		70 60`)}, want: 70},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}