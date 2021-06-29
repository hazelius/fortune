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
		0 0
		0 1
		1 0
		2 0
		3 0
		3 1`)}, want: "Yes"},
		{name: "2", args: args{r: strings.NewReader(`3
		1 0
		1 1
		3 0
		-1 0
		-1 1
		-3 0`)}, want: "No"},
		{name: "3", args: args{r: strings.NewReader(`4
		0 0
		2 9
		10 -2
		-6 -7
		0 0
		2 9
		10 -2
		-6 -7`)}, want: "Yes"},
		{name: "4", args: args{r: strings.NewReader(`6
		10 5
		-9 3
		1 -5
		-6 -5
		6 9
		-9 0
		-7 -10
		-10 -5
		5 4
		9 0
		0 -10
		-10 -2`)}, want: "Yes"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
