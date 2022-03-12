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
		2 3
		1 1
		4 1
		RRL`)}, want: "Yes"},
		{name: "2", args: args{r: strings.NewReader(`2
		1 1
		2 1
		RR`)}, want: "No"},
		{name: "3", args: args{r: strings.NewReader(`10
		1 3
		1 4
		0 0
		0 2
		0 4
		3 1
		2 4
		4 2
		4 4
		3 3
		RLRRRLRLRR`)}, want: "Yes"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
