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
		{name: "1", args: args{r: strings.NewReader(`5 4
		9 8 3 7 2
		1 6 2 9 5`)}, want: "Yes"},
		{name: "2", args: args{r: strings.NewReader(`4 90
		1 1 1 100
		1 2 3 100`)}, want: "No"},
		{name: "3", args: args{r: strings.NewReader(`4 1000000000
		1 1 1000000000 1000000000
		1 1000000000 1 1000000000`)}, want: "Yes"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
