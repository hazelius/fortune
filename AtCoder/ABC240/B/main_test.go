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
		{name: "1", args: args{r: strings.NewReader(`6
		1 4 1 2 2 1`)}, want: 3},
		{name: "2", args: args{r: strings.NewReader(`1 1`)}, want: 1},
		{name: "3", args: args{r: strings.NewReader(`11
		3 1 4 1 5 9 2 6 5 3 5`)}, want: 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
