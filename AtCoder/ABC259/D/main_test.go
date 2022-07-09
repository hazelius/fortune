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
		0 -2 3 3
		0 0 2
		2 0 2
		2 3 1
		-3 3 3`)}, want: "Yes"},
		{name: "2", args: args{r: strings.NewReader(`3
		0 1 0 3
		0 0 1
		0 0 2
		0 0 3`)}, want: "No"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
