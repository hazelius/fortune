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
		{name: "1", args: args{r: strings.NewReader(`5 3
		6 8 10 7 10
		2 3 4`)}, want: "Yes"},
		{name: "2", args: args{r: strings.NewReader(`5 2
		100 100 100 1 1
		5 4`)}, want: "No"},
		{name: "3", args: args{r: strings.NewReader(`2 1
		100 1
		2`)}, want: "No"},
		{name: "4", args: args{r: strings.NewReader(`5 2
		1 2 3 3 2
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
