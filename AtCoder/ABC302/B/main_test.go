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
		{name: "1", args: args{stdin: strings.NewReader(`6 6
		vgxgpu
		amkxks
		zhkbpp
		hykink
		esnuke
		zplvfj`)}, wantOut: `5 2
5 3
5 4
5 5
5 6
`},
		{name: "2", args: args{stdin: strings.NewReader(`5 5
		ezzzz
		zkzzz
		ezuzs
		zzznz
		zzzzs`)}, wantOut: `5 5
4 4
3 3
2 2
1 1
`},
		{name: "3", args: args{stdin: strings.NewReader(`10 10
		kseeusenuk
		usesenesnn
		kskekeeses
		nesnusnkkn
		snenuuenke
		kukknkeuss
		neunnennue
		sknuessuku
		nksneekknk
		neeeuknenk`)}, wantOut: `9 3
8 3
7 3
6 3
5 3
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
