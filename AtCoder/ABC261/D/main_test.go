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
		{name: "1", args: args{r: strings.NewReader(`6 3
		2 7 1 8 2 8
		2 10
		3 1
		5 5`)}, want: 48},
		{name: "2", args: args{r: strings.NewReader(`3 2
		1000000000 1000000000 1000000000
		1 1000000000
		3 1000000000`)}, want: 5000000000},
		{name: "3", args: args{r: strings.NewReader(`3 2
		10 11 12
		1 1000
		3 10000`)}, want: 11033},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
