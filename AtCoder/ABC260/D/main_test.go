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
		{name: "1", args: args{r: strings.NewReader(`5 2
		3 5 2 1 4`)}, want: "4 3 3 -1 4"},
		{name: "2", args: args{r: strings.NewReader(`5 1
		1 2 3 4 5`)}, want: "1 2 3 4 5"},
		{name: "3", args: args{r: strings.NewReader(`15 3
		3 14 15 9 2 6 5 13 1 7 10 11 8 12 4`)}, want: "9 9 9 15 15 6 -1 -1 6 -1 -1 -1 -1 6 15"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
