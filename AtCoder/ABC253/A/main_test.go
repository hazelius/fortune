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
		{name: "1", args: args{r: strings.NewReader("5 3 2")}, want: "Yes"},
		{name: "2", args: args{r: strings.NewReader("2 5 3")}, want: "No"},
		{name: "3", args: args{r: strings.NewReader("100 100 100")}, want: "Yes"},
		{name: "4", args: args{r: strings.NewReader("5 3 3")}, want: "Yes"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
