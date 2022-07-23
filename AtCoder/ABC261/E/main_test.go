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
		{name: "1", args: args{r: strings.NewReader(`3 10
		3 3
		2 5
		1 12`)}, want: "9 15 12"},
		{name: "2", args: args{r: strings.NewReader(`9 12
		1 1
		2 2
		3 3
		1 4
		2 5
		3 6
		1 7
		2 8
		3 9`)}, want: "0 2 1 0 5 3 3 11 2"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
