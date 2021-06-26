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
		{name: "1", args: args{strings.NewReader("5 2 3 2")}, want: 2},
		{name: "2", args: args{strings.NewReader("6 9 2 3")}, want: -1},
		{name: "3", args: args{strings.NewReader("0 2 8 3")}, want: 0},
		{name: "3", args: args{strings.NewReader("1000000 7 2 4")}, want: 1000000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
