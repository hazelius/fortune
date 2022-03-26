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
		{name: "1", args: args{r: strings.NewReader(`8
		0 3 2 6 2 1 0 0`)}, want: 4},
		{name: "2", args: args{r: strings.NewReader(`3
		2000 2000 2000`)}, want: 0},
		{name: "3", args: args{r: strings.NewReader(`10
		0 1 2 3 4 5 6 7 8 10`)}, want: 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
