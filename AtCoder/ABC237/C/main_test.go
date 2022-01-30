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
		{name: "1", args: args{r: strings.NewReader("kasaka")}, want: "Yes"},
		{name: "2", args: args{r: strings.NewReader("atcoder")}, want: "No"},
		{name: "3", args: args{r: strings.NewReader("php")}, want: "Yes"},
		{name: "4", args: args{r: strings.NewReader("aphp")}, want: "No"},
		{name: "5", args: args{r: strings.NewReader("phpa")}, want: "Yes"},
		{name: "6", args: args{r: strings.NewReader("a")}, want: "Yes"},
		{name: "7", args: args{r: strings.NewReader("t")}, want: "Yes"},
		{name: "8", args: args{r: strings.NewReader("taat")}, want: "Yes"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
