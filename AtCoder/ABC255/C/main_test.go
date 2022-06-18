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
		{name: "1", args: args{r: strings.NewReader("6 2 3 3")}, want: 1},
		{name: "2", args: args{r: strings.NewReader("0 0 0 1")}, want: 0},
		{name: "3", args: args{r: strings.NewReader("998244353 -10 -20 30")}, want: 998244363},
		{name: "4", args: args{r: strings.NewReader("-555555555555555555 -1000000000000000000 1000000 1000000000000")}, want: 444445},
		{name: "5", args: args{r: strings.NewReader("-10 10 -2 30")}, want: 0},
		{name: "6", args: args{r: strings.NewReader("10 -10 2 20")}, want: 0},
		{name: "7", args: args{r: strings.NewReader("10 -10 -2 20")}, want: 20},
		{name: "8", args: args{r: strings.NewReader("-10 10 2 20")}, want: 20},
		{name: "9", args: args{r: strings.NewReader("-10 -15 2 20")}, want: 1},
		{name: "9", args: args{r: strings.NewReader("2 -8 3 20")}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
