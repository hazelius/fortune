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
		{name: "1", args: args{r: strings.NewReader(`4 3
		1 2 3
		4 5 6
		7 8 9
		10 11 12`)}, want: "1 4 7 10\n2 5 8 11\n3 6 9 12"},
		{name: "2", args: args{r: strings.NewReader(`2 2
		1000000000 1000000000
		1000000000 1000000000`)}, want: "1000000000 1000000000\n1000000000 1000000000"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
