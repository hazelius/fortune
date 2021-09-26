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
		{name: "1", args: args{r: strings.NewReader(`3
		2 7 6`)}, want: "1 0 0 0 2 1 0 0 0 0"},
		{name: "2", args: args{r: strings.NewReader(`5
		0 1 2 3 4`)}, want: "6 0 1 1 4 0 1 1 0 2"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
