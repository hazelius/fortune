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
		{name: "1", args: args{r: strings.NewReader(`3 2
		1 1 3
		3 1`)}, want: "Yes"},
		{name: "2", args: args{r: strings.NewReader(`1 1
		1000000000
		1`)}, want: "No"},
		{name: "3", args: args{r: strings.NewReader(`5 2
		1 2 3 4 5
		5 5`)}, want: "No"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
