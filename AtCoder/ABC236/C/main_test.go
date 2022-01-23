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
		tokyo kanda akiba okachi ueno
		tokyo akiba ueno`)}, want: "Yes No Yes No Yes"},
		{name: "2", args: args{r: strings.NewReader(`7 7
		a t c o d e r
		a t c o d e r`)}, want: "Yes Yes Yes Yes Yes Yes Yes"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
