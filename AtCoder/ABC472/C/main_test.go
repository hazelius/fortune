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
		{name: "1", args: args{stdin: strings.NewReader(`5 3 83
48 73 59 90 21`)}, wantOut: `Yes
No
No
No
Yes
`},
		{name: "2", args: args{stdin: strings.NewReader(`7 4 728
187 816 349 609 255 308 175`)}, wantOut: `Yes
No
Yes
No
Yes
No
Yes
`},
		{name: "3", args: args{stdin: strings.NewReader(`10 3 1368290936
216519459 804733999 297250023 775422599 287963235 999315644 354987425 974810607 653940822 117157941`)}, wantOut: `Yes
Yes
Yes
No
Yes
Yes
No
No
Yes
Yes
`},
		{name: "4", args: args{stdin: strings.NewReader(`5 3 10
5 5 5 5 5`)}, wantOut: `Yes
Yes
No
Yes
Yes
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
