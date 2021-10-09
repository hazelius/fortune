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
		{name: "1", args: args{r: strings.NewReader(`4 50
		80 60 40 0`)}, want: 2},
		{name: "2", args: args{r: strings.NewReader(`3 90
		89 89 89`)}, want: 3},
		{name: "3", args: args{r: strings.NewReader(`2 22
		6 37`)}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
