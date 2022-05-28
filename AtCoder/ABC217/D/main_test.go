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
		2 2
		1 3
		2 2`)}, want: "5 3"},
		{name: "2", args: args{r: strings.NewReader(`5 3
		1 2
		1 4
		2 3`)}, want: "2"},
		{name: "3", args: args{r: strings.NewReader(`100 10
		1 31
		2 41
		1 59
		2 26
		1 53
		2 58
		1 97
		2 93
		1 23
		2 84`)}, want: "69 31 6 38 38"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
