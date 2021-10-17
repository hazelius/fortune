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
		1 1
		2 1
		3 1`)}, want: "3.000000000000000"},
		{name: "2", args: args{r: strings.NewReader(`3
		1 3
		2 2
		3 1`)}, want: "3.833333333333333"},
		{name: "3", args: args{r: strings.NewReader(`5
		3 9
		1 2
		4 6
		1 5
		5 3`)}, want: "8.916666666666668"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
