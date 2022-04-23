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
		{name: "1", args: args{r: strings.NewReader("AtCoder")}, want: "Yes"},
		{name: "2", args: args{r: strings.NewReader("Aa")}, want: "Yes"},
		{name: "3", args: args{r: strings.NewReader("atcoder")}, want: "No"},
		{name: "4", args: args{r: strings.NewReader("Perfect")}, want: "No"},
		{name: "5", args: args{r: strings.NewReader("ABC")}, want: "No"},
		{name: "6", args: args{r: strings.NewReader("AcaB")}, want: "Yes"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
