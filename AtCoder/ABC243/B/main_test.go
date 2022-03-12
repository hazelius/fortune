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
		want string
	}{
		{name: "1", args: args{r: strings.NewReader(`4
		1 3 5 2
		2 3 1 4`)}, want: "1 2"},
		{name: "2", args: args{r: strings.NewReader(`3
		1 2 3
		4 5 6`)}, want: "0 0"},
		{name: "3", args: args{r: strings.NewReader(`7
		4 8 1 7 9 5 6
		3 5 1 7 8 2 6`)}, want: "3 2"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
