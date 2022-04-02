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
		{name: "1", args: args{r: strings.NewReader("3 4")}, want: "0.600000000000 0.800000000000"},
		{name: "2", args: args{r: strings.NewReader("1 0")}, want: "1.000000000000 0.000000000000"},
		{name: "3", args: args{r: strings.NewReader("246 402")}, want: "0.521964870245 0.852966983083"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
