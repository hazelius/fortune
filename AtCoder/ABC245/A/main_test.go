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
		{name: "1", args: args{r: strings.NewReader("7 0 6 30")}, want: "Aoki"},
		{name: "2", args: args{r: strings.NewReader("7 30 7 30")}, want: "Takahashi"},
		{name: "3", args: args{r: strings.NewReader("0 0 23 59")}, want: "Takahashi"},
		{name: "4", args: args{r: strings.NewReader("7 2 7 1")}, want: "Aoki"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
