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
		{name: "1", args: args{stdin: strings.NewReader(`2 3
		sns
		euk`)}, wantOut: `Yes`},
		{name: "2", args: args{stdin: strings.NewReader(`2 2
		ab
		cd`)}, wantOut: `No`},
		{name: "3", args: args{stdin: strings.NewReader(`5 7
		skunsek
		nukesnu
		ukeseku
		nsnnesn
		uekukku`)}, wantOut: `Yes`},
		{name: "4", args: args{stdin: strings.NewReader(`3 3
		ssn
		neu
		ukk`)}, wantOut: `Yes`},
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
