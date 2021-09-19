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
		{name: "1", args: args{r: strings.NewReader(`3 3
		1 2
		2 3
		3 1`)}, want: 2},
		{name: "2", args: args{r: strings.NewReader(`3 5
		1 1
		2 1
		2 2
		3 2
		3 3`)}, want: 3},
		{name: "3", args: args{r: strings.NewReader(`7 5
		1 7
		7 1
		3 4
		2 6
		5 2`)}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
