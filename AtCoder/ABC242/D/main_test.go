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
		{name: "1", args: args{r: strings.NewReader(`ABC
		4
		0 1
		1 1
		1 3
		1 6`)}, want: "A B C B"},
		{name: "2", args: args{r: strings.NewReader(`CBBAACCCCC
		5
		57530144230160008 659279164847814847
		29622990657296329 861239705300265164
		509705228051901259 994708708957785197
		176678501072691541 655134104344481648
		827291290937314275 407121144297426665`)}, want: "A A C A A"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
