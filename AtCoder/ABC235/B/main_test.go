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
		{name: "1", args: args{r: strings.NewReader(`5
		1 5 10 4 2`)}, want: 10},
		{name: "2", args: args{r: strings.NewReader(`3
		100 1000 100000`)}, want: 100000},
		{name: "3", args: args{r: strings.NewReader(`4
		27 1828 1828 9242`)}, want: 1828},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
