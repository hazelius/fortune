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
		{name: "1", args: args{r: strings.NewReader("pop")}, want: "o"},
		{name: "2", args: args{r: strings.NewReader("abc")}, want: "a"},
		{name: "3", args: args{r: strings.NewReader("xxx")}, want: "-1"},
		{name: "4", args: args{r: strings.NewReader("aab")}, want: "b"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
