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
		10101
		60 45 30 40 80`)}, want: 4},
		{name: "2", args: args{r: strings.NewReader(`3
		000
		1 2 3`)}, want: 3},
		{name: "3", args: args{r: strings.NewReader(`5
		10101
		60 50 50 50 60`)}, want: 4},
		{name: "4", args: args{r: strings.NewReader(`5
		01101
		30 30 30 30 30`)}, want: 3},
		{name: "5", args: args{r: strings.NewReader(`5
		01111
		30 30 30 30 30`)}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
