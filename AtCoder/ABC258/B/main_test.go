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
		{name: "1", args: args{r: strings.NewReader(`4
		1161
		1119
		7111
		1811`)}, want: 9786},
		{name: "2", args: args{r: strings.NewReader(`10
		1111111111
		1111111111
		1111111111
		1111111111
		1111111111
		1111111111
		1111111111
		1111111111
		1111111111
		1111111111`)}, want: 1111111111},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
