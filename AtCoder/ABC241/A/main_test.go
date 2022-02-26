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
		{name: "1", args: args{r: strings.NewReader("9 0 1 2 3 4 5 6 7 8")}, want: 7},
		{name: "2", args: args{r: strings.NewReader("4 8 8 8 0 8 8 8 8 8")}, want: 4},
		{name: "3", args: args{r: strings.NewReader("0 0 0 0 0 0 0 0 0 0")}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
