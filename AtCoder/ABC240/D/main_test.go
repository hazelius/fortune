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
		{name: "1", args: args{r: strings.NewReader(`5
		3 2 3 2 2`)}, want: "1 2 3 4 3"},
		{name: "2", args: args{r: strings.NewReader(`10
		2 3 2 3 3 3 2 3 3 2`)}, want: "1 2 3 4 5 3 2 3 1 0"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
