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
		{name: "1", args: args{r: strings.NewReader(`2 2
		1 6
		4 9`)}, want: 2},
		{name: "2", args: args{r: strings.NewReader(`1 1
		10
		10`)}, want: 0},
		{name: "3", args: args{r: strings.NewReader(`6 8
		82 76 82 82 71 70
		17 39 67 2 45 35 22 24`)}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
