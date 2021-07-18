// C - MAD TEAM
// https://atcoder.jp/contests/zone2021/tasks/zone2021_c
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
		{name: "1", args: args{r: strings.NewReader(`3
		3 9 6 4 6
		6 9 3 1 1
		8 8 9 3 7`)}, want: 4},
		{name: "2", args: args{r: strings.NewReader(`5
		6 13 6 19 11
		4 4 12 11 18
		20 7 19 2 5
		15 5 12 20 7
		8 7 6 18 5`)}, want: 13},
		{name: "3", args: args{r: strings.NewReader(`10
		6 7 5 18 2
		3 8 1 6 3
		7 2 8 7 7
		6 3 3 4 7
		12 8 9 15 9
		9 8 6 1 10
		12 9 7 8 2
		10 3 17 4 10
		3 1 3 19 3
		3 14 7 13 1`)}, want: 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
