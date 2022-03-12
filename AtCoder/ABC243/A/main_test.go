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
		{name: "1", args: args{r: strings.NewReader("25 10 11 12")}, want: "T"},
		{name: "2", args: args{r: strings.NewReader("30 10 10 10")}, want: "F"},
		{name: "3", args: args{r: strings.NewReader("100000 1 1 1")}, want: "M"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
