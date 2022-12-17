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
		"a,b"c,d`)}, wantOut: `"a,b"c.d`},
		{name: "2", args: args{stdin: strings.NewReader(`5
		,,,,,`)}, wantOut: `.....`},
		{name: "3", args: args{stdin: strings.NewReader(`20
		a,"t,"c,"o,"d,"e,"r,`)}, wantOut: `a."t,"c."o,"d."e,"r.`},
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
