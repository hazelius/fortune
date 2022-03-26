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
		{name: "1", args: args{r: strings.NewReader(`1 2
		2 1
		12 14 8 2`)}, want: "6 4 2"},
		{name: "2", args: args{r: strings.NewReader(`1 1
		100 1
		10000 0 -1`)}, want: "100 -1"},
		{name: "3", args: args{r: strings.NewReader(`2 3
		5 4 3
		45 36 62 58 39 19`)}, want: "9 0 7 6"},
		{name: "4", args: args{r: strings.NewReader(`1 1
		1 0
		1 0 0`)}, want: "1 0"},
		{name: "5", args: args{r: strings.NewReader(`3 3
		1 0 0 0
		1 0 0 1 0 0 0`)}, want: "1 0 0 1"},
		{name: "6", args: args{r: strings.NewReader(`2 2
		1 2 0 
		1 4 4 0 0`)}, want: "1 2 0"},
		{name: "7", args: args{r: strings.NewReader(`2 2
		-1 -2 -3
		-1 -4 -10 -12 -9`)}, want: "1 2 3"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
