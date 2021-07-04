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
		{name: "1", args: args{r: strings.NewReader(`2 7
		1 8`)}, want: "4\n3"},
		{name: "2", args: args{r: strings.NewReader(`1 3
33`)}, want: "3"},
		{name: "3", args: args{r: strings.NewReader(`7 1000000000000
		99 8 2 4 43 5 3`)}, want: `142857142857
142857142857
142857142858
142857142857
142857142857
142857142857
142857142857`},
		{name: "4", args: args{r: strings.NewReader(`3 10
		3 2 1`)}, want: "3\n3\n4"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
