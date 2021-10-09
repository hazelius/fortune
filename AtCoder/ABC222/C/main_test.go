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
		{name: "1", args: args{r: strings.NewReader(`2 3
		GCP
		PPP
		CCC
		PPC`)}, want: "3 1 2 4"},
		{name: "2", args: args{r: strings.NewReader(`2 2
		GC
		PG
		CG
		PP`)}, want: "1 2 3 4"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
