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
		{name: "1", args: args{r: strings.NewReader("0 0 3 3")}, want: "Yes"},
		{name: "2", args: args{r: strings.NewReader("0 1 2 3")}, want: "No"},
		{name: "3", args: args{r: strings.NewReader("1000000000 1000000000 999999999 999999999")}, want: "Yes"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
