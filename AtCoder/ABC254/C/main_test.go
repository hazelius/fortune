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
		{name: "1", args: args{r: strings.NewReader(`5 2
		3 4 1 3 4`)}, want: "Yes"},
		{name: "2", args: args{r: strings.NewReader(`5 3
		3 4 1 3 4`)}, want: "No"},
		{name: "3", args: args{r: strings.NewReader(`7 5
		1 2 3 4 5 5 10`)}, want: "Yes"},
		{name: "4", args: args{r: strings.NewReader(`7 3
		1 4 2 3 5 5 10`)}, want: "No"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
