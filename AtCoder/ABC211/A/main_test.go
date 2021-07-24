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
		want float64
	}{
		{name: "1", args: args{r: strings.NewReader("130 100")}, want: 110},
		{name: "2", args: args{r: strings.NewReader("300 50")}, want: 133.33333333333334},
		{name: "3", args: args{r: strings.NewReader("123 123")}, want: 123},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
