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
		3 4 1
		7 0 0
		0 0 7
		0 0 0
		1000000000000000 1000000000000000 1000000000000000`)}, want: "2 1 0 0 900000000000000"},
		{name: "2", args: args{r: strings.NewReader(`6
		1 1 1
		1 2 1
		1 4 3
		16 4 4
		24 4 5
		17 19 1`)}, want: "0 1 2 6 8 9"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
