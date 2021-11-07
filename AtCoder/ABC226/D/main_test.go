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
		want int
	}{
		{name: "1", args: args{r: strings.NewReader(`3
		1 2
		3 6
		7 4`)}, want: 6},
		{name: "2", args: args{r: strings.NewReader(`3
		1 2
		2 2
		4 2`)}, want: 2},
		{name: "3", args: args{r: strings.NewReader(`4
		0 0
		0 1000000000
		1000000000 0
		1000000000 1000000000`)}, want: 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
