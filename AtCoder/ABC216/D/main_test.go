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
		{name: "1", args: args{r: strings.NewReader(`2 2
		2
		1 2
		2
		1 2`)}, want: "Yes"},
		{name: "2", args: args{r: strings.NewReader(`2 2
		2
		1 2
		2
		2 1`)}, want: "No"},
		{name: "3", args: args{r: strings.NewReader(`3 5
		2
		7 8
		2
		7 8
		3
		3 4 1
		3
		2 4 3
		2
		2 1`)}, want: "No"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
