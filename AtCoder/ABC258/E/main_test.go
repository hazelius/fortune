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
		{name: "1", args: args{r: strings.NewReader(`3 2 5
		3 4 1
		1
		2`)}, want: "2 3"},
		{name: "2", args: args{r: strings.NewReader(`10 7 20
		5 8 5 9 8 7 4 4 8 2
		1
		9
		10
		1000
		1000000
		1000000000
		1000000000000`)}, want: "4 3 4 5 5 5 5"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
