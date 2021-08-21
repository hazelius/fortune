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
		{name: "1", args: args{r: strings.NewReader("aab 2")}, want: "aba"},
		{name: "2", args: args{r: strings.NewReader("baba 4")}, want: "baab"},
		{name: "3", args: args{r: strings.NewReader("ydxwacbz 40320")}, want: "zyxwdcba"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
