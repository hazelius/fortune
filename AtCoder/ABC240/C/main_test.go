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
		{name: "1", args: args{r: strings.NewReader(`2 10
		3 6
		4 5`)}, want: "Yes"},
		{name: "2", args: args{r: strings.NewReader(`2 10
		10 100
		10 100`)}, want: "No"},
		{name: "3", args: args{r: strings.NewReader(`4 12
		1 8
		5 7
		3 4
		2 6`)}, want: "Yes"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
