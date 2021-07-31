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
		{name: "1", args: args{r: strings.NewReader(`5
		1 3
		1 5
		3
		2 2
		3`)}, want: "3\n7"},
		{name: "2", args: args{r: strings.NewReader(`6
		1 1000000000
		2 1000000000
		2 1000000000
		2 1000000000
		2 1000000000
		3`)}, want: "5000000000"},
		{name: "3", args: args{r: strings.NewReader(`4
		1 3
		2 5
		1 2
		3`)}, want: "2"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
