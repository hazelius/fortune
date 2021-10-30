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
		1 4
		2 4
		3 4
		4 5`)}, want: "Yes"},
		{name: "2", args: args{r: strings.NewReader(`4
		2 4
		1 4
		2 3`)}, want: "No"},
		{name: "3", args: args{r: strings.NewReader(`10
		9 10
		3 10
		4 10
		8 10
		1 10
		2 10
		7 10
		6 10
		5 10`)}, want: "Yes"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
