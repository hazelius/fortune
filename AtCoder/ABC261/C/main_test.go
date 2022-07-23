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
		{name: "1", args: args{r: strings.NewReader(`5
		newfile
		newfile
		newfolder
		newfile
		newfolder`)}, want: "newfile newfile(1) newfolder newfile(2) newfolder(1)"},
		{name: "2", args: args{r: strings.NewReader(`11
		a
		a
		a
		a
		a
		a
		a
		a
		a
		a
		a`)}, want: "a a(1) a(2) a(3) a(4) a(5) a(6) a(7) a(8) a(9) a(10)"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
