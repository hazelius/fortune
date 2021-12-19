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
		{name: "1", args: args{r: strings.NewReader(`4 4
		1 2
		1 3
		1 4
		3 4
		1 3
		1 4
		2 3
		3 4`)}, want: "Yes"},
		{name: "2", args: args{r: strings.NewReader(`5 6
		1 2
		1 3
		1 4
		3 4
		3 5
		4 5
		1 2
		1 3
		1 4
		1 5
		3 5
		4 5`)}, want: "No"},
		{name: "3", args: args{r: strings.NewReader("8 0")}, want: "Yes"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
