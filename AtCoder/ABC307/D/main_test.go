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
		{name: "1", args: args{stdin: strings.NewReader(`8
		a(b(d))c`)}, wantOut: `ac`},
		{name: "2", args: args{stdin: strings.NewReader(`5
		a(b)(`)}, wantOut: `a(`},
		{name: "3", args: args{stdin: strings.NewReader(`2
		()`)}, wantOut: ``},
		{name: "4", args: args{stdin: strings.NewReader(`6
		)))(((`)}, wantOut: `)))(((`},
		{name: "5", args: args{stdin: strings.NewReader(`4
		)b(a)()()()()`)}, wantOut: `)b`},
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
