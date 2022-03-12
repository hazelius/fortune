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
		{name: "1", args: args{r: strings.NewReader(`3 2
		URL`)}, want: 6},
		{name: "2", args: args{r: strings.NewReader(`4 500000000000000000
		RRUU`)}, want: 500000000000000000},
		{name: "3", args: args{r: strings.NewReader(`30 123456789
		LRULURLURLULULRURRLRULRRRUURRU`)}, want: 126419752371},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
