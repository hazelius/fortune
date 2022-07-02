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
		{name: "1", args: args{r: strings.NewReader(`3 4
		3 4
		2 3
		4 2`)}, want: 18},
		{name: "2", args: args{r: strings.NewReader(`10 1000000000
		3 3
		1 6
		4 7
		1 8
		5 7
		9 9
		2 4
		6 4
		5 1
		3 1`)}, want: 1000000076},
		{name: "3", args: args{r: strings.NewReader(`3 3
		3 3
		2 1
		1 1`)}, want: 10},
		{name: "4", args: args{r: strings.NewReader(`4 1
		5 5
		2 2
		4 10
		3 1`)}, want: 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
