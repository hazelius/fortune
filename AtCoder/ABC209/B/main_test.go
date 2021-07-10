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
		{name: "1", args: args{r: strings.NewReader(`2 3
		1 3`)}, want: "Yes"},
		{name: "2", args: args{r: strings.NewReader(`4 10
		3 3 4 4`)}, want: "No"},
		{name: "3", args: args{r: strings.NewReader(`8 30
		3 1 4 1 5 9 2 6`)}, want: "Yes"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
