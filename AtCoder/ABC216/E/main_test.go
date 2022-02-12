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
		{name: "1", args: args{r: strings.NewReader(`3 5
		100 50 102`)}, want: 502},
		{name: "2", args: args{r: strings.NewReader(`2 2021
		2 3`)}, want: 9},
		{name: "3", args: args{r: strings.NewReader(`3 4
		150 120 151`)}, want: 600},
		{name: "4", args: args{r: strings.NewReader(`10 9
		1 1 1 1 1 1 1 1 1 1`)}, want: 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
