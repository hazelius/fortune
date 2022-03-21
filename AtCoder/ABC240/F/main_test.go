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
		{name: "1", args: args{r: strings.NewReader(`3
		3 7
		-1 2
		2 3
		-3 2
		10 472
		-4 12
		1 29
		2 77
		-1 86
		0 51
		3 81
		3 17
		-2 31
		-4 65
		4 23
		1 1000000000
		4 1000000000`)}, want: "4 53910 2000000002000000000"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
