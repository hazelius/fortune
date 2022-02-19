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
		{name: "1", args: args{r: strings.NewReader("333")}, want: 65287.9076782217},
		{name: "2", args: args{r: strings.NewReader("634")}, want: 90086.63583462311},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
