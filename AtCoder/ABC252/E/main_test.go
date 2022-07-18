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
		1 2 1
		2 3 2
		1 3 10`)}, want: "1 2"},
		{name: "2", args: args{r: strings.NewReader(`4 6
		1 2 1
		1 3 1
		1 4 1
		2 3 1
		2 4 1
		3 4 1`)}, want: "1 2 3"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
