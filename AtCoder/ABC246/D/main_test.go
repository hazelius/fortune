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
		{name: "1", args: args{r: strings.NewReader("9")}, want: 15},
		{name: "2", args: args{r: strings.NewReader("0")}, want: 0},
		{name: "3", args: args{r: strings.NewReader("999999999989449206")}, want: 1000000000000000000},
		{name: "4", args: args{r: strings.NewReader("1")}, want: 1},
		{name: "5", args: args{r: strings.NewReader("2")}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
