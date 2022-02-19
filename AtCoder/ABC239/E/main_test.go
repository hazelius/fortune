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
		{name: "1", args: args{r: strings.NewReader(`5 2
		1 2 3 4 5
		1 4
		2 1
		2 5
		3 2
		1 2
		2 1`)}, want: "4 5"},
		// {name: "2", args: args{r: strings.NewReader(`6 2
		// 10 10 10 9 8 8
		// 1 4
		// 2 1
		// 2 5
		// 3 2
		// 6 4
		// 1 4
		// 2 2`)}, want: "9 10"},
		// {name: "3", args: args{r: strings.NewReader(`4 4
		// 1 10 100 1000
		// 1 2
		// 2 3
		// 3 4
		// 1 4
		// 2 3
		// 3 2
		// 4 1`)}, want: "1 10 100 1000"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
