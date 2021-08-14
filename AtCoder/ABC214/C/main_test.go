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
		{name: "1", args: args{r: strings.NewReader(`3
		4 1 5
		3 10 100`)}, want: "3\n7\n8"},
		{name: "2", args: args{r: strings.NewReader(`4
		100 100 100 100
		1 1 1 1`)}, want: "1\n1\n1\n1"},
		{name: "3", args: args{r: strings.NewReader(`4
		1 2 3 4
		1 2 4 7`)}, want: "1\n2\n4\n7"},
		{name: "4", args: args{r: strings.NewReader(`8
		84 87 78 16 94 36 87 93
		50 22 63 28 91 60 64 27`)}, want: "50\n22\n63\n28\n44\n60\n64\n27"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
