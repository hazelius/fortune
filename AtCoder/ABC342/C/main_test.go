package main

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

func Test_run(t *testing.T) {
	type args struct {
		stdin io.Reader
	}
	tests := []struct {
		name    string
		args    args
		wantOut string
	}{
		{name: "1", args: args{stdin: strings.NewReader(`7
		atcoder
		4
		r a
		t e
		d v
		a r`)}, wantOut: `recover`},
		{name: "2", args: args{stdin: strings.NewReader(`3
		abc
		4
		a a
		s k
		n n
		z b`)}, wantOut: `abc`},
		{name: "3", args: args{stdin: strings.NewReader(`34
		supercalifragilisticexpialidocious
		20
		g c
		l g
		g m
		c m
		r o
		s e
		a a
		o f
		f s
		e t
		t l
		d v
		p k
		v h
		x i
		h n
		n j
		i r
		s i
		u a`)}, wantOut: `laklimamriiamrmrllrmlrkramrjimrial`},
		{name: "4", args: args{stdin: strings.NewReader(`7
		atcoder
		4
		r a
		a b 
		a c
		d r`)}, wantOut: `btcoreb`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := &bytes.Buffer{}
			run(tt.args.stdin, out)
			if gotOut := out.String(); gotOut != tt.wantOut {
				t.Errorf("run() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}
