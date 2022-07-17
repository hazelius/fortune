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
		{name: "1", args: args{r: strings.NewReader("2 3 4")}, want: 12},
		{name: "2", args: args{r: strings.NewReader("1 5 5")}, want: 0},
		{name: "3", args: args{r: strings.NewReader("10 5 5")}, want: 3942349900},
		{name: "4", args: args{r: strings.NewReader("10 1 1")}, want: 2584},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
