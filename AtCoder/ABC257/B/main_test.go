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
		{name: "1", args: args{r: strings.NewReader(`5 3 5
		1 3 4
		3 3 1 1 2`)}, want: "2 4 5"},
		{name: "2", args: args{r: strings.NewReader(`2 2 2
		1 2
		1 2`)}, want: "1 2"},
		{name: "3", args: args{r: strings.NewReader(`10 6 9
		1 3 5 7 8 9
		1 2 3 4 5 6 5 6 2`)}, want: "2 5 6 7 9 10"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
