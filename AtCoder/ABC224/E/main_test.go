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
		{name: "1", args: args{r: strings.NewReader(`3 3 7
		1 1 4
		1 2 7
		2 1 3
		2 3 5
		3 1 2
		3 2 5
		3 3 5`)}, want: "1 0 2 0 3 1 0"},
		{name: "2", args: args{r: strings.NewReader(`5 7 20
		2 7 8
		2 6 4
		4 1 9
		1 5 4
		2 2 7
		5 5 2
		1 7 2
		4 6 6
		1 4 1
		2 1 10
		5 6 9
		5 3 3
		3 7 9
		3 6 3
		4 3 4
		3 3 10
		4 2 1
		3 5 4
		1 2 6
		4 7 9`)}, want: "2 4 1 5 3 6 6 2 7 0 0 4 1 5 3 0 5 2 4 0"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
