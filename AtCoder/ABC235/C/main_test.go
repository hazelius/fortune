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
		{name: "1", args: args{r: strings.NewReader(`6 8
		1 1 2 3 1 2
		1 1
		1 2
		1 3
		1 4
		2 1
		2 2
		2 3
		4 1`)}, want: "1 2 5 -1 3 6 -1 -1"},
		{name: "2", args: args{r: strings.NewReader(`3 2
		0 1000000000 999999999
		1000000000 1
		123456789 1`)}, want: "2 -1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
