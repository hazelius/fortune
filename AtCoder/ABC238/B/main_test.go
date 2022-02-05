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
		90 180 45 195`)}, want: 120},
		{name: "2", args: args{r: strings.NewReader(`1
		1`)}, want: 359},
		{name: "3", args: args{r: strings.NewReader(`10
		215 137 320 339 341 41 44 18 241 149`)}, want: 170},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
