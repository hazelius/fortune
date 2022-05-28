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
		{name: "1", args: args{r: strings.NewReader(`8
		1 3
		1 2
		3
		1 2
		1 7
		3
		2 2 3
		3`)}, want: "1 5 4"},
		{name: "2", args: args{r: strings.NewReader(`4
		1 10000
		1 1000
		2 100 3
		1 10`)}, want: ""},
		{name: "3", args: args{r: strings.NewReader(`10
		1 2
		3
		1 3 
		1 10 
		1 10 
		3
		2 2 1 
		2 10 1
		3
		3`)}, want: "0 8 7 7"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
