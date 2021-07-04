// E - White and Black Balls
// https://atcoder.jp/contests/abc205/tasks/abc205_e
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
		want int
	}{
		{name: "1", args: args{r: strings.NewReader("2 3 1")}, want: 9},
		{name: "2", args: args{r: strings.NewReader("1 0 0")}, want: 0},
		{name: "3", args: args{r: strings.NewReader("1000000 1000000 1000000")}, want: 192151600},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
