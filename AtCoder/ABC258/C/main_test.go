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
		want string
	}{
		{name: "1", args: args{r: strings.NewReader(`3 3
		abc
		2 2
		1 1
		2 2`)}, want: "b a"},
		{name: "2", args: args{r: strings.NewReader(`10 8
		dsuccxulnl
		2 4
		2 7
		1 2
		2 7
		1 1
		1 2
		1 3
		2 5`)}, want: "c u c u"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
