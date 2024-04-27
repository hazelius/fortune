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
		{name: "1", args: args{stdin: strings.NewReader(`3
		abc
		def
		ghi
		abc
		bef
		ghi`)}, wantOut: `2 1`},
		{name: "2", args: args{stdin: strings.NewReader(`1
		f
		q`)}, wantOut: `1 1`},
		{name: "3", args: args{stdin: strings.NewReader(`10
		eixfumagit
		vtophbepfe
		pxbfgsqcug
		ugpugtsxzq
		bvfhxyehfk
		uqyfwtmglr
		jaitenfqiq
		acwvufpfvv
		jhaddglpva
		aacxsyqvoj
		eixfumagit
		vtophbepfe
		pxbfgsqcug
		ugpugtsxzq
		bvfhxyehok
		uqyfwtmglr
		jaitenfqiq
		acwvufpfvv
		jhaddglpva
		aacxsyqvoj`)}, wantOut: `5 9`},
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
