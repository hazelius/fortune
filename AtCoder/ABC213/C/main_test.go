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
		{name: "1", args: args{r: strings.NewReader(`4 5 2
3 2
		2 5`)}, want: `2 1
1 2`},
		{name: "2", args: args{r: strings.NewReader(`1000000000 1000000000 10
1 1
10 10
100 100
1000 1000
10000 10000
100000 100000
1000000 1000000
10000000 10000000
100000000 100000000
1000000000 1000000000`)}, want: `1 1
2 2
3 3
4 4
5 5
6 6
7 7
8 8
9 9
10 10`},
		{name: "3", args: args{r: strings.NewReader(`8 8 4
		1 1
		1 2
		5 6
		100 2008`)}, want: `1 1
1 2
2 3
3 4`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
