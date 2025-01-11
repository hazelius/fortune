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
		{name: "1", args: args{stdin: strings.NewReader(`400 500 600 700 800`)}, wantOut: `ABCDE
BCDE
ACDE
ABDE
ABCE
ABCD
CDE
BDE
ADE
BCE
ACE
BCD
ABE
ACD
ABD
ABC
DE
CE
BE
CD
AE
BD
AD
BC
AC
AB
E
D
C
B
A
`},
		{name: "2", args: args{stdin: strings.NewReader(`800 800 900 900 1000`)}, wantOut: `ABCDE
ACDE
BCDE
ABCE
ABDE
ABCD
CDE
ACE
ADE
BCE
BDE
ABE
ACD
BCD
ABC
ABD
CE
DE
AE
BE
CD
AC
AD
BC
BD
AB
E
C
D
A
B
`},
		{name: "3", args: args{stdin: strings.NewReader(`128 256 512 1024 2048`)}, wantOut: `ABCDE
BCDE
ACDE
CDE
ABDE
BDE
ADE
DE
ABCE
BCE
ACE
CE
ABE
BE
AE
E
ABCD
BCD
ACD
CD
ABD
BD
AD
D
ABC
BC
AC
C
AB
B
A
`},
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
