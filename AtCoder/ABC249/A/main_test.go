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
		{name: "1", args: args{r: strings.NewReader("4 3 3 6 2 5 10")}, want: "Takahashi"},
		{name: "2", args: args{r: strings.NewReader("3 1 4 1 5 9 2")}, want: "Aoki"},
		{name: "3", args: args{r: strings.NewReader("1 1 1 1 1 1 1")}, want: "Draw"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
