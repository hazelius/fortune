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
		{name: "1", args: args{r: strings.NewReader(`5
		3 1 4 1 5
		4
		1 5 1
		2 4 3
		1 5 2
		1 3 3`)}, want: "2 0 0 1"},
		{name: "2", args: args{r: strings.NewReader(`10
		2 5 9 5 2 1 4 9 2 2
		2
		1 10 2
		2 10 2`)}, want: "4 3"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
