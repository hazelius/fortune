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
		{name: "1", args: args{r: strings.NewReader(`2
		1 1
		2 3`)}, want: 5},
		{name: "2", args: args{r: strings.NewReader(`3
		2 2 2
		2 2 2`)}, want: 1},
		{name: "3", args: args{r: strings.NewReader(`10
		1 2 3 4 5 6 7 8 9 10
		1 4 9 16 25 36 49 64 81 100`)}, want: 978222082},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
