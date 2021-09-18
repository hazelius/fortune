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
		{name: "1", args: args{r: strings.NewReader(`bacdefghijklmnopqrstuvwxzy
		4
		abx
		bzz
		bzy
		caa`)}, want: "bzz bzy abx caa"},
		{name: "2", args: args{r: strings.NewReader(`zyxwvutsrqponmlkjihgfedcba
		5
		a
		ab
		abc
		ac
		b`)}, want: "b a ac ab abc"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
