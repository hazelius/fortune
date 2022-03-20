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
		{name: "1", args: args{r: strings.NewReader(`4 4 4 1 3 2
		1 2
		2 3
		3 4
		1 4`)}, want: 4},
		{name: "2", args: args{r: strings.NewReader(`6 5 10 1 2 3
		2 3
		2 4
		4 6
		3 6
		1 5`)}, want: 0},
		{name: "3", args: args{r: strings.NewReader(`10 15 20 4 4 6
		2 6
		2 7
		5 7
		4 5
		2 4
		3 7
		1 7
		1 4
		2 9
		5 10
		1 3
		7 8
		7 9
		1 6
		1 2`)}, want: 952504739},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
